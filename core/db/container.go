package db

import (
	"github.com/jinzhu/gorm"
)

type Container struct {
	gorm.Model
	AccountId      uint    `json:"accountId"`
	ContainerNameHmac      string    `json:"containerNameHmac"`
	LatestRecordId      int    `json:"latestRecordId"`
}

type ContainerRecord struct {
	gorm.Model
	ContainerId      uint    `json:"containerId"`
	AccountId      uint    `json:"accountId"`
	PayloadCiphertext      string    `json:"payloadCiphertext" sql:"type:text"`
	SessionKeyCiphertext      string    `json:"sessionKeyCiphertext" sql:"type:text"` // set only when returning records to the client
}

type ContainerSessionKeyShare struct {
	gorm.Model
	//ContainerSessionKeyId      int    `json:"containerSessionKeyId"`
	ContainerId      uint    `json:"containerId"` // experimenting without SessionKey
	AccountId      uint    `json:"accountId"`
	ToAccountId      uint    `json:"toAccountId"`
	SessionKeyCiphertext      string    `json:"sessionKeyCiphertext" sql:"type:text"`
}

