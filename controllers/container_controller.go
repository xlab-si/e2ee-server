package controllers

import (
	"github.com/xlab-si/e2ee-server/core/db"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	//"github.com/gorilla/mux"
	//"log"
	//"github.com/jeffail/gabs"
)

type ContainerResponseMessage struct {
    Success bool `json:"success"`
    Error string `json:"error"`
    Records []db.ContainerRecord `json:"records"`
}

type ContainerCreateChunk struct {
    ToAccountId int `json:"toAccountId"`
    SessionKeyCiphertext string `json:"sessionKeyCiphertext"`
}

type RecordCreateChunk struct {
    ContainerNameHmac string `json:"containerNameHmac"`
    PayloadCiphertext string `json:"payloadCiphertext"`
}

type ContainerShareChunk struct {
    ContainerNameHmac string `json:"containerNameHmac"`
    ToAccountId uint `json:"toAccountId"`
    SessionKeyCiphertext string `json:"sessionKeyCiphertext"`
}

type ContainerUnshareChunk struct {
    ContainerNameHmac string `json:"containerNameHmac"`
    ToAccountId uint `json:"toAccountId"`
}

func ContainerGet(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_, _, accountId := ExtractTokenInfo(r)
        p := strings.SplitN(r.URL.RequestURI()[1:], "/", 3)
        p0 := strings.SplitN(p[1], "?", 3)
	containerNameHmac := p0[0]

	container := db.FindContainer(containerNameHmac)
	var err1 = ""
	var success bool
	success = true
 	if container == (db.Container{}) {
	    err1 = "container does not exist"	
	    success = false
	} 

	records := db.GetContainerRecords(container.ID, accountId)
	
	var m = ContainerResponseMessage{
	    Success: success,
	    Error: err1,
	    Records: records,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m); err != nil {
		panic(err)
	}
}

func ContainerCreate(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_, _, accountId := ExtractTokenInfo(r)
        p := strings.SplitN(r.URL.RequestURI()[1:], "/", 3)
	containerNameHmac := p[1]

	var chunk ContainerCreateChunk
        body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	//jsonParsed, err1 := gabs.ParseJSON([]byte(body))
        //log.Println(err1)
        //log.Println(jsonParsed)

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

	var sessionKeyCiphertext string
	sessionKeyCiphertext = chunk.SessionKeyCiphertext

	db.CreateContainer(accountId, containerNameHmac)
	//log.Println("Container created: " + strconv.Itoa(containerId))
	db.CreateContainerSessionKeyShare(containerNameHmac, sessionKeyCiphertext, accountId, accountId)
	
	var m = ContainerResponseMessage{
	    Success: true,
	    Error: "",
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m); err != nil {
		panic(err)
	}
}

func ContainerDelete(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	//_, _, accountId := ExtractTokenInfo(r)
        p := strings.SplitN(r.URL.RequestURI()[1:], "/", 3)
	containerNameHmac := p[1]
	
        db.DeleteContainer(containerNameHmac)
	
	var m = ContainerResponseMessage{
	    Success: true,
	    Error: "",
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m); err != nil {
		panic(err)
	}
}

func ContainerRecordCreate(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_, _, accountId := ExtractTokenInfo(r)
	var chunk RecordCreateChunk
        //body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	// records might contain large files, thus no limitation is 
 	// applied here, perhaps some large value might be applied?
        body, err := ioutil.ReadAll(r.Body)

	//jsonParsed, err2 := gabs.ParseJSON([]byte(body))
        //log.Println(err2)
        //log.Println(jsonParsed)

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
	
	container := db.FindContainer(chunk.ContainerNameHmac)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var err1 string
	if container == (db.Container{}) {
		err1 = "container does not exist"	
		var m = ContainerResponseMessage{
	    		Success: false,
	    		Error: err1,
		}
		if err := json.NewEncoder(w).Encode(m); err != nil {
			panic(err)
		}
	} else {
		db.CreateContainerRecord(container.ID, accountId, chunk.PayloadCiphertext)
		var m = ContainerResponseMessage{
	    		Success: true,
	    		Error: "",
		}
		if err := json.NewEncoder(w).Encode(m); err != nil {
			panic(err)
		}
	}	
}

func ContainerShare(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_, _, accountId := ExtractTokenInfo(r)

	var chunk ContainerShareChunk
        body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

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

	db.CreateContainerSessionKeyShare(chunk.ContainerNameHmac, chunk.SessionKeyCiphertext, accountId, chunk.ToAccountId)

	var m = ContainerResponseMessage{
	    Success: true,
	    Error: "",
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m); err != nil {
		panic(err)
	}
}

func ContainerUnshare(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_, _, accountId := ExtractTokenInfo(r)

	var chunk ContainerUnshareChunk
        body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

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

	db.DeleteContainerSessionKeyShare(chunk.ContainerNameHmac, accountId, chunk.ToAccountId)

	var m = ContainerResponseMessage{
	    Success: true,
	    Error: "",
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m); err != nil {
		panic(err)
	}
}





