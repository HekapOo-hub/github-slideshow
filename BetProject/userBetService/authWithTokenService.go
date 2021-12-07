package main

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	takenTTL   = 1 * time.Hour
	signingKey = "afho457iewnflfp2452oivjnsqojr"
)

type tokenClaims struct {
	Role string
	Id   int
	jwt.StandardClaims
}

type tokenMaker interface {
	AuthorizeWithToken(name, password string) (string, error)
}
type AuthWithTokenService struct {
	*AuthorizationService
}

func (a *AuthWithTokenService) AuthorizeWithToken(name, password string) (string, error) {
	role, id, err := a.Authorize(name, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(takenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		}, Role: role, Id: id,
	})

	return token.SignedString([]byte(signingKey))

}
