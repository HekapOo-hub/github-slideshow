package main

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	takenTTL   = 1 * time.Hour
	signingKey = "afho457iewnflfp2452oivjnsqojr"
)

type TokenError struct {
	Message string
}

func (t *TokenError) Error() string {
	return t.Message
}

type tokenClaims struct {
	Role string
	Id   int
	jwt.StandardClaims
}

type authentication interface {
	validToken() (int, bool)
	decode(string) (*tokenClaims, error)
	CheckRole(string) (bool, error)
}

type AuthenticationService struct {
	repository
}

func (a *AuthenticationService) decode(token string) (*tokenClaims, error) {
	tokenType, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := tokenType.Claims.(*tokenClaims); ok && tokenType.Valid {
		return claims, nil
	} else {

		return nil, err
	}
}
func (a *AuthenticationService) validToken(token string) (string, error) {
	tClaims, err := a.decode(token)
	if err != nil {
		return "", &TokenError{Message: "Token decode error"}
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tClaims)
	str, err := jwtToken.SignedString([]byte(signingKey))
	if token == str {
		return tClaims.Role, err
	} else {
		return "", &TokenError{Message: "this is not your token!"}
	}
}
func (a *AuthenticationService) CheckRole(roles []string, token string) (bool, error) {
	r, err := a.validToken(token)
	if err != nil {
		return false, err
	}
	for _, role := range roles {
		if r == role {
			return true, nil
		}
	}
	return false, nil
}
