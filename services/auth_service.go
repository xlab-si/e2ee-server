package services

import (
	"e2ee/api/parameters"
	"e2ee/core/authentication"
	"e2ee/core/db"
	"encoding/json"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
)

func Login(credentials db.UserCredentials) (int, []byte) {
	authBackend := authentication.InitJWTAuthenticationBackend()
	dbUser := db.GetUser(credentials.Username) // todo: err mngt

	if authBackend.Authenticate(dbUser, credentials) {
		var err error
		token := dbUser.Token
		if token == "" {
			token, err = authBackend.GenerateToken(dbUser.Uuid, dbUser.Username, dbUser.ID)
			if err != nil {
			    return http.StatusInternalServerError, []byte("")
			}
			db.AddToken(dbUser.Username, token)
		}
		response, _ := json.Marshal(parameters.TokenAuthentication{token})
		return http.StatusOK, response
	}

	return http.StatusUnauthorized, []byte("")
}

func RefreshToken(uuid string, username string, accountId uint) []byte {
	authBackend := authentication.InitJWTAuthenticationBackend()
	token, err := authBackend.GenerateToken(uuid, username, accountId)
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(parameters.TokenAuthentication{token})
	if err != nil {
		panic(err)
	}
	return response
}

func Logout(req *http.Request) error {
	authBackend := authentication.InitJWTAuthenticationBackend()
	tokenRequest, err := jwt.ParseFromRequest(req, func(token *jwt.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})
	if err != nil {
		return err
	}
	tokenString := req.Header.Get("Authorization")
	return authBackend.Logout(tokenString, tokenRequest)
}
