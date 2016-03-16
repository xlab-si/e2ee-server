package controllers

import (
	"github.com/mancabizjak/e2ee-server/core/db"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	//"github.com/gorilla/mux"
	//"log"
	//"github.com/jeffail/gabs"
)

type ResponseMessage struct {
    Success bool `json:"success"`
    Error string `json:"error"`
}

type AccountResponseMessage struct {
    Success bool `json:"success"`
    Account db.Account `json:"account"`
    Sid int `json:"sid"`
}

type AccountExistsResponse struct {
    Exists bool `json:"exists"`
}

func AccountExists(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_, _, accountId := ExtractTokenInfo(r)
	account := db.FindAccount(accountId)
	
	exists := false
	if account.AccountId != 0 {
		exists = true
	}

	var m = AccountExistsResponse{
	    Exists: exists,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m); err != nil {
		panic(err)
	}
}

func AccountGet(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_, _, accountId := ExtractTokenInfo(r)

	account := db.FindAccount(accountId)
	
	var m = AccountResponseMessage{
	    Success: true,
    	    Account: account,
	}

	if account.AccountId != 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(m); err != nil {
			panic(err)
		}
		return
	}

	m = AccountResponseMessage{
	    Success: false,
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(m); err != nil {
		panic(err)
	}
}

func AccountCreate(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_, username, accountId := ExtractTokenInfo(r)

	var account db.Account
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &account); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	var m = ResponseMessage{
	    Success: true,
	}

	account.Username = username
	account.AccountId = accountId
	db.StoreAccount(account)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(m); err != nil {
		panic(err)
	}
}

