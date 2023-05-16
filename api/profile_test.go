package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestCreateProfile(t *testing.T) {
	userID := uuid.New()
	profile := GenRandProfile(userID)

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
				addAuthorization(t, request, tokenMaker, authorizationBearerType, userID.String(), time.Minute)
			},
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
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationBearerType, userID.String(), time.Minute)
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
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationBearerType, userID.String(), time.Minute)
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
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationBearerType, userID.String(), time.Minute)
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
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationBearerType, userID.String(), time.Minute)
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
		{
			Name: "Unauthorized",
			Body: gin.H{
				"country":           profile.Country,
				"measurementSystem": profile.MeasurementSystem,
				"bodyWeight":        profile.BodyWeight,
				"bodyFat":           profile.BodyFat,
				"timeZoneOffset":    profile.TimezoneOffset,
				"userID":            userID.String(),
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationBearerType, userID.String(), time.Minute)
			},
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

			url := "/createProfile"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkRes(t, recorder)
		})
	}
}

func TestGetProfile(t *testing.T) {
	userID := uuid.New()
	profile := GenRandProfile(userID)

	testCases := []struct {
		Name       string
		UserID     string
		setupAuth  func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs func(store *mockdb.MockStore)
		checkRes   func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			Name:   "Not Found",
			UserID: userID.String(),
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationBearerType, userID.String(), time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProfile(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(db.Profile{}, sql.ErrNoRows)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			Name:   "Internal Error",
			UserID: userID.String(),
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationBearerType, userID.String(), time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProfile(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(db.Profile{}, sql.ErrConnDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			Name:   "OK",
			UserID: userID.String(),
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationBearerType, userID.String(), time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
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
		{
			Name:       "Unauthorized",
			UserID:     userID.String(),
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

			url := fmt.Sprintf("/getProfile/%s", tc.UserID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkRes(t, recorder)
		})
	}
}

func TestUpdateProfile(t *testing.T) {
	userID := uuid.New()
	profile := GenRandProfile(userID)
	updateParams, updatedProfile := GenRandUpdateProfile(userID, profile.ID)

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
			Name: "Not Found -> Get Profile",
			Body: gin.H{
				"country":           "",
				"measurementSystem": "",
				"bodyWeight":        0,
				"bodyFat":           0,
				"timeZoneOffset":    -999,
				"userID":            userID.String(),
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProfile(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(db.Profile{}, sql.ErrNoRows)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			Name: "Internal Error -> Get Profile",
			Body: gin.H{
				"country":           "",
				"measurementSystem": "",
				"bodyWeight":        0,
				"bodyFat":           0,
				"timeZoneOffset":    -999,
				"userID":            userID.String(),
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProfile(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(db.Profile{}, sql.ErrConnDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			Name: "Internal Error -> Update Profile",
			Body: gin.H{
				"country":           updatedProfile.Country,
				"measurementSystem": updatedProfile.MeasurementSystem,
				"bodyWeight":        updatedProfile.BodyWeight,
				"bodyFat":           updatedProfile.BodyFat,
				"timeZoneOffset":    updatedProfile.TimezoneOffset,
				"userID":            userID.String(),
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProfile(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(profile, nil)
				store.EXPECT().UpdateProfile(gomock.Any(), gomock.Eq(updateParams)).Times(1).Return(nil, sql.ErrConnDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			Name: "Internal Error -> Get Profile After Update",
			Body: gin.H{
				"country":           updatedProfile.Country,
				"measurementSystem": updatedProfile.MeasurementSystem,
				"bodyWeight":        updatedProfile.BodyWeight,
				"bodyFat":           updatedProfile.BodyFat,
				"timeZoneOffset":    updatedProfile.TimezoneOffset,
				"userID":            userID.String(),
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProfile(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(profile, nil)
				store.EXPECT().UpdateProfile(gomock.Any(), gomock.Eq(updateParams)).Times(1).Return(nil, nil)
				store.EXPECT().GetProfile(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(db.Profile{}, sql.ErrConnDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			Name: "OK",
			Body: gin.H{
				"country":           updatedProfile.Country,
				"measurementSystem": updatedProfile.MeasurementSystem,
				"bodyWeight":        updatedProfile.BodyWeight,
				"bodyFat":           updatedProfile.BodyFat,
				"timeZoneOffset":    updatedProfile.TimezoneOffset,
				"userID":            userID.String(),
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProfile(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(profile, nil)
				store.EXPECT().UpdateProfile(gomock.Any(), gomock.Eq(updateParams)).Times(1).Return(nil, nil)
				store.EXPECT().GetProfile(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(updatedProfile, nil)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				validateNewProfileResponse(t, recorder.Body, ProfileResponse{
					Country:           updatedProfile.Country,
					MeasurementSystem: updatedProfile.MeasurementSystem,
					BodyWeight:        updatedProfile.BodyWeight,
					BodyFat:           updatedProfile.BodyFat,
					TimeZoneOffset:    updatedProfile.TimezoneOffset,
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

			data, err := json.Marshal(tc.Body)
			require.NoError(t, err)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := "/updateProfile"
			request, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkRes(t, recorder)
		})
	}
}

func TestDeleteProfile(t *testing.T) {
	userID := uuid.New()
	profile := GenRandProfile(userID)

	testCases := []struct {
		Name       string
		UserID     string
		buildStubs func(store *mockdb.MockStore)
		checkRes   func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			Name:   "Not Found -> Get Profile",
			UserID: userID.String(),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProfile(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(db.Profile{}, sql.ErrNoRows)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			Name:   "Internal Error -> Get Profile",
			UserID: userID.String(),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProfile(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(db.Profile{}, sql.ErrConnDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			Name:   "Internal Error -> Delete Profile",
			UserID: userID.String(),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProfile(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(profile, nil)
				store.EXPECT().DeleteProfile(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(sql.ErrConnDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			Name:   "OK -> No Content",
			UserID: userID.String(),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProfile(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(profile, nil)
				store.EXPECT().DeleteProfile(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(nil)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNoContent, recorder.Code)
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

			url := fmt.Sprintf("/deleteProfile/%s", tc.UserID)
			request, err := http.NewRequest(http.MethodDelete, url, nil)
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
		TimezoneOffset:    int32(util.RandomInt(100, 800)),
	}
}

func GenRandUpdateProfile(userID uuid.UUID, profileID int32) (db.UpdateProfileParams, db.Profile) {
	params := db.UpdateProfileParams{
		Country: sql.NullString{
			String: strings.ToUpper(util.RandomStr(3)),
			Valid:  true,
		},
		MeasurementSystem: sql.NullString{
			String: util.RandomStr(5),
			Valid:  true,
		},
		BodyWeight: sql.NullFloat64{
			Float64: 0,
			Valid:   false,
		},
		BodyFat: sql.NullFloat64{
			Float64: float64(util.RandomInt(10, 20)),
			Valid:   true,
		},
		TimezoneOffset: sql.NullInt32{
			Int32: int32(util.RandomInt(100, 800)),
			Valid: true,
		},
		UserID: userID.String(),
	}

	patchedProfile := db.Profile{
		Country:           params.Country.String,
		ID:                profileID,
		MeasurementSystem: params.MeasurementSystem.String,
		BodyWeight:        params.BodyWeight.Float64,
		BodyFat:           params.BodyFat.Float64,
		TimezoneOffset:    params.TimezoneOffset.Int32,
		UserID:            []byte(userID.String()),
	}

	return params, patchedProfile
}

func validateNewProfileResponse(t *testing.T, body *bytes.Buffer, res ProfileResponse) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var profile ProfileResponse
	err = json.Unmarshal(data, &profile)
	require.NoError(t, err)
	require.Equal(t, profile, res)
}
