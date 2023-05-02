package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
					Password:     user.Password,
				}
				store.EXPECT().GetUserByEmail(gomock.Any(), user.EmailAddress).Times(1).Return(getUserByEmailRes, nil)
				store.EXPECT().NewUserTx(gomock.Any(), gomock.Eq(args)).Times(1).Return(newUserTxRes, nil)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				validateUserResponse(t, recorder.Body, newUserTxRes)
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

			url := "/createUser"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkRes(t, recorder)
		})
	}
}

func GenRandUser(userID uuid.UUID) db.User {
	return db.User{
		ID:                []byte(userID.String()),
		LastName:          util.RandomStr(10),
		FirstName:         util.RandomStr(5),
		Password:          util.RandomStr(10),
		EmailAddress:      util.RandomStr(5) + "@gmail.com",
		PasswordChangedAt: time.Date(1970, time.January, 01, 01, 00, 00, 00, time.Now().Location()),
	}
}

func validateUserResponse(t *testing.T, body *bytes.Buffer, res db.NewUserTxResults) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var acc db.NewUserTxResults
	err = json.Unmarshal(data, &acc)
	require.NoError(t, err)
	require.Equal(t, acc, res)
}
