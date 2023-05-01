package api

import (
	"testing"
	"time"

	db "github.com/AntoninoAdornetto/mogged-lift-tracker-service/db/sqlc"
	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/google/uuid"
)

func TestCreateUser(t *testing.T) {
	userID := uuid.New()
	GenRandUser(userID)
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
