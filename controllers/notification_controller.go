package controllers

import (
	"github.com/xlab-si/e2ee-server/core/db"
	"encoding/json"
	"net/http"
	//log "github.com/Sirupsen/logrus"
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
	accountId, _ := GetAccountInfo(r)

	messages := db.GetNotifications(accountId)
	rmessages := []Notification{}
	for _, msg := range messages {
		/*log.WithFields(log.Fields{
                        "accountId": accountId,
                }).Info(msg)*/

		account := db.FindAccount(msg.FromAccountId) // msg.FromAccountId is 0, which it should not be
		username := account.Username
		var rm = Notification {
			FromUsername: username,
			HeadersCiphertext: msg.HeadersCiphertext,	
			PayloadCiphertext: msg.PayloadCiphertext,	
			CreatedAt: msg.CreatedAt,
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
	accountId, _ := GetAccountInfo(r)

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




