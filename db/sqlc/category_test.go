package db

import (
	"context"
	"testing"

	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/stretchr/testify/require"
)

// asserts both create and get one queries
func TestCreateCategory(t *testing.T) {
	GenRandCategory(t)
}

func TestListCategories(t *testing.T) {
	n := 5
	categories := make([]Category, n)

	for i := 0; i < n; i++ {
		categories[i] = GenRandCategory(t)
	}

	query, err := testQueries.ListCategories(context.Background())
	require.NoError(t, err)

	for _, v := range query {
		require.NotEmpty(t, v.ID)
		require.NotEmpty(t, v.Name)
	}
}

func TestUpdateCategoryName(t *testing.T) {
	category := GenRandCategory(t)
	newCategoryName := util.RandomStr(5)

	_, err := testQueries.UpdateCategory(context.Background(), UpdateCategoryParams{
		ID:   category.ID,
		Name: newCategoryName,
	})
	require.NoError(t, err)

	query, err := testQueries.GetCategory(context.Background(), category.ID)
	require.NoError(t, err)
	require.NotZero(t, query.ID)
	require.NotEqual(t, query.Name, category.Name)
	require.Equal(t, query.Name, newCategoryName)
	require.Equal(t, query.ID, category.ID)
}

func TestDeleteCategory(t *testing.T) {
	category := GenRandCategory(t)

	_, err := testQueries.DeleteCategory(context.Background(), category.ID)
	require.NoError(t, err)

	query, err := testQueries.GetCategory(context.Background(), category.ID)
	require.Error(t, err)
	require.Empty(t, query.ID)
	require.Empty(t, query.Name)
}

func GenRandCategory(t *testing.T) Category {
	categoryName := util.RandomStr(5)
	record, err := testQueries.CreateCategory(context.Background(), categoryName)
	require.NoError(t, err)

	id, err := record.LastInsertId()
	require.NoError(t, err)

	category, err := testQueries.GetCategory(context.Background(), int32(id))
	require.NoError(t, err)
	require.NotZero(t, category.ID)
	require.Equal(t, category.Name, categoryName)
	require.Equal(t, category.ID, int32(id))

	return category
}
