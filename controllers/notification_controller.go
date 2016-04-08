package controllers

import (
	"github.com/mancabizjak/e2ee-server/core/db"
	"encoding/json"
	//"io"
	//"io/ioutil"
	"net/http"

	//"github.com/gorilla/mux"
	//"github.com/jeffail/gabs"
	//"log"
)

type NotificationsPacket struct {
    Success bool `json:"success"`
    Error string `json:"error"`
    Notifications []Notification `json:"messages"`
}

type NotificationsDeleteResponse struct {
    Success bool `json:"success"`
    Error string `json:"error"`
}

func NotificationsGet(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_, _, accountId := ExtractTokenInfo(r)

	messages := db.GetNotifications(accountId)
	rmessages := []Notification{}
	for _, me := range messages {
		account := db.FindAccount(me.FromAccountId) // me.FromAccountId is 0, which it should not be
		username := account.Username
		var rm = Notification {
			FromUsername: username,
			HeadersCiphertext: me.HeadersCiphertext,	
			PayloadCiphertext: me.PayloadCiphertext,	
		}
                rmessages = append(rmessages, rm)
        }


	var m = NotificationsPacket{
	    Success: true,
	    Error: "",
	    Notifications: rmessages,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m); err != nil {
                panic(err)
        }
}

func NotificationsDelete(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_, _, accountId := ExtractTokenInfo(r)

	db.DBDeleteNotifications(accountId)
	var m = NotificationsDeleteResponse{
	    Success: true,
	    Error: "",
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m); err != nil {
                panic(err)
        }
}




