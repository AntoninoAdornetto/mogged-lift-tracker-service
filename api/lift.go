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
	LIFT_NOT_FOUND         = "lift with specified ID '%d' does not exist"
	MUSCLE_GROUP_NOT_FOUND = "exercise with specified name '%s' does not exist"
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
	WorkoutID int32 `uri:"workout_id" binding:"required"`
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

type getMaxLiftsRequest struct {
	Limit int32 `uri:"limit" binding:"required"`
}

func (server *Server) getMaxLifts(ctx *gin.Context) {
	userID := ctx.MustGet(authorizationPayloadKey).(*token.Payload).UserID

	req := getMaxLiftsRequest{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	lifts, err := server.store.GetMaxLifts(ctx, db.GetMaxLiftsParams{
		Limit:  req.Limit,
		UserID: userID,
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

type getMaxLiftsByExerciseRequest struct {
	ExerciseName string `uri:"exercise_name" binding:"required"`
}

func (server *Server) getMaxLiftsByExercise(ctx *gin.Context) {
	userID := ctx.MustGet(authorizationPayloadKey).(*token.Payload).UserID

	req := getMaxLiftsByExerciseRequest{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	exercise, err := server.store.GetExerciseByName(ctx, db.GetExerciseByNameParams{Name: req.ExerciseName, UserID: userID})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf(EXERCISE_NAME_NOT_FOUND, req.ExerciseName)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	lifts, err := server.store.GetMaxLiftsByExercise(ctx, db.GetMaxLiftsByExerciseParams{
		ExerciseName: exercise.Name,
		UserID:       userID,
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

type getMaxLiftsByMuscleGroupRequest struct {
	MuscleGroup string `uri:"muscle_group" binding:"required"`
}

type getMaxLiftsByMuscleGroupResponse struct {
	MuscleGroup  string  `json:"muscleGroup"`
	ExerciseName string  `json:"exerciseName"`
	WeightLifted float64 `json:"weightLifted"`
	Reps         int32   `json:"reps"`
}

func (server *Server) getMaxLiftsByMuscleGroup(ctx *gin.Context) {
	userID := ctx.MustGet(authorizationPayloadKey).(*token.Payload).UserID

	req := getMaxLiftsByMuscleGroupRequest{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	muscleGroup, err := server.store.GetMuscleGroupByName(ctx, req.MuscleGroup)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf(MUSCLE_GROUP_NOT_FOUND, req.MuscleGroup)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	lifts, err := server.store.GetMaxLiftsByMuscleGroup(ctx, db.GetMaxLiftsByMuscleGroupParams{MuscleGroup: muscleGroup.Name, UserID: userID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := make([]getMaxLiftsByMuscleGroupResponse, len(lifts))
	for i, v := range lifts {
		res[i] = getMaxLiftsByMuscleGroupResponse{
			MuscleGroup:  v.MuscleGroup,
			Reps:         v.Reps,
			ExerciseName: v.ExerciseName,
			WeightLifted: v.WeightLifted,
		}
	}

	ctx.JSON(http.StatusOK, res)
}

type getMaxRepLiftsRequest struct {
	Limit int32 `uri:"limit"`
}

func (server *Server) getMaxRepLifts(ctx *gin.Context) {
	userID := ctx.MustGet(authorizationPayloadKey).(*token.Payload).UserID

	req := getMaxRepLiftsRequest{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	lifts, err := server.store.GetMaxRepLifts(ctx, db.GetMaxRepLiftsParams{UserID: userID, Limit: req.Limit})
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

type updateLiftRequest struct {
	ExerciseName string  `json:"exerciseName"`
	WeightLifted float64 `json:"weightLifted"`
	Reps         int32   `json:"reps"`
	SetType      string  `json:"setType"`
	ID           int64   `json:"id" binding:"required"`
}

func (server *Server) updateLift(ctx *gin.Context) {
	userID := ctx.MustGet(authorizationPayloadKey).(*token.Payload).UserID

	req := updateLiftRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	query, err := server.store.GetLift(ctx, db.GetLiftParams{ID: req.ID, UserID: userID})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf(LIFT_NOT_FOUND, req.ID)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	args := db.UpdateLiftParams{
		ExerciseName: sql.NullString{String: req.ExerciseName, Valid: req.ExerciseName != ""},
		WeightLifted: sql.NullFloat64{Float64: req.WeightLifted, Valid: req.WeightLifted > 0},
		SetType:      sql.NullString{String: req.SetType, Valid: req.SetType != ""},
		Reps:         sql.NullInt32{Int32: req.Reps, Valid: req.Reps > 0},
		ID:           query.ID,
		UserID:       userID,
	}

	err = server.store.UpdateLift(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	lift, err := server.store.GetLift(ctx, db.GetLiftParams{ID: query.ID, UserID: userID})
	if err != nil {
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

type deleteLiftRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) deleteLift(ctx *gin.Context) {
	userID := ctx.MustGet(authorizationPayloadKey).(*token.Payload).UserID

	req := deleteLiftRequest{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	query, err := server.store.GetLift(ctx, db.GetLiftParams{ID: req.ID, UserID: userID})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf(LIFT_NOT_FOUND, req.ID)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.DeleteLift(ctx, db.DeleteLiftParams{ID: query.ID, UserID: userID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
