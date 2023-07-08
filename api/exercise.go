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

func exerciseRes(ex db.Exercise, id string) ExerciseResponse {
	return ExerciseResponse{
		ID:               ex.ID,
		Category:         ex.Category,
		MuscleGroup:      ex.MuscleGroup,
		IsStock:          ex.Isstock,
		MostRepsLifted:   ex.MostRepsLifted,
		MostWeightLifted: ex.MostWeightLifted,
		RestTimer:        ex.RestTimer,
		UserID:           id,
		Name:             ex.Name,
	}
}

type createExerciseRequest struct {
	Name        string `json:"exerciseName" binding:"required"`
	MuscleGroup string `json:"muscleGroup"  binding:"required"`
	Category    string `json:"category"     binding:"required"`
}

func (server *Server) createExercise(ctx *gin.Context) {
	req := createExerciseRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID := ctx.MustGet(authorizationPayloadKey).(*token.Payload).UserID

	args := db.CreateExerciseParams{
		Name:        req.Name,
		MuscleGroup: req.MuscleGroup,
		Category:    req.Category,
		UserID:      userID,
	}

	exerciseID, err := server.store.CreateExercise(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	exercise, err := server.store.GetExercise(
		ctx,
		db.GetExerciseParams{ID: int32(exerciseID), UserID: userID},
	)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf(EXERCISE_NOT_FOUND, exerciseID)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, exerciseRes(exercise, userID))
}

type getExerciseRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getExercise(ctx *gin.Context) {
	req := getExerciseRequest{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID := ctx.MustGet(authorizationPayloadKey).(*token.Payload).UserID

	args := db.GetExerciseParams{
		ID:     req.ID,
		UserID: userID,
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

	ctx.JSON(http.StatusOK, exerciseRes(exercise, userID))
}

type getExerciseByNameRequest struct {
	Name string `uri:"exercise" binding:"required"`
}

func (server *Server) getExerciseByName(ctx *gin.Context) {
	req := getExerciseByNameRequest{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID := ctx.MustGet(authorizationPayloadKey).(*token.Payload).UserID

	args := db.GetExerciseByNameParams{
		Name:   req.Name,
		UserID: userID,
	}

	exercise, err := server.store.GetExerciseByName(ctx, args)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(
				http.StatusNotFound,
				errorResponse(fmt.Errorf(EXERCISE_NAME_NOT_FOUND, req.Name)),
			)
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, exerciseRes(exercise, userID))
}

// @TODO: add Server Side Pagination/Sorting
func (server *Server) listExercises(ctx *gin.Context) {
	userID := ctx.MustGet(authorizationPayloadKey).(*token.Payload).UserID
	exercises, err := server.store.ListExercises(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//@TODO: user_id returned as binary
	ctx.JSON(http.StatusOK, exercises)
}

type updateExerciseRequest struct {
	ID          int32  `json:"id"           binding:"required"`
	Name        string `json:"exerciseName"`
	MuscleGroup string `json:"muscleGroup"`
	Category    string `json:"category"`
	RestTimer   string `json:"restTimer"`
}

func (server *Server) updateExercise(ctx *gin.Context) {
	req := updateExerciseRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID := ctx.MustGet(authorizationPayloadKey).(*token.Payload).UserID
	getExerciseArgs := db.GetExerciseParams{ID: req.ID, UserID: userID}

	_, err := server.store.GetExercise(ctx, getExerciseArgs)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf(EXERCISE_NOT_FOUND, req.ID)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//@TODO - reps & weight should be updated through WO/TX
	args := db.UpdateExerciseParams{
		Name: sql.NullString{
			String: req.Name,
			Valid:  req.Name != "",
		},
		MuscleGroup: sql.NullString{
			String: req.MuscleGroup,
			Valid:  req.MuscleGroup != "",
		},
		Category: sql.NullString{
			String: req.Category,
			Valid:  req.Category != "",
		},
		RestTimer: sql.NullString{
			String: req.RestTimer,
			Valid:  req.RestTimer != "",
		},
		ID:     req.ID,
		UserID: userID,
	}

	err = server.store.UpdateExercise(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	exercise, err := server.store.GetExercise(ctx, getExerciseArgs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, exerciseRes(exercise, userID))
}

type deleteExerciseRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) deleteExercise(ctx *gin.Context) {
	req := deleteExerciseRequest{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID := ctx.MustGet(authorizationPayloadKey).(*token.Payload).UserID

	exercise, err := server.store.GetExercise(
		ctx,
		db.GetExerciseParams{ID: req.ID, UserID: userID},
	)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf(EXERCISE_NOT_FOUND, req.ID)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.DeleteExercise(
		ctx,
		db.DeleteExerciseParams{ID: exercise.ID, UserID: userID},
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
