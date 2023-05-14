package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	mockdb "github.com/AntoninoAdornetto/mogged-lift-tracker-service/db/mock"
	db "github.com/AntoninoAdornetto/mogged-lift-tracker-service/db/sqlc"
	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type eqCreateUserParamsMatcher struct {
	args     db.CreateUserParams
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	args, ok := x.(db.CreateUserParams)
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

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.args, e.password)
}

func EqCreateUserParams(args db.CreateUserParams, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{args, password}
}

func TestCreateUser(t *testing.T) {
	userID := uuid.New()
	user := GenRandUser(userID)

	getUserByEmailRes := db.GetUserByEmailRow{
		FirstName:         user.FirstName,
		ID:                userID.String(),
		LastName:          user.LastName,
		EmailAddress:      user.EmailAddress,
		Password:          user.Password,
		PasswordChangedAt: user.PasswordChangedAt,
	}

	newUserTxRes := db.NewUserTxResults{
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		EmailAddress: user.EmailAddress,
		ID:           userID,
	}

	testCases := []struct {
		Name       string
		Body       gin.H
		buildStubs func(store *mockdb.MockStore)
		checkRes   func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			Name: "OK",
			Body: gin.H{
				"firstName":    user.FirstName,
				"lastName":     user.LastName,
				"emailAddress": user.EmailAddress,
				"password":     user.Password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				args := db.CreateUserParams{
					FirstName:    user.FirstName,
					LastName:     user.LastName,
					EmailAddress: user.EmailAddress,
				}
				store.EXPECT().GetUserByEmail(gomock.Any(), user.EmailAddress).Times(1).Return(getUserByEmailRes, nil)
				store.EXPECT().NewUserTx(gomock.Any(), EqCreateUserParams(args, user.Password)).Times(1).Return(newUserTxRes, nil)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				validateNewUserResponse(t, recorder.Body, newUserTxRes)
			},
		},
		{
			Name: "Bad Request",
			Body: gin.H{},
			buildStubs: func(store *mockdb.MockStore) {
				args := db.CreateUserParams{}
				store.EXPECT().NewUserTx(gomock.Any(), gomock.Eq(args)).Times(0).Return(db.NewUserTxResults{}, sql.ErrTxDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			Name: "Internal Error",
			Body: gin.H{
				"firstName":    user.FirstName,
				"lastName":     user.LastName,
				"emailAddress": "notFound@gmail.com",
				"password":     user.Password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				args := db.CreateUserParams{
					FirstName:    user.FirstName,
					LastName:     user.LastName,
					EmailAddress: "notFound@gmail.com",
				}
				store.EXPECT().NewUserTx(gomock.Any(), EqCreateUserParams(args, user.Password)).Times(1).Return(db.NewUserTxResults{}, sql.ErrTxDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			Name: "Not Found",
			Body: gin.H{
				"firstName":    user.FirstName,
				"lastName":     user.LastName,
				"emailAddress": user.EmailAddress,
				"password":     user.Password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				args := db.CreateUserParams{
					FirstName:    user.FirstName,
					LastName:     user.LastName,
					EmailAddress: user.EmailAddress,
				}
				store.EXPECT().NewUserTx(gomock.Any(), EqCreateUserParams(args, user.Password)).Times(1).Return(newUserTxRes, nil)
				store.EXPECT().GetUserByEmail(gomock.Any(), user.EmailAddress).Times(1).Return(db.GetUserByEmailRow{}, sql.ErrNoRows)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			Name: "Internal Error -> GetUser After TX",
			Body: gin.H{
				"firstName":    user.FirstName,
				"lastName":     user.LastName,
				"emailAddress": user.EmailAddress,
				"password":     user.Password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				args := db.CreateUserParams{
					FirstName:    user.FirstName,
					LastName:     user.LastName,
					EmailAddress: user.EmailAddress,
				}
				store.EXPECT().NewUserTx(gomock.Any(), EqCreateUserParams(args, user.Password)).Times(1).Return(newUserTxRes, nil)
				store.EXPECT().GetUserByEmail(gomock.Any(), user.EmailAddress).Times(1).Return(db.GetUserByEmailRow{}, sql.ErrConnDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
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

			url := "/createUser"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkRes(t, recorder)
		})
	}
}

func TestGetUserByEmail(t *testing.T) {
	userID := uuid.New()
	user := GenRandUser(userID)

	getUserByEmailRes := db.GetUserByEmailRow{
		EmailAddress:      user.EmailAddress,
		ID:                userID.String(),
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		PasswordChangedAt: user.PasswordChangedAt,
		Password:          user.Password,
	}

	response := UserResponse{
		FirstName:         getUserByEmailRes.FirstName,
		LastName:          getUserByEmailRes.LastName,
		EmailAddress:      getUserByEmailRes.EmailAddress,
		PasswordChangedAt: getUserByEmailRes.PasswordChangedAt,
		ID:                userID.String(),
	}

	testCases := []struct {
		Name         string
		EmailAddress string
		buildStubs   func(store *mockdb.MockStore)
		checkRes     func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			Name:         "OK",
			EmailAddress: user.EmailAddress,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUserByEmail(gomock.Any(), user.EmailAddress).Times(1).Return(getUserByEmailRes, nil)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				validateGetUserByEmailResponse(t, recorder.Body, response)
			},
		},
		{
			Name:         "Bad Request",
			EmailAddress: "test",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUserByEmail(gomock.Any(), "test").Times(0).Return(db.GetUserByEmailRow{}, sql.ErrConnDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			Name:         "Not Found",
			EmailAddress: "thurnis@gmail.com",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUserByEmail(gomock.Any(), gomock.Eq("thurnis@gmail.com")).Times(1).Return(db.GetUserByEmailRow{}, sql.ErrNoRows)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			Name:         "Internal Error",
			EmailAddress: user.EmailAddress,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUserByEmail(gomock.Any(), user.EmailAddress).Times(1).Return(db.GetUserByEmailRow{}, sql.ErrConnDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
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

			url := fmt.Sprintf("/getUserByEmail/%s", tc.EmailAddress)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkRes(t, recorder)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	userID := uuid.New()
	user := GenRandUser(userID)

	preGetUserByID := db.GetUserByIdRow{
		FirstName:         user.FirstName,
		EmailAddress:      user.EmailAddress,
		ID:                userID.String(),
		LastName:          user.LastName,
		Password:          user.Password,
		PasswordChangedAt: user.PasswordChangedAt,
	}

	newUserData := db.UpdateUserParams{
		FirstName: sql.NullString{
			Valid:  true,
			String: util.RandomStr(5),
		},
		LastName: sql.NullString{
			Valid:  true,
			String: util.RandomStr(6),
		},
		EmailAddress: sql.NullString{
			Valid:  true,
			String: util.RandomStr(7) + "@gmail.com",
		},
		UserID: userID.String(),
	}

	postGetUserByID := db.GetUserByIdRow{
		FirstName:         newUserData.FirstName.String,
		LastName:          newUserData.LastName.String,
		EmailAddress:      newUserData.EmailAddress.String,
		ID:                userID.String(),
		PasswordChangedAt: time.Now(),
		Password:          user.Password,
	}

	response := UserResponse{
		FirstName:         postGetUserByID.FirstName,
		LastName:          postGetUserByID.LastName,
		EmailAddress:      postGetUserByID.EmailAddress,
		PasswordChangedAt: postGetUserByID.PasswordChangedAt,
		ID:                userID.String(),
	}

	testCases := []struct {
		Name       string
		Body       gin.H
		buildStubs func(store *mockdb.MockStore)
		checkRes   func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			Name: "OK",
			Body: gin.H{
				"firstName":    response.FirstName,
				"lastName":     response.LastName,
				"emailAddress": response.EmailAddress,
				"id":           userID.String(),
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUserById(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(preGetUserByID, nil)
				store.EXPECT().UpdateUser(gomock.Any(), gomock.Eq(newUserData)).Times(1).Return(nil)
				store.EXPECT().GetUserById(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(postGetUserByID, nil)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				validateGetUserByEmailResponse(t, recorder.Body, response)
			},
		},
		{
			Name: "Bad Request",
			Body: gin.H{},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUserById(gomock.Any(), gomock.Any()).Times(0).Return(db.GetUserByIdRow{}, sql.ErrConnDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			Name: "Internal Error",
			Body: gin.H{
				"firstName":    response.FirstName,
				"lastName":     response.LastName,
				"emailAddress": response.EmailAddress,
				"id":           userID.String(),
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUserById(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(preGetUserByID, nil)
				store.EXPECT().UpdateUser(gomock.Any(), gomock.Eq(newUserData)).Times(1).Return(sql.ErrConnDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			Name: "Not Found",
			Body: gin.H{
				"firstName":    response.FirstName,
				"lastName":     response.LastName,
				"emailAddress": response.EmailAddress,
				"id":           userID.String(),
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUserById(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(db.GetUserByIdRow{}, sql.ErrNoRows)
				store.EXPECT().UpdateUser(gomock.Any(), gomock.Eq(newUserData)).Times(0).Return(sql.ErrConnDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			Name: "Internal Error -> Get User",
			Body: gin.H{
				"firstName":    response.FirstName,
				"lastName":     response.LastName,
				"emailAddress": response.EmailAddress,
				"id":           userID.String(),
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUserById(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(db.GetUserByIdRow{}, sql.ErrConnDone)
				store.EXPECT().UpdateUser(gomock.Any(), gomock.Eq(newUserData)).Times(0).Return(sql.ErrConnDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			Name: "Internal Error -> Get User After Updating",
			Body: gin.H{
				"firstName":    response.FirstName,
				"lastName":     response.LastName,
				"emailAddress": response.EmailAddress,
				"id":           userID.String(),
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUserById(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(preGetUserByID, nil)
				store.EXPECT().UpdateUser(gomock.Any(), gomock.Eq(newUserData)).Times(1).Return(nil)
				store.EXPECT().GetUserById(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(db.GetUserByIdRow{}, sql.ErrConnDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
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

			data, err := json.Marshal(tc.Body)
			require.NoError(t, err)

			url := "/updateUser"
			request, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkRes(t, recorder)
		})
	}
}

func TestChangePassword(t *testing.T) {
	userID := uuid.New()
	user := GenRandUser(userID)

	getUserByIdRes := db.GetUserByIdRow{
		EmailAddress:      user.EmailAddress,
		ID:                userID.String(),
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		PasswordChangedAt: user.PasswordChangedAt,
		Password:          user.Password,
	}

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
			Name: "Not Found -> Get User",
			Body: gin.H{
				"id":              userID.String(),
				"currentPassword": user.Password,
				"newPassword":     "shannonsbosoms@69",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUserById(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(db.GetUserByIdRow{}, sql.ErrNoRows)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			Name: "Internal Error -> Get User",
			Body: gin.H{
				"id":              userID.String(),
				"currentPassword": user.Password,
				"newPassword":     "shannonsbosoms@69",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUserById(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(db.GetUserByIdRow{}, sql.ErrConnDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			Name: "Unauthorized -> incorrect current password",
			Body: gin.H{
				"id":              userID.String(),
				"currentPassword": util.RandomStr(9),
				"newPassword":     "shannonsbosoms@69",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUserById(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(getUserByIdRes, nil)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			Name: "Internal Error -> change pasword",
			Body: gin.H{
				"id":              userID.String(),
				"currentPassword": user.Password,
				"newPassword":     "shannonsbosoms@69",
			},
			buildStubs: func(store *mockdb.MockStore) {
				args := db.ChangePasswordParams{
					UserID:   userID.String(),
					Password: "shannonsbosoms@69",
				}
				store.EXPECT().GetUserById(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(getUserByIdRes, nil)
				store.EXPECT().ChangePassword(gomock.Any(), gomock.Eq(args)).Times(1).Return(sql.ErrConnDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			Name: "OK -> No Content",
			Body: gin.H{
				"id":              userID.String(),
				"currentPassword": user.Password,
				"newPassword":     "shannonsbosoms@69",
			},
			buildStubs: func(store *mockdb.MockStore) {
				args := db.ChangePasswordParams{
					UserID:   userID.String(),
					Password: "shannonsbosoms@69",
				}
				store.EXPECT().GetUserById(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(getUserByIdRes, nil)
				store.EXPECT().ChangePassword(gomock.Any(), gomock.Eq(args)).Times(1).Return(nil)
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

			data, err := json.Marshal(tc.Body)
			require.NoError(t, err)

			url := "/changePassword"
			request, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkRes(t, recorder)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	userID := uuid.New()
	user := GenRandUser(userID)

	getUserByIdRes := db.GetUserByIdRow{
		FirstName:         user.FirstName,
		ID:                userID.String(),
		LastName:          user.LastName,
		EmailAddress:      user.EmailAddress,
		Password:          user.Password,
		PasswordChangedAt: user.PasswordChangedAt,
	}

	testCases := []struct {
		Name       string
		UserID     string
		buildStubs func(store *mockdb.MockStore)
		checkRes   func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			Name:   "Not found -> Get User",
			UserID: userID.String(),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUserById(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(db.GetUserByIdRow{}, sql.ErrNoRows)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			Name:   "Internal Error -> Get User",
			UserID: userID.String(),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUserById(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(db.GetUserByIdRow{}, sql.ErrConnDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			Name:   "Internal Error -> Delete User",
			UserID: userID.String(),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUserById(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(getUserByIdRes, nil)
				store.EXPECT().DeleteUser(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(sql.ErrConnDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			Name:   "OK -> No Content",
			UserID: userID.String(),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUserById(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(getUserByIdRes, nil)
				store.EXPECT().DeleteUser(gomock.Any(), gomock.Eq(userID.String())).Times(1).Return(nil)
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

			url := fmt.Sprintf("/deleteUser/%s", tc.UserID)
			request, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkRes(t, recorder)
		})
	}
}

func GenRandUser(userID uuid.UUID) db.User {
	hashedPassword, _ := util.HashPassword(util.RandomStr(10))
	return db.User{
		ID:                []byte(userID.String()),
		LastName:          util.RandomStr(10),
		FirstName:         util.RandomStr(5),
		Password:          hashedPassword,
		EmailAddress:      util.RandomStr(5) + "@gmail.com",
		PasswordChangedAt: time.Date(1970, time.January, 01, 01, 00, 00, 00, time.Now().Location()),
	}
}

func validateNewUserResponse(t *testing.T, body *bytes.Buffer, res db.NewUserTxResults) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var user db.NewUserTxResults
	err = json.Unmarshal(data, &user)
	require.NoError(t, err)
	require.Equal(t, user, res)
}

func validateGetUserByEmailResponse(t *testing.T, body *bytes.Buffer, res UserResponse) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var user UserResponse
	err = json.Unmarshal(data, &user)
	require.NoError(t, err)
	require.WithinDuration(t, user.PasswordChangedAt, res.PasswordChangedAt, time.Minute)
	require.Equal(t, user.ID, res.ID)
	require.Equal(t, user.FirstName, res.FirstName)
	require.Equal(t, user.LastName, res.LastName)
	require.Equal(t, user.EmailAddress, res.EmailAddress)
}
