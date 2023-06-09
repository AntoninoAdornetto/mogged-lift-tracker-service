package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/AntoninoAdornetto/mogged-lift-tracker-service/db/mock"
	db "github.com/AntoninoAdornetto/mogged-lift-tracker-service/db/sqlc"
	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/token"
	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateExercise(t *testing.T) {
	userID := uuid.New()
	exercise := GenRandExercise(userID.String())
	exerciseRes := ExerciseResponse{
		Name:             exercise.Name,
		ID:               exercise.ID,
		MuscleGroup:      exercise.MuscleGroup,
		Category:         exercise.Category,
		IsStock:          exercise.Isstock,
		MostWeightLifted: exercise.MostWeightLifted,
		MostRepsLifted:   exercise.MostRepsLifted,
		RestTimer:        exercise.RestTimer,
		UserID:           userID.String(),
	}

	testCases := []struct {
		Name       string
		Body       gin.H
		setupAuth  func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs func(store *mockdb.MockStore)
		checkRes   func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			Name: "Bad Request",
			Body: gin.H{},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(
					t,
					request,
					tokenMaker,
					authorizationBearerType,
					userID.String(),
					time.Minute,
				)
			},
			buildStubs: func(store *mockdb.MockStore) {},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			Name: "Internal Error",
			Body: gin.H{
				"exerciseName": exercise.Name,
				"muscleGroup":  exercise.MuscleGroup,
				"category":     exercise.Category,
				"userID":       userID.String(),
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(
					t,
					request,
					tokenMaker,
					authorizationBearerType,
					userID.String(),
					time.Minute,
				)
			},
			buildStubs: func(store *mockdb.MockStore) {
				args := db.CreateExerciseParams{
					Name:        exercise.Name,
					MuscleGroup: exercise.MuscleGroup,
					Category:    exercise.Category,
					UserID:      userID.String(),
				}
				store.EXPECT().
					CreateExercise(gomock.Any(), gomock.Eq(args)).
					Times(1).
					Return(int64(0), sql.ErrConnDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			Name: "Not Found -> Get Exercise",
			Body: gin.H{
				"exerciseName": exercise.Name,
				"muscleGroup":  exercise.MuscleGroup,
				"category":     exercise.Category,
				"userID":       userID.String(),
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(
					t,
					request,
					tokenMaker,
					authorizationBearerType,
					userID.String(),
					time.Minute,
				)
			},
			buildStubs: func(store *mockdb.MockStore) {
				args := db.CreateExerciseParams{
					Name:        exercise.Name,
					MuscleGroup: exercise.MuscleGroup,
					Category:    exercise.Category,
					UserID:      userID.String(),
				}
				store.EXPECT().
					CreateExercise(gomock.Any(), gomock.Eq(args)).
					Times(1).
					Return(int64(exercise.ID), nil)
				store.EXPECT().
					GetExercise(gomock.Any(), gomock.Eq(db.GetExerciseParams{ID: exercise.ID, UserID: userID.String()})).
					Times(1).
					Return(db.Exercise{}, sql.ErrNoRows)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			Name: "Internal Error -> Get Exercise",
			Body: gin.H{
				"exerciseName": exercise.Name,
				"muscleGroup":  exercise.MuscleGroup,
				"category":     exercise.Category,
				"userID":       userID.String(),
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(
					t,
					request,
					tokenMaker,
					authorizationBearerType,
					userID.String(),
					time.Minute,
				)
			},
			buildStubs: func(store *mockdb.MockStore) {
				args := db.CreateExerciseParams{
					Name:        exercise.Name,
					MuscleGroup: exercise.MuscleGroup,
					Category:    exercise.Category,
					UserID:      userID.String(),
				}
				store.EXPECT().
					CreateExercise(gomock.Any(), gomock.Eq(args)).
					Times(1).
					Return(int64(exercise.ID), nil)
				store.EXPECT().
					GetExercise(gomock.Any(), gomock.Eq(db.GetExerciseParams{ID: exercise.ID, UserID: userID.String()})).
					Times(1).
					Return(db.Exercise{}, sql.ErrConnDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			Name: "OK",
			Body: gin.H{
				"exerciseName": exercise.Name,
				"muscleGroup":  exercise.MuscleGroup,
				"category":     exercise.Category,
				"userID":       userID.String(),
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(
					t,
					request,
					tokenMaker,
					authorizationBearerType,
					userID.String(),
					time.Minute,
				)
			},
			buildStubs: func(store *mockdb.MockStore) {
				args := db.CreateExerciseParams{
					Name:        exercise.Name,
					MuscleGroup: exercise.MuscleGroup,
					Category:    exercise.Category,
					UserID:      userID.String(),
				}
				store.EXPECT().
					CreateExercise(gomock.Any(), gomock.Eq(args)).
					Times(1).
					Return(int64(exercise.ID), nil)
				store.EXPECT().
					GetExercise(gomock.Any(), gomock.Eq(db.GetExerciseParams{ID: exercise.ID, UserID: userID.String()})).
					Times(1).
					Return(exercise, nil)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				validateExerciseResponse(t, recorder.Body, exerciseRes)
			},
		},
		{
			Name: "Unauthorized",
			Body: gin.H{
				"exerciseName": exercise.Name,
				"muscleGroup":  exercise.MuscleGroup,
				"category":     exercise.Category,
				"userID":       userID.String(),
			},
			setupAuth:  func(t *testing.T, request *http.Request, tokenMaker token.Maker) {},
			buildStubs: func(store *mockdb.MockStore) {},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			payload, err := json.Marshal(tc.Body)
			require.NoError(t, err)

			url := "/api/createExercise"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkRes(t, recorder)
		})
	}
}

func TestGetExercise(t *testing.T) {
	userID := uuid.New()
	_ = GenRandUser(userID)
	exercise := GenRandExercise(userID.String())

	testCases := []struct {
		Name       string
		id         int32
		setupAuth  func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs func(store *mockdb.MockStore)
		checkRes   func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			Name: "Bad Request",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(
					t,
					request,
					tokenMaker,
					authorizationBearerType,
					userID.String(),
					time.Minute,
				)
			},
			buildStubs: func(store *mockdb.MockStore) {
				args := db.GetExerciseParams{
					ID: exercise.ID,
				}
				store.EXPECT().
					GetExercise(gomock.Any(), gomock.Eq(args)).
					Times(0).
					Return(db.Exercise{}, sql.ErrConnDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			Name: "Not Found",
			id:   1000,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(
					t,
					request,
					tokenMaker,
					authorizationBearerType,
					userID.String(),
					time.Minute,
				)
			},
			buildStubs: func(store *mockdb.MockStore) {
				args := db.GetExerciseParams{
					ID:     1000,
					UserID: userID.String(),
				}
				store.EXPECT().
					GetExercise(gomock.Any(), gomock.Eq(args)).
					Times(1).
					Return(db.Exercise{}, sql.ErrNoRows)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			Name: "Internal Server Error",
			id:   1000,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(
					t,
					request,
					tokenMaker,
					authorizationBearerType,
					userID.String(),
					time.Minute,
				)
			},
			buildStubs: func(store *mockdb.MockStore) {
				args := db.GetExerciseParams{
					ID:     1000,
					UserID: userID.String(),
				}
				store.EXPECT().
					GetExercise(gomock.Any(), gomock.Eq(args)).
					Times(1).
					Return(db.Exercise{}, sql.ErrConnDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			Name: "OK",
			id:   exercise.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(
					t,
					request,
					tokenMaker,
					authorizationBearerType,
					userID.String(),
					time.Minute,
				)
			},
			buildStubs: func(store *mockdb.MockStore) {
				args := db.GetExerciseParams{
					ID:     exercise.ID,
					UserID: userID.String(),
				}
				store.EXPECT().
					GetExercise(gomock.Any(), gomock.Eq(args)).
					Times(1).
					Return(db.Exercise{
						ID:               exercise.ID,
						Category:         exercise.Category,
						MostRepsLifted:   exercise.MostRepsLifted,
						MostWeightLifted: exercise.MostWeightLifted,
						MuscleGroup:      exercise.MuscleGroup,
						Name:             exercise.Name,
						RestTimer:        exercise.RestTimer,
						UserID:           []byte(userID.String()),
						Isstock:          exercise.Isstock,
					}, nil)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/api/getExercise/%d", tc.id)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkRes(t, recorder)
		})
	}
}

func TestGetExerciseByName(t *testing.T) {
	userID := uuid.New()
	_ = GenRandUser(userID)
	exercise := GenRandExercise(userID.String())

	testCases := []struct {
		Name         string
		ExerciseName string
		setupAuth    func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs   func(store *mockdb.MockStore)
		checkRes     func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			Name:         "Not Found",
			ExerciseName: exercise.Name,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(
					t,
					request,
					tokenMaker,
					authorizationBearerType,
					userID.String(),
					time.Minute,
				)
			},
			buildStubs: func(store *mockdb.MockStore) {
				args := db.GetExerciseByNameParams{
					Name:   exercise.Name,
					UserID: userID.String(),
				}
				store.EXPECT().
					GetExerciseByName(gomock.Any(), gomock.Eq(args)).
					Times(1).
					Return(db.Exercise{}, sql.ErrNoRows)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			Name:         "Internal Error",
			ExerciseName: exercise.Name,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(
					t,
					request,
					tokenMaker,
					authorizationBearerType,
					userID.String(),
					time.Minute,
				)
			},
			buildStubs: func(store *mockdb.MockStore) {
				args := db.GetExerciseByNameParams{
					Name:   exercise.Name,
					UserID: userID.String(),
				}
				store.EXPECT().
					GetExerciseByName(gomock.Any(), gomock.Eq(args)).
					Times(1).
					Return(db.Exercise{}, sql.ErrConnDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			Name:         "Unauthorized",
			ExerciseName: exercise.Name,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			},
			buildStubs: func(store *mockdb.MockStore) {
				args := db.GetExerciseByNameParams{
					Name:   exercise.Name,
					UserID: userID.String(),
				}
				store.EXPECT().
					GetExerciseByName(gomock.Any(), gomock.Eq(args)).
					Times(0)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			Name:         "OK",
			ExerciseName: exercise.Name,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(
					t,
					request,
					tokenMaker,
					authorizationBearerType,
					userID.String(),
					time.Minute,
				)
			},
			buildStubs: func(store *mockdb.MockStore) {
				args := db.GetExerciseByNameParams{
					Name:   exercise.Name,
					UserID: userID.String(),
				}
				store.EXPECT().
					GetExerciseByName(gomock.Any(), gomock.Eq(args)).
					Times(1).
					Return(db.Exercise{
						ID:               exercise.ID,
						Name:             exercise.Name,
						MuscleGroup:      exercise.MuscleGroup,
						Category:         exercise.Category,
						Isstock:          exercise.Isstock,
						MostWeightLifted: exercise.MostWeightLifted,
						MostRepsLifted:   exercise.MostRepsLifted,
						RestTimer:        exercise.RestTimer,
						UserID:           []byte(userID.String()),
					}, nil)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/api/getByName/%s", tc.ExerciseName)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkRes(t, recorder)
		})
	}
}

func GenRandExercise(userID string) db.Exercise {
	return db.Exercise{
		ID:               int32(util.RandomInt(1, 10)),
		Name:             util.RandomStr(5),
		MuscleGroup:      util.RandomStr(5),
		Category:         util.RandomStr(5),
		Isstock:          false,
		MostWeightLifted: float64(util.RandomInt(100, 200)),
		MostRepsLifted:   int32(util.RandomInt(6, 25)),
		RestTimer:        "00:02:20s",
		UserID:           []byte(userID),
	}
}

func validateExerciseResponse(t *testing.T, body *bytes.Buffer, res ExerciseResponse) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var exercise ExerciseResponse
	err = json.Unmarshal(data, &exercise)
	require.NoError(t, err)

	require.Equal(t, exercise, res)
}
