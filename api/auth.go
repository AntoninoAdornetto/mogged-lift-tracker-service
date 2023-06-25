package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	db "github.com/AntoninoAdornetto/mogged-lift-tracker-service/db/sqlc"
	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/token"
	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/gin-gonic/gin"
)

const (
	SESSION_NOT_FOUND    = "Failed to find active session. Please login again"
	SESSION_EXPIRED      = "Session has expired. Please login again"
	SESSION_MISMATCH     = "Session does not belong to requested user. Please login again"
	SESSION_INVALID      = "Invalid session. Please login again"
	INVALID_SESSION_ARGS = "Session credentials were not provided. Please login again"
)

type loginRequest struct {
	EmailAddress string `json:"emailAddress" binding:"required,email"`
	Password     string `json:"password"     binding:"required"`
}

type loginResponse struct {
	AccessToken           string       `json:"accessToken"`
	AccessTokenExpiresAt  time.Time    `json:"accessTokenExpiresAt"`
	RefreshToken          string       `json:"refreshToken"`
	RefreshTokenExpiresAt time.Time    `json:"refreshTokenExpiresAt"`
	SessionID             string       `json:"sessionID"`
	User                  UserResponse `json:"user"`
}

func (server *Server) login(ctx *gin.Context) {
	req := loginRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUserByEmail(ctx, req.EmailAddress)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(
				http.StatusNotFound,
				errorResponse(fmt.Errorf(USEREMAIL_NOT_FOUND, req.EmailAddress)),
			)
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if err = util.ValidatePassword(req.Password, user.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("incorrect password")))
		return
	}

	token, accessTokenPayload, err := server.tokenMaker.CreateToken(
		user.ID,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	refreshToken, refreshTokenPayload, err := server.tokenMaker.CreateToken(
		user.ID,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.DeleteSession(ctx, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshTokenPayload.ID.String(),
		RefreshToken: refreshToken,
		UserID:       user.ID,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		ExpiresAt:    refreshTokenPayload.ExpiredAt,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	session, err := server.store.GetSession(ctx, refreshTokenPayload.ID.String())
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(
				http.StatusNotFound,
				errorResponse(errors.New("session not found, please login again")),
			)
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "session_id",
		Value:    session.ID,
		Secure:   true,
		HttpOnly: true,
		Expires:  session.ExpiresAt,
		SameSite: http.SameSiteStrictMode,
		Path:     "/api",
	})

	ctx.JSON(http.StatusOK, loginResponse{
		AccessToken:           token,
		RefreshToken:          session.RefreshToken,
		AccessTokenExpiresAt:  accessTokenPayload.ExpiredAt,
		RefreshTokenExpiresAt: refreshTokenPayload.ExpiredAt,
		SessionID:             refreshTokenPayload.ID.String(),
		User: UserResponse{
			ID:                user.ID,
			FirstName:         user.FirstName,
			LastName:          user.LastName,
			EmailAddress:      user.EmailAddress,
			PasswordChangedAt: user.PasswordChangedAt,
		},
	})
}

type renewTokenRequest struct {
	RefreshToken        string `json:"refreshToken"`
	RefreshTokenPayload *token.Payload
	SessionID           string
	Validated           bool
}

type renewTokenResponse struct {
	AccessToken          string    `json:"accessToken"`
	AccessTokenExpiresAt time.Time `json:"accessTokenExpiresAt"`
}

func (server *Server) renewToken(ctx *gin.Context) {
	req := &renewTokenRequest{}

	if err := ctx.ShouldBindJSON(req); err != nil {
		sessionID, err := ctx.Cookie("session_id")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New(INVALID_SESSION_ARGS)))
			return
		}
		req.SessionID = sessionID
	} else {
		payload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New(SESSION_EXPIRED)))
			return
		}
		req.SessionID = payload.ID.String()
		req.Validated = true
		req.RefreshTokenPayload = payload
	}

	session, err := server.store.GetSession(ctx, req.SessionID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(errors.New(SESSION_NOT_FOUND)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// this check is invoked only when the session ID is pulled from http cookie.
	// when caller makes request with refresh token, this check is not hit
	if !req.Validated {
		payload, err := server.tokenMaker.VerifyToken(session.RefreshToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New(SESSION_EXPIRED)))
			return
		}
		req.RefreshTokenPayload = payload
	}

	if session.IsBanned {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("session is blocked")))
		return
	}

	if session.UserID != req.RefreshTokenPayload.UserID {
		ctx.JSON(
			http.StatusUnauthorized,
			errorResponse(errors.New(SESSION_MISMATCH)),
		)
		return
	}

	if time.Now().After(session.ExpiresAt) {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New(SESSION_EXPIRED)))
		return
	}

	accessToken, accessTokenPayload, err := server.tokenMaker.CreateToken(
		req.RefreshTokenPayload.UserID,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, renewTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessTokenPayload.ExpiredAt,
	})
}
