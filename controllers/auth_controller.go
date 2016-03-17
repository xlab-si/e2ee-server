package controllers

import (
	"github.com/xlab-si/e2ee-server/services"
	"github.com/xlab-si/e2ee-server/core/db"
	"encoding/json"
	"net/http"
	"io"
	"io/ioutil"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials db.UserCredentials
        body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
                panic(err)
        }
        if err := r.Body.Close(); err != nil {
                panic(err)
        }
        if err := json.Unmarshal(body, &credentials); err != nil {
                w.Header().Set("Content-Type", "application/json; charset=UTF-8")
                w.WriteHeader(422) // unprocessable entity
                if err := json.NewEncoder(w).Encode(err); err != nil {
                        panic(err)
                }
                return
        }

	responseStatus, token := services.Login(credentials)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(token)
}

func RefreshToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	uuid, username, accountId := ExtractTokenInfo(r)
	w.Header().Set("Content-Type", "application/json")
	w.Write(services.RefreshToken(uuid, username, accountId))
}

func Logout(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	err := services.Logout(r)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
