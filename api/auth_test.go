package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
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

type eqCreateSessionParamsMatcher struct {
	args     loginRequest
	password string
}

func (e eqCreateSessionParamsMatcher) Matches(x interface{}) bool {
	args, ok := x.(loginRequest)
	if !ok {
		return false
	}

	err := util.ValidatePassword(e.password, args.Password)
	if err != nil {
		return false
	}

	e.args.Password = args.Password
	return reflect.DeepEqual(e.args, args)
}

func (e eqCreateSessionParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.args, e.password)
}

func EqCreateSessionParams(args loginRequest, password string) gomock.Matcher {
	return eqCreateSessionParamsMatcher{args, password}
}

func TestLogin(t *testing.T) {
	userID := uuid.New()
	loginResponse := GenLoginResponse(userID, t)

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
			Name: "Get User -> Not Found",
			Body: gin.H{
				"emailAddress": "random@gmail.com",
				"password":     "random",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Eq("random@gmail.com")).
					Times(1).
					Return(db.GetUserByEmailRow{}, sql.ErrNoRows)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			Name: "Get User -> Internal Error",
			Body: gin.H{
				"emailAddress": loginResponse.User.EmailAddress,
				"password":     "pass",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Eq(loginResponse.User.EmailAddress)).
					Times(1).
					Return(db.GetUserByEmailRow{}, sql.ErrConnDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			Name: "UnAuthorized",
			Body: gin.H{
				"emailAddress": loginResponse.User.EmailAddress,
				"password":     "pass1",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Eq(loginResponse.User.EmailAddress)).
					Times(1).
					Return(db.GetUserByEmailRow{
						EmailAddress:      loginResponse.User.EmailAddress,
						ID:                loginResponse.User.ID,
						FirstName:         loginResponse.User.FirstName,
						LastName:          loginResponse.User.LastName,
						PasswordChangedAt: loginResponse.User.PasswordChangedAt,
						Password:          "pass",
					}, nil)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		// @TODO - fix Create Session -> Internal Error Test
		// {
		// 	Name: "Create Session -> Internal Error",
		// 	Body: gin.H{
		// 		"emailAddress": loginResponse.User.EmailAddress,
		// 		"password":     "random",
		// 	},
		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		args := db.CreateSessionParams{
		// 			ID:           loginResponse.SessionID,
		// 			UserID:       loginResponse.User.ID,
		// 			ClientIp:     "",
		// 			UserAgent:    "",
		// 			ExpiresAt:    loginResponse.AccessTokenExpiresAt,
		// 			RefreshToken: loginResponse.RefreshToken,
		// 		}
		// 		store.EXPECT().
		// 			GetUserByEmail(gomock.Any(), gomock.Eq(loginResponse.User.EmailAddress)).
		// 			Times(1).
		// 			Return(db.GetUserByEmailRow{
		// 				EmailAddress:      loginResponse.User.EmailAddress,
		// 				ID:                loginResponse.User.ID,
		// 				FirstName:         loginResponse.User.FirstName,
		// 				LastName:          loginResponse.User.LastName,
		// 				PasswordChangedAt: loginResponse.User.PasswordChangedAt,
		// 				Password:          "pass",
		// 			}, nil)

		// 		// store.EXPECT().
		// 		// 	CreateSession(gomock.Any(), gomock.Eq(args)).
		// 		// 	Times(1).
		// 		// 	Return(sql.ErrConnDone)
		// 	},
		// 	checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		// 		require.Equal(t, http.StatusInternalServerError, recorder.Code)
		// 	},
		// },
		// @TODO: Fix Get User -> Ok test
		// {
		// 	Name: "Get User -> OK",
		// 	Body: gin.H{
		// 		"emailAddress": loginResponse.User.EmailAddress,
		// 		"password":     "pass",
		// 	},
		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		store.EXPECT().
		// 			GetUserByEmail(gomock.Any(), gomock.Eq(loginResponse.User.EmailAddress)).
		// 			Times(1).
		// 			Return(db.GetUserByEmailRow{
		// 				EmailAddress:      loginResponse.User.EmailAddress,
		// 				ID:                loginResponse.User.ID,
		// 				FirstName:         loginResponse.User.FirstName,
		// 				LastName:          loginResponse.User.LastName,
		// 				PasswordChangedAt: loginResponse.User.PasswordChangedAt,
		// 				Password:          "pass",
		// 			}, nil)
		// 	},
		// 	checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		// 		require.Equal(t, http.StatusOK, recorder.Code)
		// 	},
		// },
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

			url := "/api/auth/login"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkRes(t, recorder)
		})
	}
}

func GenLoginResponse(userID uuid.UUID, t *testing.T) loginResponse {
	user := GenRandUser(userID)
	maker, err := token.NewJWTMaker(util.RandomStr(32))
	require.NoError(t, err)

	accessToken, accessPayload, err := maker.CreateToken(string(user.ID), time.Minute)
	require.NoError(t, err)

	refreshToken, refreshPayload, err := maker.CreateToken(string(user.ID), time.Minute)

	return loginResponse{
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		SessionID:             uuid.NewString(),
		User: UserResponse{
			FirstName:         user.FirstName,
			LastName:          user.LastName,
			EmailAddress:      user.EmailAddress,
			PasswordChangedAt: user.PasswordChangedAt,
			ID:                string(user.ID),
		},
	}
}
