package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mockdb "github.com/AntoninoAdornetto/mogged-lift-tracker-service/db/mock"
	db "github.com/AntoninoAdornetto/mogged-lift-tracker-service/db/sqlc"
	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateProfile(t *testing.T) {
	userID := uuid.New()
	profile := GenRandProfile(userID)

	testCases := []struct {
		Name       string
		Body       gin.H
		buildStubs func(store *mockdb.MockStore)
		checkRes   func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			Name: "Bad Request",
			Body: gin.H{},
			buildStubs: func(store *mockdb.MockStore) {
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			Name: "Internal Error",
			Body: gin.H{
				"country":           profile.Country,
				"measurementSystem": profile.MeasurementSystem,
				"bodyWeight":        profile.BodyWeight,
				"bodyFat":           profile.BodyFat,
				"timeZoneOffset":    profile.TimezoneOffset,
				"userID":            userID.String(),
			},
			buildStubs: func(store *mockdb.MockStore) {
				args := db.CreateProfileParams{
					Country:           profile.Country,
					MeasurementSystem: profile.MeasurementSystem,
					BodyWeight:        profile.BodyWeight,
					BodyFat:           profile.BodyFat,
					TimezoneOffset:    profile.TimezoneOffset,
					UserID:            userID.String(),
				}
				store.EXPECT().CreateProfile(gomock.Any(), gomock.Eq(args)).Times(1).Return(int64(0), sql.ErrConnDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			Name: "Not Found -> Get Profile",
			Body: gin.H{
				"country":           profile.Country,
				"measurementSystem": profile.MeasurementSystem,
				"bodyWeight":        profile.BodyWeight,
				"bodyFat":           profile.BodyFat,
				"timeZoneOffset":    profile.TimezoneOffset,
				"userID":            userID.String(),
			},
			buildStubs: func(store *mockdb.MockStore) {
				args := db.CreateProfileParams{
					Country:           profile.Country,
					MeasurementSystem: profile.MeasurementSystem,
					BodyWeight:        profile.BodyWeight,
					BodyFat:           profile.BodyFat,
					TimezoneOffset:    profile.TimezoneOffset,
					UserID:            userID.String(),
				}
				store.EXPECT().CreateProfile(gomock.Any(), gomock.Eq(args)).Times(1).Return(int64(profile.ID), nil)
				store.EXPECT().GetProfile(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(db.Profile{}, sql.ErrNoRows)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			Name: "Internal Error -> Get Profile",
			Body: gin.H{
				"country":           profile.Country,
				"measurementSystem": profile.MeasurementSystem,
				"bodyWeight":        profile.BodyWeight,
				"bodyFat":           profile.BodyFat,
				"timeZoneOffset":    profile.TimezoneOffset,
				"userID":            userID.String(),
			},
			buildStubs: func(store *mockdb.MockStore) {
				args := db.CreateProfileParams{
					Country:           profile.Country,
					MeasurementSystem: profile.MeasurementSystem,
					BodyWeight:        profile.BodyWeight,
					BodyFat:           profile.BodyFat,
					TimezoneOffset:    profile.TimezoneOffset,
					UserID:            userID.String(),
				}
				store.EXPECT().CreateProfile(gomock.Any(), gomock.Eq(args)).Times(1).Return(int64(profile.ID), nil)
				store.EXPECT().GetProfile(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(db.Profile{}, sql.ErrConnDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			Name: "Internal Error -> Get Profile",
			Body: gin.H{
				"country":           profile.Country,
				"measurementSystem": profile.MeasurementSystem,
				"bodyWeight":        profile.BodyWeight,
				"bodyFat":           profile.BodyFat,
				"timeZoneOffset":    profile.TimezoneOffset,
				"userID":            userID.String(),
			},
			buildStubs: func(store *mockdb.MockStore) {
				args := db.CreateProfileParams{
					Country:           profile.Country,
					MeasurementSystem: profile.MeasurementSystem,
					BodyWeight:        profile.BodyWeight,
					BodyFat:           profile.BodyFat,
					TimezoneOffset:    profile.TimezoneOffset,
					UserID:            userID.String(),
				}
				store.EXPECT().CreateProfile(gomock.Any(), gomock.Eq(args)).Times(1).Return(int64(profile.ID), nil)
				store.EXPECT().GetProfile(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(profile, nil)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				validateNewProfileResponse(t, recorder.Body, ProfileResponse{
					Country:           profile.Country,
					MeasurementSystem: profile.MeasurementSystem,
					BodyWeight:        profile.BodyWeight,
					BodyFat:           profile.BodyFat,
					TimeZoneOffset:    profile.TimezoneOffset,
				})
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

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			payload, err := json.Marshal(tc.Body)
			require.NoError(t, err)

			url := "/createProfile"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkRes(t, recorder)
		})
	}
}

func GenRandProfile(userID uuid.UUID) db.Profile {
	return db.Profile{
		Country:           strings.ToUpper(util.RandomStr(3)),
		ID:                int32(util.RandomInt(1, 100)),
		UserID:            []byte(userID.String()),
		MeasurementSystem: "Imperial",
		BodyWeight:        float64(util.RandomInt(150, 220)),
		BodyFat:           float64(util.RandomInt(15, 20)),
		TimezoneOffset:    240,
	}
}

func validateNewProfileResponse(t *testing.T, body *bytes.Buffer, res ProfileResponse) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var profile ProfileResponse
	err = json.Unmarshal(data, &profile)
	require.NoError(t, err)
	require.Equal(t, profile, res)
}
