package main

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

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

type User struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time `gorm:"index";sql:"DEFAULT:'current_timestamp'"`
	UpdatedAt time.Time `gorm:"index";sql:"DEFAULT:'current_timestamp'"`
	Email     string   `gorm:"index"`
	FirstName string   `gorm:"index"`
	LastName  string   `gorm:"index"`
	Password  string `json:",omitempty"`
	Active    bool     `gorm:"index;default:true";sql:"not null"`
	UserGroupKey string `gorm:"index"`
}

