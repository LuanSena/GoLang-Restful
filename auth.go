package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

var validToken = "itsover8000"

func Authenticate(w http.ResponseWriter, r *http.Request) int {
	var status int
	var res []byte
	var err error
	params := r.URL.Query()

	token := strings.Join(params["Token"], "")
	if token == validToken {
		status = http.StatusOK
	} else {
		status = http.StatusUnauthorized
		log.Printf(token)
		w.WriteHeader(status)
		res, err = json.Marshal(jsonErr{Code: http.StatusUnauthorized, Text: "Unauthorized"})
		if err != nil {
			panic(err)
		}
		w.Write(res)
	}

	return status
}
