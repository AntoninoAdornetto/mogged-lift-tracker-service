package db

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// asserts both create and get one queries
func TestCreateTemplate(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)
	GenRandTemplate(t, userId.String())
}

func TestListTemplates(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)
	n := 5
	templates := make([]Template, n)

	query, err := testQueries.ListTemplates(context.Background(), userId.String())
	require.NoError(t, err)
	require.Empty(t, query)

	for i := 0; i < n; i++ {
		templates[i] = GenRandTemplate(t, userId.String())
	}

	query, err = testQueries.ListTemplates(context.Background(), userId.String())
	require.NoError(t, err)
	require.Len(t, query, n)
	for i, v := range query {
		require.Equal(t, v.Name, templates[i].Name)
		require.Equal(t, v.Lifts, templates[i].Lifts)
		require.WithinDuration(t, v.DateLastUsed, templates[i].DateLastUsed, time.Second)
		require.Equal(t, v.ID, templates[i].ID)
		require.Equal(t, v.Lifts, templates[i].Lifts)
	}
}

// @todo - assert for date updates
func TestUpdateTemplate(t *testing.T) {
	user := GenRandUser(t)
	userId, err := uuid.Parse(user.ID)
	require.NoError(t, err)

	template := GenRandTemplate(t, userId.String())
	newTemplateValues := struct {
		NewTemplateName string
		NewLifts        json.RawMessage
		NewDateLastUsed int64
	}{
		NewTemplateName: util.RandomStr(5),
		NewLifts:        GenRandTemplate(t, userId.String()).Lifts,
	}

	_, err = testQueries.UpdateTemplate(context.Background(), UpdateTemplateParams{
		Name:      newTemplateValues.NewTemplateName,
		CreatedBy: userId.String(),
		ID:        template.ID,
		Lifts:     newTemplateValues.NewLifts,
		// DateLastUsed: time.UnixMilli(newTemplateValues.NewDateLastUsed),
	})
	require.NoError(t, err)

	query, err := testQueries.GetTemplate(context.Background(), GetTemplateParams{
		ID:        template.ID,
		CreatedBy: userId.String(),
	})
	require.NoError(t, err)
	require.Equal(t, query.ID, template.ID)
	require.Equal(t, query.Lifts, newTemplateValues.NewLifts)
	require.Equal(t, query.Name, newTemplateValues.NewTemplateName)
}

func GenRandTemplate(t *testing.T, userId string) Template {
	template := &Template{}
	workout := GenRandWorkout(t, userId)

	args := CreateTemplateParams{
		Name:      util.RandomStr(8),
		Lifts:     workout.Lifts,
		CreatedBy: userId,
	}

	record, err := testQueries.CreateTemplate(context.Background(), args)
	require.NoError(t, err)
	id, err := record.LastInsertId()
	require.NoError(t, err)
	require.NotZero(t, id)

	query, err := testQueries.GetTemplate(context.Background(), GetTemplateParams{
		ID:        int32(id),
		CreatedBy: userId,
	})
	require.NoError(t, err)
	createdBy, err := uuid.FromBytes(query.CreatedBy)
	require.NoError(t, err)
	require.Equal(t, query.ID, int32(id))
	require.Equal(t, createdBy.String(), userId)
	require.Equal(t, query.Name, args.Name)
	require.Equal(t, query.Lifts, args.Lifts)

	template = &query
	return *template
}
