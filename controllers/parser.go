package controllers

import (
        //jwt "github.com/dgrijalva/jwt-go"
        //"github.com/xlab-si/e2ee-server/core/authentication"
        "net/http"
	"github.com/jeffail/gabs"
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

func GetAccountInfo(r *http.Request) (string, string) {
 	auth := r.Header.Get("Authorization")
	if (strings.Fields(auth)[1] == "thisistesttoken") {
		// this works only when authorization is not required -
		// when authorization is required, such tokens are rejected
		// prior to this call
		return strings.Fields(auth)[3], strings.Fields(auth)[2]
	}
        auth_token := strings.Fields(auth)[1]

	url := "https://www.googleapis.com/oauth2/v3/userinfo?access_token="
        response, err := http.Get(url+auth_token)

        if err != nil {
                fmt.Printf("%s", err)
                os.Exit(1)
        } else {
                defer response.Body.Close()
                contents, err := ioutil.ReadAll(response.Body)
                if err != nil {
                        fmt.Printf("%s", err)
                        os.Exit(1)
                }
                fmt.Printf("%s\n", string(contents))

		j, err1 := gabs.ParseJSON([]byte(string(contents)))
		if err1 != nil {
			fmt.Println(err1)
		}
		id := j.Path("sub").Data().(string)
		email := j.Path("email").Data().(string)
		return id, email
	}
	return "", ""
}

