package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

// CreateToken creates a new token for a specific username and duration
func (jwtMaker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	//1- Generate new payload
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	//2- Generate token using HS256 signing method with payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	//3- Encrypt token with private secret key
	var encryptedToken string
	if encryptedToken, err = token.SignedString([]byte(jwtMaker.secretKey)); err != nil {
		return "", err
	}
	return encryptedToken, nil
}

// VerifyToken checks if the token is valid or not
func (jwtMaker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	// The purpose behind this keyFunc is to validate the signing method of the passed token
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(jwtMaker.secretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		vErr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(vErr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}
