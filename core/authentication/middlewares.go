package authentication

import (
	//jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"encoding/json"
	"strings"
	"os"
	"io/ioutil"
    	"fmt"
	"golang.org/x/oauth2/jws"
)

const clientID string = "250838009887-b3i9oidc73koc3u3rulat30gpma20qmn.apps.googleusercontent.com"

func RequireTokenAuthentication(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
    	auth := req.Header.Get("Authorization")
    	auth_token := strings.Fields(auth)[1]

	// TODO: validation of token should be offline
	url := "https://www.googleapis.com/oauth2/v3/tokeninfo?access_token="
	response, err := http.Get(url+auth_token)

	var set jws.ClaimSet
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
	    	err = json.Unmarshal([]byte(string(contents)), &set)
		if (set.Aud == clientID) {
			next(rw, req)
		} else {
			rw.WriteHeader(http.StatusUnauthorized)
		}
    	}
}
