package controllers

import (
	"net/http"
	"encoding/json"
	"net/url"
)

type VersionCheckMessage struct {
    Success bool `json:"success"`
    Error string `json:"version"`
}

func VersionCheck(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	u, _ := url.Parse(r.URL.String())
	queryParams := u.Query()
	var success bool
	success = false
	if queryParams["v"][0] == "0.0.4" {
		success = true
	}
        var m = VersionCheckMessage{
            Success: success,
            Error: "",
        }
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusOK)
        if err := json.NewEncoder(w).Encode(m); err != nil {
                panic(err)
        }
}
