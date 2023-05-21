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
	TEMPLATE_NOT_FOUND = "template with specified ID '%d' does not exist"
)

type TemplateResponse struct {
	ID           int32                `json:"id"`
	Name         string               `json:"name"`
	Lifts        map[string][]db.Lift `json:"lifts"`
	DateLastUsed time.Time            `json:"dateLastUsed"`
	CreatedBy    string               `json:"createdBy"`
}

type createTemplateRequest struct {
	Name      string          `json:"name" binding:"required"`
	Lifts     json.RawMessage `json:"lifts" binding:"required"`
	CreatedBy string          `json:"createdBy" binding:"required"`
}

func (server *Server) createTemplate(ctx *gin.Context) {
	userID := ctx.MustGet(authorizationPayloadKey).(*token.Payload).UserID

	req := createTemplateRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	record, err := server.store.CreateTemplate(ctx, db.CreateTemplateParams{
		Name:      req.Name,
		Lifts:     req.Lifts,
		CreatedBy: userID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	templateID, err := record.LastInsertId()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	template, err := server.store.GetTemplate(ctx, db.GetTemplateParams{ID: int32(templateID), CreatedBy: userID})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf(TEMPLATE_NOT_FOUND, templateID)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, template)
}
