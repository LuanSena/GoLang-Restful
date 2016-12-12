package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to invoice RESFUL Api")
}

func getAllInvoices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	params := r.URL.Query()
	fmt.Println(params)
}

func invoiceInclude(w http.ResponseWriter, r *http.Request) {
	var invoice Invoice
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &invoice); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
}

func getInvoiceBydoc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	invoiceDoc := vars["invoiceDoc"]
	fmt.Fprintln(w, "Numero do documento:", invoiceDoc)
}

func invoiceDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	invoiceDoc := vars["invoiceDoc"]
	fmt.Fprintln(w, "Numero do documento:", invoiceDoc)
}
