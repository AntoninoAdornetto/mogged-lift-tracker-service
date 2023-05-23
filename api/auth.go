package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	db "github.com/AntoninoAdornetto/mogged-lift-tracker-service/db/sqlc"
	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	EmailAddress string `json:"emailAddress" binding:"required,email"`
	Password     string `json:"password" binding:"required"`
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
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf(USEREMAIL_NOT_FOUND, req.EmailAddress)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if err = util.ValidatePassword(req.Password, user.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("incorrect password")))
		return
	}

	token, accessTokenPayload, err := server.tokenMaker.CreateToken(user.ID, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	refreshToken, refreshTokenPayload, err := server.tokenMaker.CreateToken(user.ID, server.config.RefreshTokenDuration)
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
			ctx.JSON(http.StatusNotFound, errorResponse(errors.New("session not found, please login again")))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

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
	RefreshToken string `json:"refreshToken" binding:"required"`
}

type renewTokenResponse struct {
	AccessToken          string    `json:"accessToken"`
	AccessTokenExpiresAt time.Time `json:"accessTokenExpiresAt"`
}

// Look into a cookie to store the refresh token so that we can potentially keep the user
// logged in for a good amount of time.
func (server *Server) renewToken(ctx *gin.Context) {
	req := renewTokenRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	refreshTokenPayload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("session has expired, please login again")))
		return
	}

	session, err := server.store.GetSession(ctx, refreshTokenPayload.ID.String())
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(errors.New("failed to fetch session, please login again")))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if session.IsBanned {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("session is blocked")))
		return
	}

	if session.UserID != refreshTokenPayload.UserID {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("session does not belong to requested user")))
		return
	}

	if session.RefreshToken != req.RefreshToken {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("invalid session, please login again")))
		return
	}

	if time.Now().After(session.ExpiresAt) {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("expired session")))
		return
	}

	accessToken, accessTokenPayload, err := server.tokenMaker.CreateToken(refreshTokenPayload.UserID, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, renewTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessTokenPayload.ExpiredAt,
	})
}
