package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	db "github.com/AntoninoAdornetto/mogged-lift-tracker-service/db/sqlc"
	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/token"
	"github.com/gin-gonic/gin"
)

const (
	WORKOUT_NOT_FOUND = "workout with specified ID '%d' does not exist"
)

type WorkoutResponse struct {
	ID            int32           `json:"id"`
	CompletedDate time.Time       `json:"completedDate"`
	Duration      string          `json:"duration"`
	UserID        string          `json:"userID"`
	Lifts         json.RawMessage `json:"lifts"`
}

type createWorkoutRequest struct {
	UserID   string          `json:"userID"   binding:"required"`
	Duration string          `json:"duration" binding:"required"`
	Lifts    json.RawMessage `json:"lifts"    binding:"required"`
}

func (server *Server) createWorkout(ctx *gin.Context) {
	req := createWorkoutRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	workout, err := server.store.WorkoutTx(ctx, db.WorkoutTxParams{
		UserID:   req.UserID,
		LiftsMap: req.Lifts,
		Duration: req.Duration,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, WorkoutResponse{
		ID:       workout.ID,
		Duration: workout.Duration,
		Lifts:    workout.Lifts,
		UserID:   req.UserID,
	})
}

type getWorkoutRequest struct {
	ID     int32  `uri:"id"      binding:"required"`
	UserID string `uri:"user_id" binding:"required"`
}

func (server *Server) getWorkout(ctx *gin.Context) {
	req := getWorkoutRequest{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	workout, err := server.store.GetWorkout(
		ctx,
		db.GetWorkoutParams{ID: req.ID, UserID: req.UserID},
	)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf(WORKOUT_NOT_FOUND, req.ID)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, WorkoutResponse{
		ID:            workout.ID,
		Duration:      workout.Duration,
		Lifts:         workout.Lifts,
		UserID:        req.UserID,
		CompletedDate: workout.CompletedDate.Time,
	})
}

type listWorkoutsResponse struct {
	Workouts []db.Workout `json:"workouts"`
}

func (server *Server) listWorkouts(ctx *gin.Context) {
	authHeader := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	workouts, err := server.store.ListWorkouts(ctx, authHeader.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//@TODO: remove binary value for userID and create new struct to match response body of other
	// JSON responses
	ctx.JSON(http.StatusOK, listWorkoutsResponse{
		Workouts: workouts,
	})
}

type updateWorkoutRequest struct {
	ID            int32           `json:"id"`
	Duration      string          `json:"duration"`
	Lifts         json.RawMessage `json:"lifts"`
	CompletedDate time.Time       `json:"completedDate"`
}

func (server *Server) updateWorkout(ctx *gin.Context) {
	req := updateWorkoutRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authHeader := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	args := db.UpdateWorkoutParams{
		ID:    req.ID,
		Lifts: req.Lifts,
		Duration: sql.NullString{
			String: req.Duration,
			Valid:  req.Duration != "",
		},
		CompletedDate: sql.NullTime{
			Time:  req.CompletedDate,
			Valid: req.CompletedDate.Before(time.Now()),
		},
		UserID: authHeader.UserID,
	}

	err := server.store.UpdateWorkout(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	workout, err := server.store.GetWorkout(
		ctx,
		db.GetWorkoutParams{ID: req.ID, UserID: authHeader.UserID},
	)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf(WORKOUT_NOT_FOUND, req.ID)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, WorkoutResponse{
		ID:            workout.ID,
		CompletedDate: workout.CompletedDate.Time,
		Duration:      workout.Duration,
		Lifts:         workout.Lifts,
		UserID:        authHeader.UserID,
	})
}

type deleteWorkoutRequest struct {
	ID int32 `uri:"id"`
}

func (server *Server) deleteWorkout(ctx *gin.Context) {
	req := deleteWorkoutRequest{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID := ctx.MustGet(authorizationPayloadKey).(*token.Payload).UserID

	_, err := server.store.GetWorkout(ctx, db.GetWorkoutParams{ID: req.ID, UserID: userID})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf(WORKOUT_NOT_FOUND, req.ID)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.DeleteWorkout(ctx, db.DeleteWorkoutParams{ID: req.ID, UserID: userID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
