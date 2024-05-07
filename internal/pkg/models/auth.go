package models

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	JWTSigningKey = "helloWorld"
	UserIdCtxKey  = "userId"
)

type TokenClaims struct {
	UserId int `json:"userId"`
	jwt.RegisteredClaims
}

type SignUp struct {
	Fullname string `json:"fullname"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (s *SignUp) validate() error {
	if len(s.Fullname) == 0 {
		return fmt.Errorf("Required field: fullname is empty")
	}
	if len(s.Login) == 0 {
		return fmt.Errorf("Required field: fullname is empty")
	}
	if len(s.Password) > 8 {
		return fmt.Errorf("Password length should be minimum 8 characters")
	}
	return nil
}

type SignIn struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserProfile struct {
	Fullname     string    `json:"fullname"`
	Login        string    `json:"login" `
	RegisteredAt time.Time `json:"registeredAt"`
}
