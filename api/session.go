package api

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	SESSION_NOT_ESTABLISHED = "No session has been established. Login or register for an account"
)

type validateSessionResponse struct {
	IsLoggedIn bool `json:"isLoggedIn"`
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

	ctx.JSON(http.StatusOK, validateSessionResponse{
		IsLoggedIn: time.Now().Before(session.ExpiresAt),
	})
}

// type getSessionResponse struct {
// 	IsLoggedIn bool `json:"isLoggedIn"`
// }

// func (server *Server) getSession(ctx *gin.Context) {
// 	sessionID, err := ctx.Cookie("session_id")
// 	if err != nil {
// 		ctx.JSON(http.StatusNotFound, errorResponse(errors.New(SESSION_NOT_ESTABLISHED)))
// 		return
// 	}

// 	session, err := server.store.GetSession(ctx, sessionID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(errors.New(SESSION_NOT_FOUND)))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, getSessionResponse{
// 		IsLoggedIn: time.Now().Before(session.ExpiresAt),
// 	})
// }
