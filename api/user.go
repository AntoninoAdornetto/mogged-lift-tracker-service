package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/AntoninoAdornetto/mogged-lift-tracker-service/db/sqlc"
	"github.com/gin-gonic/gin"
)

type userResponse struct {
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	EmailAddress      string    `json:"emailAddress"`
	PasswordChangedAt time.Time `json:"passwordChangedAt"`
	ID                string    `json:"id"`
}

type createUserRequest struct {
	FirstName    string `json:"firstName" binding:"required"`
	LastName     string `json:"lastName" binding:"required"`
	EmailAddress string `json:"emailAddress" binding:"required,email"`
	Password     string `json:"password" binding:"required,gt=8"`
}

func (server *Server) createUser(ctx *gin.Context) {
	req := &createUserRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.CreateUserParams{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		EmailAddress: req.EmailAddress,
		Password:     req.Password,
	}

	_, err := server.store.NewUserTx(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	user, err := server.store.GetUserByEmail(ctx, req.EmailAddress)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userResponse{
		ID:                user.ID,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		EmailAddress:      user.EmailAddress,
		PasswordChangedAt: user.PasswordChangedAt,
	})
}

type getUserRequest struct {
	EmailAddress string `uri:"email" binding:"required,email"`
}

func (server *Server) getUserByEmail(ctx *gin.Context) {
	req := &getUserRequest{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUserByEmail(ctx, req.EmailAddress)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userResponse{
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
	ID           string `json:"id" binding:"required"`
}

func (server *Server) updateUser(ctx *gin.Context) {
	req := &updateUserRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err := server.store.GetUserById(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var firstName, lastName, emailAddress sql.NullString

	if req.FirstName != "" {
		firstName.String = req.FirstName
		firstName.Valid = true
	} else {
		firstName.Valid = false
	}

	if req.LastName != "" {
		lastName.String = req.LastName
		lastName.Valid = true
	} else {
		lastName.Valid = false
	}

	if req.EmailAddress != "" {
		emailAddress.String = req.EmailAddress
		emailAddress.Valid = true
	} else {
		emailAddress.Valid = false
	}

	err = server.store.UpdateUser(ctx, db.UpdateUserParams{
		FirstName:    firstName,
		LastName:     lastName,
		EmailAddress: emailAddress,
		UserID:       req.ID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
