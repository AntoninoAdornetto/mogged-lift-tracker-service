package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/AntoninoAdornetto/mogged-lift-tracker-service/db/sqlc"
	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/token"
	"github.com/gin-gonic/gin"
)

const (
	LIFT_NOT_FOUND = "lift with specified ID '%d' does not exist"
)

type LiftResponse struct {
	ID           int64   `json:"id"`
	ExerciseName string  `json:"exerciseName"`
	WeightLifted float64 `json:"weightLifted"`
	Reps         int32   `json:"reps"`
	SetType      string  `json:"setType"`
	UserID       string  `json:"userID"`
	WorkoutID    int32   `json:"workoutID"`
}

type getLiftRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) getLift(ctx *gin.Context) {
	userID := ctx.MustGet(authorizationPayloadKey).(*token.Payload).UserID

	req := getLiftRequest{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	lift, err := server.store.GetLift(ctx, db.GetLiftParams{ID: req.ID, UserID: userID})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf(LIFT_NOT_FOUND, req.ID)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, LiftResponse{
		ID:           lift.ID,
		Reps:         lift.Reps,
		WeightLifted: lift.WeightLifted,
		ExerciseName: lift.ExerciseName,
		SetType:      lift.SetType,
		UserID:       userID,
		WorkoutID:    lift.WorkoutID,
	})
}

type listLiftsFromWorkoutRequest struct {
	WorkoutID int32 `uri:"workout_id"`
}

func (server *Server) listLiftsFromWorkout(ctx *gin.Context) {
	userID := ctx.MustGet(authorizationPayloadKey).(*token.Payload).UserID

	req := listLiftsFromWorkoutRequest{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	lifts, err := server.store.ListLiftsFromWorkout(ctx, db.ListLiftsFromWorkoutParams{
		WorkoutID: req.WorkoutID,
		UserID:    userID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := make([]LiftResponse, len(lifts))
	for i, v := range lifts {
		res[i] = LiftResponse{
			ID:           v.ID,
			Reps:         v.Reps,
			WeightLifted: v.WeightLifted,
			ExerciseName: v.ExerciseName,
			SetType:      v.SetType,
			UserID:       userID,
			WorkoutID:    v.WorkoutID,
		}
	}

	ctx.JSON(http.StatusOK, res)
}
