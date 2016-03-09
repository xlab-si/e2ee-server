package db

import (
	"github.com/jinzhu/gorm"
)

type Message struct {
	gorm.Model
	FromAccountId      uint    `json:"fromAccountId"`
	ToAccountId      uint    `json:"toAccountId"`
	HeadersCiphertext      string    `json:"headersCiphertext" sql:"type:text"`
	PayloadCiphertext      string    `json:"payloadCiphertext" sql:"type:text"`
}


