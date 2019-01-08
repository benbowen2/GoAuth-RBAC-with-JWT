package main

import (
	"encoding/json"
	"net/http"
)

func SendAsJSONResponse(w http.ResponseWriter, v interface{}, desiredHttpStatusCode int) {
	jsonBody, _ := json.Marshal(v)
	SendJSONResponse(w, jsonBody, desiredHttpStatusCode)
}

func SendJSONResponse(w http.ResponseWriter, json []byte, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(httpStatusCode)
	w.Write(json)
}

func SendErrorJSONResponse(w http.ResponseWriter, errorMessage string, httpStatusCode int) {
	response := ApiErrorResponse{
		Message: errorMessage,
	}

	jsonBody, _ := json.Marshal(response)
	SendJSONResponse(w, jsonBody, httpStatusCode)
	return
}
