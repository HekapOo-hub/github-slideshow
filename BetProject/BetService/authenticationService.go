package main

import (
	"fmt"
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
	validToken(string) (*tokenClaims, error)
	decode(string) (*tokenClaims, error)
	CheckRole([]string, string) (int, error)
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
		fmt.Println(claims)
		return claims, nil
	} else {

		return nil, err
	}
}
func (a *AuthenticationService) validToken(token string) (*tokenClaims, error) {
	tClaims, err := a.decode(token)
	if err != nil {
		return nil, &TokenError{Message: "Token decode error"}
	}
	fmt.Println(tClaims.Role + " Role")
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tClaims)
	str, err := jwtToken.SignedString([]byte(signingKey))
	if token == str {
		return tClaims, err
	} else {
		return nil, &TokenError{Message: "this is not your token!"}
	}
}
func (a *AuthenticationService) CheckRole(roles []string, token string) (int, error) {
	tClaims, err := a.validToken(token)
	if err != nil {
		return -1, err
	}
	for _, role := range roles {
		if tClaims.Role == role {
			return tClaims.Id, nil
		}
	}
	return -1, nil
}
