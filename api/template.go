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
	ID           int32           `json:"id"`
	Name         string          `json:"name"`
	Exercises    json.RawMessage `json:"exercises"`
	DateLastUsed string          `json:"dateLastUsed"`
	CreatedBy    string          `json:"createdBy"`
}

func templateResponse(t db.Template, userID string) TemplateResponse {
	return TemplateResponse{
		ID:           t.ID,
		Name:         t.Name,
		Exercises:    t.Exercises,
		DateLastUsed: t.DateLastUsed.Format("2006-01-02"),
		CreatedBy:    userID,
	}
}

type createTemplateRequest struct {
	Name      string          `json:"name" binding:"required"`
	Exercises json.RawMessage `json:"exercises" binding:"required"`
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
		Exercises: req.Exercises,
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

	ctx.JSON(http.StatusOK, templateResponse(template, userID))
}

type getTemplateRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getTemplate(ctx *gin.Context) {
	userID := ctx.MustGet(authorizationPayloadKey).(*token.Payload).UserID

	req := getTemplateRequest{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	template, err := server.store.GetTemplate(ctx, db.GetTemplateParams{ID: req.ID, CreatedBy: userID})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf(TEMPLATE_NOT_FOUND, req.ID)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, templateResponse(template, userID))
}

func (server *Server) listTemplates(ctx *gin.Context) {
	userID := ctx.MustGet(authorizationPayloadKey).(*token.Payload).UserID

	list, err := server.store.ListTemplates(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	templates := make([]TemplateResponse, len(list))
	for i, template := range list {
		templates[i] = templateResponse(template, userID)
	}

	ctx.JSON(http.StatusOK, templates)
}

type updateTemplateRequest struct {
	ID           int32           `json:"id" binding:"required"`
	Name         string          `json:"name"`
	Exercises    json.RawMessage `json:"exercises"`
	DateLastUsed string          `json:"dateLastUsed"`
	CreatedBy    string          `json:"createdBy" binding:"required"`
}

func (server *Server) updateTemplate(ctx *gin.Context) {
	userID := ctx.MustGet(authorizationPayloadKey).(*token.Payload).UserID

	req := updateTemplateRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	query, err := server.store.GetTemplate(ctx, db.GetTemplateParams{ID: req.ID, CreatedBy: userID})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNoContent, errorResponse(fmt.Errorf(TEMPLATE_NOT_FOUND, req.ID)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var t time.Time
	if req.DateLastUsed != "" {
		t, err = time.Parse("2006-01-02", req.DateLastUsed)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("incorrect date format for value: %s. Should be YYYY-MM-DD", req.DateLastUsed)))
			return
		}
	}

	args := db.UpdateTemplateParams{
		Name:         sql.NullString{String: req.Name, Valid: req.Name != ""},
		DateLastUsed: sql.NullTime{Time: t, Valid: !t.IsZero()},
		Exercises:    req.Exercises,
		ID:           query.ID,
		CreatedBy:    userID,
	}

	err = server.store.UpdateTemplate(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	template, err := server.store.GetTemplate(ctx, db.GetTemplateParams{ID: query.ID, CreatedBy: userID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, templateResponse(template, userID))
}

type deleteTemplateRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) deleteTemplate(ctx *gin.Context) {
	userID := ctx.MustGet(authorizationPayloadKey).(*token.Payload).UserID

	req := deleteTemplateRequest{}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	query, err := server.store.GetTemplate(ctx, db.GetTemplateParams{ID: req.ID, CreatedBy: userID})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf(TEMPLATE_NOT_FOUND, req.ID)))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.DeleteTemplate(ctx, db.DeleteTemplateParams{ID: query.ID, CreatedBy: userID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
