package controllers

import (
	"e2ee/core/db"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	//"github.com/gorilla/mux"
	"log"
	//"github.com/jeffail/gabs"
)

type Peer struct {
    AccountId uint `json:"accountId"`
    Username string `json:"username"`
    PubKey string `json:"pubKey"`
    SignKeyPub string `json:"signKeyPub"`
}

type PeerMessage struct {
    Success bool `json:"success"`
    Peer Peer `json:"peer"`
}

type Notification struct {
    FromUsername string `json:"fromUsername"`
    ToAccountId uint `json:"toAccountId"`
    HeadersCiphertext string `json:"headersCiphertext"`
    PayloadCiphertext string `json:"payloadCiphertext"`
}

type NotificationResponse struct {
    Success bool `json:"success"`
    Error string `json:"error"`
    MessageId uint `json:"messageId"`
}

func PeerGet(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
        p := strings.SplitN(r.URL.RequestURI()[1:], "/", 3)
        peerName := p[1]
	log.Println(peerName)
	account := db.FindAccountByName(peerName)
	success := false 
	if account.AccountId != 0 {
	    success = true
	}

	var peer = Peer{
	    AccountId: account.AccountId,
	    Username: account.Username,
	    PubKey: account.PubKey,	
	    SignKeyPub: account.SignKeyPub,
	}
	var m = PeerMessage{
	    Success: success,
	    Peer: peer,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m); err != nil {
                panic(err)
        }
}

func PeerNotify(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_, _, accountId := ExtractTokenInfo(r)

        body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
        log.Println(err)
	var chunk Notification

	if err != nil {
                panic(err)
        }
        if err := r.Body.Close(); err != nil {
                panic(err)
        }
        if err := json.Unmarshal(body, &chunk); err != nil {
                w.Header().Set("Content-Type", "application/json; charset=UTF-8")
                w.WriteHeader(422) // unprocessable entity
                if err := json.NewEncoder(w).Encode(err); err != nil {
                        panic(err)
                }
                return
        }

	messageId := db.CreateNotification(accountId, chunk.ToAccountId, chunk.HeadersCiphertext, chunk.PayloadCiphertext)

	success := false 
	if messageId != 0 {
	    success = true
	}

	var m = NotificationResponse{
	    Success: success,
	    Error: "",
	    MessageId: messageId,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m); err != nil {
                panic(err)
        }
}





