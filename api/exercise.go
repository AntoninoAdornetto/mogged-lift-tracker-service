package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/AntoninoAdornetto/mogged-lift-tracker-service/db/sqlc"
	"github.com/gin-gonic/gin"
)

const (
	EXERCISE_NOT_FOUND = "exercise with specified ID '%d' does not exist"
)

type ExerciseResponse struct {
	ID               int32   `json:"ID"`
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
