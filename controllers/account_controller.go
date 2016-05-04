package controllers

import (
	"github.com/xlab-si/e2ee-server/core/db"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	log "github.com/Sirupsen/logrus"

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
}

type AccountExistsResponse struct {
   	Exists bool `json:"exists"`
}

func AccountExists(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	accountId, _ := GetAccountInfo(r)
	account := db.FindAccount(accountId)
	
	exists := false
	if account.AccountId != "" {
		exists = true
	} else {
		log.WithFields(log.Fields{
                	"accountId": accountId,
	        }).Info("Account not found - this is expected when signing in for the first time")
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
	accountId, _ := GetAccountInfo(r)
	account := db.FindAccount(accountId)
	
	var m = AccountResponseMessage{
	    Success: true,
    	    Account: account,
	}

	if account.AccountId != "" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(m); err != nil {
			panic(err)
		}
		return
	} else {
		log.WithFields(log.Fields{
                	"accountId": accountId,
	        }).Error("Account not found")
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
	accountId, username := GetAccountInfo(r)

	var account db.Account
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		log.WithFields(log.Fields{
                	"accountId": accountId,
	        }).Error("Reading body request failed")
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &account); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		log.WithFields(log.Fields{
                	"accountId": accountId,
	        }).Error("Unprocessable entity when trying to create an account")
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

