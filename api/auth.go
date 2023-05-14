package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	EmailAddress string `json:"emailAddress" binding:"required,email"`
	Password     string `json:"password" binding:"required"`
}

type loginResponse struct {
	AccessToken string       `json:"accessToken"`
	User        UserResponse `json:"user"`
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

	token, err := server.tokenMaker.CreateToken(user.ID, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, loginResponse{
		AccessToken: token,
		User: UserResponse{
			ID:                user.ID,
			FirstName:         user.FirstName,
			LastName:          user.LastName,
			EmailAddress:      user.EmailAddress,
			PasswordChangedAt: user.PasswordChangedAt,
		},
	})
}
