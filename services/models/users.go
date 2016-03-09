package models

type User struct {
	UUID     string `json:"uuid" form:"-"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	AccountId int `json:"accountId" form:"accountId"`
	Token string `json:"token" form:"token"`
}
