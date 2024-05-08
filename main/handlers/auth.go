package handlers

import (
	"errors"
	"hibernum-backend/main/service"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_pwd")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Authentication struct{}

func NewAuthentication() *Authentication {
	return &Authentication{}
}

func (auth *Authentication) CreateToken(username string) (string, error) {
	expTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
			IssuedAt:  time.Now().Local().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (auth *Authentication) ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (auth *Authentication) LoginHandler(username string, password string, w http.ResponseWriter) {
	err := service.ValidateCredentials(username, password)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) || errors.Is(err, service.ErrInvalidCredentials) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token, err := auth.CreateToken(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", token)
	w.WriteHeader(http.StatusOK)
}
