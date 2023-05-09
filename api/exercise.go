package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/AntoninoAdornetto/mogged-lift-tracker-service/db/sqlc"
	"github.com/gin-gonic/gin"
)

const (
	EXERCISE_NOT_FOUND      = "exercise with specified ID '%d' does not exist"
	EXERCISE_NAME_NOT_FOUND = "exercise with specified name '%s' does not exist"
)

type ExerciseResponse struct {
	ID               int32   `json:"id"`
	Name             string  `json:"exerciseName"`
	MuscleGroup      string  `json:"muscleGroup"`
	Category         string  `json:"category"`
	IsStock          bool    `json:"isStockExercise"`
	MostWeightLifted float64 `json:"mostWeightLifted"`
	MostRepsLifted   int32   `json:"mostRepsLifted"`
	RestTimer        string  `json:"restTimer"`
	UserID           string  `json:"userID"`
}

type createExerciseRequest struct {
	Name        string `json:"exerciseName" binding:"required"`
	MuscleGroup string `json:"muscleGroup" binding:"required"`
	Category    string `json:"category" binding:"required"`
	UserID      string `json:"userID" binding:"required"`
}

func (server *Server) createExercise(ctx *gin.Context) {
	req := createExerciseRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.CreateExerciseParams{
		Name:        req.Name,
		MuscleGroup: req.MuscleGroup,
		Category:    req.Category,
		UserID:      req.UserID,
	}

	result, err := server.store.CreateExercise(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	exerciseID, err := result.LastInsertId()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	exercise, err := server.store.GetExercise(ctx, db.GetExerciseParams{ID: int32(exerciseID), UserID: req.UserID})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf(EXERCISE_NOT_FOUND, exerciseID)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ExerciseResponse{
		ID:               exercise.ID,
		Category:         exercise.Category,
		IsStock:          exercise.Isstock,
		MostRepsLifted:   exercise.MostRepsLifted,
		MostWeightLifted: exercise.MostWeightLifted,
		MuscleGroup:      exercise.MuscleGroup,
		Name:             exercise.Name,
		RestTimer:        exercise.RestTimer,
		UserID:           req.UserID,
	})
}

type getExerciseRequest struct {
	ID     int32  `uri:"id" binding:"required"`
	UserID string `uri:"user_id" binding:"required"`
}

func (server *Server) getExercise(ctx *gin.Context) {
	req := getExerciseRequest{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.GetExerciseParams{
		ID:     req.ID,
		UserID: req.UserID,
	}

	exercise, err := server.store.GetExercise(ctx, args)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf(EXERCISE_NOT_FOUND, req.ID)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ExerciseResponse{
		ID:               exercise.ID,
		Category:         exercise.Category,
		IsStock:          exercise.Isstock,
		MostRepsLifted:   exercise.MostRepsLifted,
		MostWeightLifted: exercise.MostWeightLifted,
		MuscleGroup:      exercise.MuscleGroup,
		Name:             exercise.Name,
		RestTimer:        exercise.RestTimer,
		UserID:           req.UserID,
	})
}

type getExerciseByNameRequest struct {
	Name   string `uri:"exercise_name" binding:"required"`
	UserID string `uri:"user_id" binding:"required"`
}

func (server *Server) getExerciseByName(ctx *gin.Context) {
	req := getExerciseByNameRequest{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.GetExerciseByNameParams{
		Name:   req.Name,
		UserID: req.UserID,
	}

	exercise, err := server.store.GetExerciseByName(ctx, args)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf(EXERCISE_NAME_NOT_FOUND, req.Name)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ExerciseResponse{
		ID:               exercise.ID,
		Category:         exercise.Category,
		IsStock:          exercise.Isstock,
		MostRepsLifted:   exercise.MostRepsLifted,
		MostWeightLifted: exercise.MostWeightLifted,
		MuscleGroup:      exercise.MuscleGroup,
		Name:             exercise.Name,
		RestTimer:        exercise.RestTimer,
		UserID:           req.UserID,
	})
}

type listExercisesRequest struct {
	UserID string `uri:"user_id" binding:"required"`
}

func (server *Server) listExercises(ctx *gin.Context) {
	req := listExercisesRequest{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	exercises, err := server.store.ListExercises(ctx, req.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, exercises)
}
