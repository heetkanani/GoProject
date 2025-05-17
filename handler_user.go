package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (apiConfig *apiConfig)handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct{
		Name string `name`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err:= decoder.Decode(&params)
	if err != nil{
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// apiConfig.DB.CreateUser()
	responseWithError(w, 400, "Something went wrong")
}

//this is for getting the user by API key
func (apiConfig *apiConfig)handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err!=nil{
		responseWithError(w, 403, fmt.Sprintf("auth error: %v", err));
	}
}
