package tests

import (
	"backend-master-class/token"
	"backend-master-class/util"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestPasetoMaker(t *testing.T) {
	t.Run("When token is valid", func(t *testing.T) {
		secretKey := util.RandomString(32)
		pasetoMaker, err := token.NewPasetoMaker(secretKey)
		assert.NoError(t, err)
		username := util.RandomOwner()
		duration := time.Minute
		issuedAt := time.Now()
		expiredAt := issuedAt.Add(duration)

		paseto, err := pasetoMaker.CreateToken(username, duration)
		assert.NoError(t, err)
		assert.NotEmpty(t, paseto)

		payload, err := pasetoMaker.VerifyToken(paseto)
		assert.NoError(t, err)
		assert.NotEmpty(t, payload)
		assert.NotZero(t, payload.ID)
		assert.Equal(t, username, payload.Username)
		assert.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
		assert.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
	})

	t.Run("When token is expired", func(t *testing.T) {
		secretKey := util.RandomString(32)
		pasetoMaker, err := token.NewPasetoMaker(secretKey)
		assert.NoError(t, err)

		username := util.RandomOwner()
		duration := -(time.Minute)

		paseto, err := pasetoMaker.CreateToken(username, duration)
		assert.NoError(t, err)
		assert.NotEmpty(t, paseto)

		payload, err := pasetoMaker.VerifyToken(paseto)
		assert.Error(t, err)
		assert.EqualError(t, err, token.ErrExpiredToken.Error())
		assert.Nil(t, payload)
	})

	t.Run("When token has invalid algo none", func(t *testing.T) {
		payload, err := token.NewPayload(util.RandomOwner(), time.Minute)
		assert.NoError(t, err)

		//2- Generate token using HS256 signing method with payload
		jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
		signedToken, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
		assert.NoError(t, err)

		maker, err := token.NewJWTMaker(util.RandomString(36))
		assert.NoError(t, err)

		payload, err = maker.VerifyToken(signedToken)
		assert.Error(t, err)
		assert.EqualError(t, err, token.ErrInvalidToken.Error())
		assert.Nil(t, payload)

	})
}
