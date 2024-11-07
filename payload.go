package main

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type Payload struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	FullName  string    `json:"full_name"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, fullName string, hour time.Duration) (*Payload, error) {
	tokenID := uuid.NewString()

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		FullName:  fullName,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(hour * time.Hour),
	}

	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	
	return nil
}
