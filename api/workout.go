package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	db "github.com/AntoninoAdornetto/mogged-lift-tracker-service/db/sqlc"
	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/token"
	"github.com/gin-gonic/gin"
)

type WorkoutResponse struct {
	ID       int32           `json:"id"`
	Duration string          `json:"duration"`
	Lifts    json.RawMessage `json:"lifts"`
	UserID   string          `json:"userID"`
}

type createWorkoutRequest struct {
	UserID   string          `json:"userID" binding:"required"`
	Duration string          `json:"duration" binding:"required"`
	Lifts    json.RawMessage `json:"lifts" binding:"required"`
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
	ID     int32  `uri:"id" binding:"required"`
	UserID string `uri:"user_id" binding:"required"`
}

func (server *Server) getWorkout(ctx *gin.Context) {
	req := getWorkoutRequest{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	workout, err := server.store.GetWorkout(ctx, db.GetWorkoutParams{ID: req.ID, UserID: req.UserID})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
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
