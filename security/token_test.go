package security

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

func TestNewToken(t *testing.T) {
	id := bson.NewObjectId().Hex()
	token, err := NewToken(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestNewTokenPayload(t *testing.T) {
	id := bson.NewObjectId().Hex()
	token, err := NewToken(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	payload, err := NewTokenPayload(token)
	assert.NoError(t, err)
	assert.NotNil(t, payload)
	assert.Equal(t, id, payload.UserId)

	tokenExpired := getExpiredToken(id)
	payload, err = NewTokenPayload(tokenExpired)
	assert.Error(t, err)
	assert.Nil(t, payload)
}

func getExpiredToken(userId string) string {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 5 * -1).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    userId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtSecretKey)
	return tokenString
}
