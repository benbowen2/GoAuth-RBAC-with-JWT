package main

import "github.com/dgrijalva/jwt-go"

type ApiErrorResponse struct {
	Message     string `json:"message"`
	Title       string
	Description string
	Code        string
}

func (api ApiErrorResponse) Error() string {
	return api.Message
}

type JwtToken struct {
	Token string `json:"token"`
}

type TokenClaims struct {
	ID    uint   `json:"id"`
	Scope int    `json:"scope"`
	Email string `json:"email"`
	jwt.StandardClaims
}
