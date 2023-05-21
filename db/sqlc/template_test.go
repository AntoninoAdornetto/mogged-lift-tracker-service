package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// asserts both create and get one queries
func TestCreateTemplate(t *testing.T) {
	user := GenRandUser(t)
	GenRandTemplate(t, user.ID)
}

func TestListTemplates(t *testing.T) {
	user := GenRandUser(t)
	n := 5
	templates := make([]Template, n)

	query, err := testQueries.ListTemplates(context.Background(), user.ID)
	require.NoError(t, err)
	require.Empty(t, query)

	for i := 0; i < n; i++ {
		templates[i] = GenRandTemplate(t, user.ID)
	}

	query, err = testQueries.ListTemplates(context.Background(), user.ID)
	require.NoError(t, err)
	require.Len(t, query, n)
	for i, v := range query {
		require.Equal(t, v.Name, templates[i].Name)
		require.Equal(t, v.Exercises, templates[i].Exercises)
		require.WithinDuration(t, v.DateLastUsed, templates[i].DateLastUsed, time.Second)
		require.Equal(t, v.ID, templates[i].ID)
		require.Equal(t, v.Exercises, templates[i].Exercises)
	}
}

func TestUpdateTemplateName(t *testing.T) {
	user := GenRandUser(t)
	template := GenRandTemplate(t, user.ID)
	newTemplateName := sql.NullString{String: util.RandomStr(5), Valid: true}

	err := testQueries.UpdateTemplate(context.Background(), UpdateTemplateParams{Name: newTemplateName, ID: template.ID, CreatedBy: user.ID})
	require.NoError(t, err)

	query, err := testQueries.GetTemplate(context.Background(), GetTemplateParams{ID: template.ID, CreatedBy: user.ID})
	require.NoError(t, err)
	require.NotZero(t, query.ID)
	require.Equal(t, query.Name, newTemplateName.String)
	require.NotEqual(t, query.Name, template.Name)
}

func TestUpdateTemplateExercises(t *testing.T) {
	user := GenRandUser(t)
	template := GenRandTemplate(t, user.ID)
	newExercises := GenRandTemplate(t, user.ID).Exercises

	err := testQueries.UpdateTemplate(context.Background(), UpdateTemplateParams{Exercises: newExercises, CreatedBy: user.ID, ID: template.ID})
	require.NoError(t, err)

	query, err := testQueries.GetTemplate(context.Background(), GetTemplateParams{ID: template.ID, CreatedBy: user.ID})
	require.NoError(t, err)
	require.NotZero(t, query.ID)
	require.Equal(t, query.Exercises, newExercises)
	require.NotEqual(t, query.Exercises, template.Exercises)
}

func TestUpdateTemplatesDateLastUsed(t *testing.T) {
	user := GenRandUser(t)
	template := GenRandTemplate(t, user.ID)
	newDateLastUsed := time.Now()

	err := testQueries.UpdateTemplate(context.Background(), UpdateTemplateParams{
		ID: template.ID,
		DateLastUsed: sql.NullTime{
			Time:  newDateLastUsed,
			Valid: true,
		},
		CreatedBy: user.ID,
	})
	require.NoError(t, err)

	query, err := testQueries.GetTemplate(context.Background(), GetTemplateParams{ID: template.ID, CreatedBy: user.ID})
	require.NoError(t, err)
	require.NotZero(t, query.ID)
	require.NotEqual(t, query.DateLastUsed, template.DateLastUsed)
	require.WithinDuration(t, query.DateLastUsed, newDateLastUsed, time.Hour*24)
}

func TestDeleteTemplate(t *testing.T) {
	user := GenRandUser(t)
	template := GenRandTemplate(t, user.ID)

	err := testQueries.DeleteTemplate(context.Background(), DeleteTemplateParams{
		ID:        template.ID,
		CreatedBy: user.ID,
	})
	require.NoError(t, err)

	query, err := testQueries.GetTemplate(context.Background(), GetTemplateParams{
		ID:        template.ID,
		CreatedBy: user.ID,
	})
	require.Error(t, err)
	require.Zero(t, query.ID)
}

func GenRandTemplate(t *testing.T, userId string) Template {
	template := &Template{}
	workout := GenRandWorkout(t, userId)

	args := CreateTemplateParams{
		Name:      util.RandomStr(8),
		Exercises: workout.Lifts,
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
	require.Equal(t, query.Exercises, args.Exercises)

	template = &query
	return *template
}
