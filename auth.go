package main

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

var AvailablePermissions = map[string]int{
	"allow_api":          1,
	"canDoSomething":     2,
	"canDoSomethingElse": 4,
}

var PermissionRoles = map[string][]string{
	"is_basic_user": {
		"allow_api",
		"canDoSomething",
	},
	"is_premium_user": {
		"canDoSomethingElse",
	},
}

var PermissionGroups = map[string][]string{
	"basic": {
		"is_basic_user",
	},
	"premium": {
		"is_basic_user",
		"is_premium_user",
	},
}

func CalculateUserScope(user User) (scope int) {
	roles := PermissionGroups[user.UserGroupKey]
	scope = 0
	for _, r := range roles {
		for _, p := range PermissionRoles[r] {
			scope += AvailablePermissions[p]
		}
	}
	return
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hasPermission(scope int, permissions []string) error {
	for _, p := range permissions {
		if AvailablePermissions[p]&scope == 0 {
			return errors.New("permission not in scope")
		}
	}
	return nil
}

func TokenIsValid(tokenString string) (claims TokenClaims, err error) {
	claims = TokenClaims{}
	_, err = jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSECRET), nil
	})
	return
}

func AuthorizeUser(permissions []string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		tokenString := strings.Replace(req.Header.Get("Authorization"), "JWT ", "", -1)
		if tokenString == "" {
			SendErrorJSONResponse(w, INTERNALSERVICEERROR, http.StatusInternalServerError)
			return
		}

		token, err := TokenIsValid(tokenString)
		if err != nil {
			log.Error(err)
			SendErrorJSONResponse(w, TOKENNOTVALID, http.StatusInternalServerError)
			return
		}

		err = hasPermission(token.Scope, permissions)
		if err != nil {
			log.Error(err)
			SendErrorJSONResponse(w, PERMISSIONNOTGIVEN, http.StatusInternalServerError)
			return
		}
		handler(w, req)
	}
}
