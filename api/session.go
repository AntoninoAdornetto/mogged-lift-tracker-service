package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	SESSION_NOT_ESTABLISHED = "No session has been established. Login or register for an account"
)

type validateSessionResponse struct {
	IsLoggedIn   bool   `json:"isLoggedIn"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"emailAddress"`
	UserID       string `json:"userID"`
}

func (server *Server) validateSession(ctx *gin.Context) {
	sessionID, err := ctx.Cookie("session_id")
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(errors.New(SESSION_NOT_ESTABLISHED)))
		return
	}

	session, err := server.store.GetSession(ctx, sessionID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(errors.New(SESSION_NOT_FOUND)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	user, err := server.store.GetUserById(ctx, session.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(
				http.StatusNotFound,
				errorResponse(fmt.Errorf(USERID_NOT_FOUND, session.UserID)),
			)
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// @TODO - Get Profile settings. Such as measurement system, timezone?
	ctx.JSON(http.StatusOK, validateSessionResponse{
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		EmailAddress: user.EmailAddress,
		IsLoggedIn:   time.Now().Before(session.ExpiresAt),
		UserID:       user.ID,
	})
}
