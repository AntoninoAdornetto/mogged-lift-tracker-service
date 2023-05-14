package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ExpiredTokenError = errors.New("token has expired")
	InvalidTokenError = errors.New("token is invalid")
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserID    string    `json:"userID"`
	IssuedAt  time.Time `json:"issuedAt`
	ExpiredAt time.Time `json:"expiredAt`
}

func NewPayload(userID string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		UserID:    userID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ExpiredTokenError
	}
	return nil
}
