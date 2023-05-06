package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/AntoninoAdornetto/mogged-lift-tracker-service/db/sqlc"
	"github.com/gin-gonic/gin"
)

const (
	PROFILE_NOT_FOUND = "user profile with specified user id '%s' does not exist"
)

type ProfileResponse struct {
	Country           string  `json:"country"`
	MeasurementSystem string  `json:"measurementSystem"`
	BodyWeight        float64 `json:"bodyWeight"`
	BodyFat           float64 `json:"bodyFat"`
	TimeZoneOffset    int32   `json:"timeZoneOffset"`
}

type createProfileRequest struct {
	Country           string  `json:"country" binding:"required"`
	MeasurementSystem string  `json:"measurementSystem" binding:"required"`
	BodyWeight        float64 `json:"bodyWeight" binding:"required"`
	BodyFat           float64 `json:"bodyFat" binding:"required"`
	TimeZoneOffset    int32   `json:"timeZoneOffset" binding:"required"`
	UserID            string  `json:"userID" binding:"required"`
}

func (server *Server) createProfile(ctx *gin.Context) {
	req := createProfileRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.CreateProfileParams{
		Country:           req.Country,
		MeasurementSystem: req.MeasurementSystem,
		BodyWeight:        req.BodyWeight,
		BodyFat:           req.BodyFat,
		TimezoneOffset:    req.TimeZoneOffset,
		UserID:            req.UserID,
	}

	_, err := server.store.CreateProfile(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	profile, err := server.store.GetProfile(ctx, req.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf(PROFILE_NOT_FOUND, req.UserID)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ProfileResponse{
		Country:           profile.Country,
		MeasurementSystem: profile.MeasurementSystem,
		BodyWeight:        profile.BodyWeight,
		BodyFat:           profile.BodyFat,
		TimeZoneOffset:    profile.TimezoneOffset,
	})
}

type getProfileRequest struct {
	UserID string `uri:"user_id" binding:"required"`
}

func (server *Server) getProfile(ctx *gin.Context) {
	req := getProfileRequest{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadGateway, errorResponse(err))
		return
	}

	profile, err := server.store.GetProfile(ctx, req.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf(PROFILE_NOT_FOUND, req.UserID)))
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ProfileResponse{
		Country:           profile.Country,
		MeasurementSystem: profile.MeasurementSystem,
		BodyWeight:        profile.BodyWeight,
		BodyFat:           profile.BodyFat,
		TimeZoneOffset:    profile.TimezoneOffset,
	})
}

// type updateProfileRequest struct {
// 	Country           string  `json:"country" binding:"required"`
// 	MeasurementSystem string  `json:"measurementSystem" binding:"required"`
// 	BodyWeight        float64 `json:"bodyWeight" binding:"required"`
// 	BodyFat           float64 `json:"bodyFat" binding:"required"`
// 	TimeZoneOffset    int32   `json:"timeZoneOffset" binding:"required"`
// }

// func (server *Server) updateProfile(ctx *gin.Context) {
// 	req := updateProfileRequest{}
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	var country, measurementSystem sql.NullString
// 	var bodyWeight, bodyFat sql.NullFloat64
// 	var timeZoneOffset sql.NullInt32
// }
