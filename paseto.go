package main

import (
	"time"

	"github.com/o1egl/paseto"
)

var symmetricKey = []byte("rxrXY4lkiqKgRfeK0gOrOXxaHTKiGRl5")

func CreateToken(email string, fullName string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(email, fullName, duration)
	if err != nil {
		return "", nil, err
	}

	token, err := paseto.NewV2().Encrypt(symmetricKey, payload, nil)
	if err != nil {
		return "", nil, err
	}

	return token, payload, nil
}

func VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := paseto.NewV2().Decrypt(token, symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
