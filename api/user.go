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
	USERID_NOT_FOUND            = "user with specified ID '%s' does not exist"
	USEREMAIL_NOT_FOUND         = "%s is not registered to an account"
	LOGGED_IN_USER_DOESNT_MATCH = "authenticated user does not match the requested account"
)

type UserResponse struct {
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	EmailAddress      string    `json:"emailAddress"`
	PasswordChangedAt time.Time `json:"passwordChangedAt"`
	ID                string    `json:"id"`
}

type createUserRequest struct {
	FirstName    string `json:"firstName"    binding:"required"`
	LastName     string `json:"lastName"     binding:"required"`
	EmailAddress string `json:"emailAddress" binding:"required,email"`
	Password     string `json:"password"     binding:"required,gt=8"`
}

func (server *Server) createUser(ctx *gin.Context) {
	req := createUserRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.CreateUserParams{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		EmailAddress: req.EmailAddress,
		Password:     hashedPassword,
	}

	_, err = server.store.NewUserTx(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
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

	ctx.JSON(http.StatusOK, UserResponse{
		ID:                user.ID,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		EmailAddress:      user.EmailAddress,
		PasswordChangedAt: user.PasswordChangedAt,
	})
}

type getUserByEmailRequest struct {
	EmailAddress string `uri:"email" binding:"required,email"`
}

func (server *Server) getUserByEmail(ctx *gin.Context) {
	req := getUserByEmailRequest{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authHeader := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

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

	if user.ID != authHeader.UserID {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New(LOGGED_IN_USER_DOESNT_MATCH)))
		return
	}

	ctx.JSON(http.StatusOK, UserResponse{
		ID:                user.ID,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		EmailAddress:      user.EmailAddress,
		PasswordChangedAt: user.PasswordChangedAt,
	})
}

type updateUserRequest struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"emailAddress" binding:"omitempty,email"`
	ID           string `json:"id"           binding:"required"`
}

func (server *Server) updateUser(ctx *gin.Context) {
	req := updateUserRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authHeader := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	user, err := server.store.GetUserById(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf(USERID_NOT_FOUND, req.ID)))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if user.ID != authHeader.UserID {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New(LOGGED_IN_USER_DOESNT_MATCH)))
		return
	}

	args := db.UpdateUserParams{
		FirstName: sql.NullString{
			Valid:  req.FirstName != "",
			String: req.FirstName,
		},
		LastName: sql.NullString{
			Valid:  req.LastName != "",
			String: req.LastName,
		},
		EmailAddress: sql.NullString{
			Valid:  req.EmailAddress != "",
			String: req.EmailAddress,
		},
		UserID: req.ID,
	}

	err = server.store.UpdateUser(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	user, err = server.store.GetUserById(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, UserResponse{
		ID:                user.ID,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		EmailAddress:      user.EmailAddress,
		PasswordChangedAt: user.PasswordChangedAt,
	})
}

type changePasswordRequest struct {
	ID              string `json:"id"              binding:"required"`
	CurrentPassword string `json:"currentPassword" binding:"required,gt=8"`
	NewPassword     string `json:"newPassword"     binding:"required,gt=8"`
}

func (server *Server) changePassword(ctx *gin.Context) {
	req := changePasswordRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authHeader := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	user, err := server.store.GetUserById(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf(USERID_NOT_FOUND, req.ID)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if user.ID != authHeader.UserID {
		ctx.JSON(http.StatusUnauthorized, errors.New(LOGGED_IN_USER_DOESNT_MATCH))
		return
	}

	if user.Password != req.CurrentPassword {
		ctx.JSON(
			http.StatusUnauthorized,
			errorResponse(errors.New("your current password is incorrect")),
		)
		return
	}

	args := db.ChangePasswordParams{
		Password: req.NewPassword,
		UserID:   req.ID,
	}

	err = server.store.ChangePassword(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

type deleteUserRequest struct {
	ID string `uri:"id" binding:"required"`
}

func (server *Server) deleteUser(ctx *gin.Context) {
	req := deleteUserRequest{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authHeader := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	user, err := server.store.GetUserById(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf(USERID_NOT_FOUND, req.ID)))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if user.ID != authHeader.UserID {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New(LOGGED_IN_USER_DOESNT_MATCH)))
		return
	}

	err = server.store.DeleteUser(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
