package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", LoginUser).Methods(http.MethodPost).Name("Authorize User")
	r.HandleFunc("/something", AuthorizeUser([]string{"canDoSomething"}, doSomething))
	r.HandleFunc("/something/else", AuthorizeUser([]string{"canDoSomethingElse"}, doSomethingElse))

	loggedHandler := handlers.LoggingHandler(os.Stdout, r)
	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf("%s:%d", Address, Port), loggedHandler))
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		SendErrorJSONResponse(w, INTERNALSERVICEERROR, http.StatusInternalServerError)
		return
	}

	var UserCredentials struct {
		Email    string
		Password string
	}
	err = json.Unmarshal(body, &UserCredentials)
	if err != nil {
		SendErrorJSONResponse(w, INVALIDCREDENTIALS, http.StatusBadRequest)
	}

	db, err := NewMySQLDBConnection()
	defer db.Close()
	if err != nil {
		log.Info(DBCONNECTIONERROR)
		SendErrorJSONResponse(w, INTERNALSERVICEERROR, http.StatusInternalServerError)
		return
	}

	var User User
	results := db.Where(&User{Email: UserCredentials.Email, Active: true}).First(&User)
	if results.Error != nil || !CheckPasswordHash(UserCredentials.Password, string(User.Password)) {
		SendErrorJSONResponse(w, INVALIDEMAILORPASSWORD, http.StatusBadRequest)
		return
	}

	claims := TokenClaims{
		ID:    User.ID,
		Email: User.Email,
		Scope: CalculateUserScope(User),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    "Notes",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWTSECRET))
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
}

func doSomething(w http.ResponseWriter, r *http.Request) {
	type Something struct {
		Foo string `json:foo`
		Bar int    `json:bar`
	}
	var something Something
	SendAsJSONResponse(w, something, http.StatusOK)

}

func doSomethingElse(w http.ResponseWriter, r *http.Request) {
	type Something struct {
		Foo string `json:foo`
		Bar int    `json:bar`
	}
	var something Something
	SendAsJSONResponse(w, something, http.StatusOK)
}
