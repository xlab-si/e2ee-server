package services

type Account struct {
	ContainerNameHmacKeyCiphertext        string       `json:"containerNameHmacKeyCiphertext"`
	HmacKeyCiphertext      string    `json:"hmacKeyCiphertext"`
	KeypairCiphertext string      `json:"keypairCiphertext"`
	KeypairMac string      `json:"keypairMac"`
	KeypairMacSalt string      `json:"keypairMacSalt"`
	KeypairSalt string      `json:"keypairSalt"`
	PubKey string      `json:"pubKey"`
	SignKeyPrivateCiphertext string      `json:"signKeyPrivateCiphertext"`
	SignKeyPrivateMac string      `json:"signKeyPrivateMac"`
	SignKeyPrivateMacSalt string      `json:"signKeyPrivateMacSalt"`
	SignKeyPub string      `json:"signKeyPub"`
	Username string      `json:"username"`
}

/*

import (
        "encoding/json"
        "io"
        "io/ioutil"

        //"log"
        //"github.com/jeffail/gabs"
)

type AccountResponseMessage struct {
    Success bool `json:"success"`
    Account Account `json:"account"`
    Sid int `json:"sid"`
}

func Account AccountGet() {

}
*/






