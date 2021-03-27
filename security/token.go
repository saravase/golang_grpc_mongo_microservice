package security

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	ErrInvalidToken = errors.New("invalid jwt token")
	jwtSecretKey    = []byte(os.Getenv("JWT_SECRET_KEY"))
)

func NewToken(userId string) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    userId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

func parseJwtCallback(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return jwtSecretKey, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, parseJwtCallback)
}

type TokenPayload struct {
	UserId    string
	CreatedAt time.Time
	ExpiredAt time.Time
}

func ExtractToken(r *http.Request) (string, error) {
	header := strings.TrimSpace(r.Header.Get("Authorization"))
	splitted := strings.Split(header, " ")
	if len(splitted) != 2 {
		return "", ErrInvalidToken
	}

	return splitted[1], nil
}

func NewTokenPayload(tokenString string) (*TokenPayload, error) {
	token, err := ParseToken(tokenString)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !token.Valid || !ok {
		return nil, ErrInvalidToken
	}

	id, _ := claims["iss"].(string)
	issuedAt, _ := claims["iat"].(int64)
	expiredAt, _ := claims["exp"].(int64)

	return &TokenPayload{
		UserId:    id,
		CreatedAt: time.Unix(issuedAt, 0),
		ExpiredAt: time.Unix(expiredAt, 0),
	}, nil
}
