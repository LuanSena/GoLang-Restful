package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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

	var res []byte
	var err error

	authenticate := Authenticate(w, r)
	if authenticate == http.StatusUnauthorized {
		return
	}

	params := r.URL.Query()
	invoices, err := GetAllInvoices(params)
	if err != nil {
		panic(err)
	}
	res, err = json.Marshal(invoices)
	if err != nil {
		panic(err)
	}

	if invoices == nil {
		w.WriteHeader(http.StatusNotFound)
		res, err = json.Marshal(jsonErr{Code: http.StatusNotFound, Text: "Not Found"})
		if err != nil {
			panic(err)
		}
	}

	w.Write(res)
}

func invoiceInclude(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var err error
	var res []byte

	authenticate := Authenticate(w, r)
	if authenticate == http.StatusUnauthorized {
		return
	}

	invoice := new(Invoice)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &invoice)
	if err != nil {
		panic(err)
	}

	inv, err := GetInvoiceByDoc(invoice.Document)

	if inv != (Invoice{}) {
		w.WriteHeader(http.StatusConflict)
		res, err = json.Marshal(jsonErr{Code: http.StatusConflict,
			Text: "Este documento já existe"})
		if err != nil {
			panic(err)
		}
	} else {
		inv, err := CreateInvoice(invoice)
		if err != nil {
			panic(err)
		}
		invoiceCreated, err := GetInvoiceByDoc(inv.Document)
		if err != nil {
			panic(err)
		}
		res, err = json.Marshal(invoiceCreated)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusCreated)
	}

	w.Write(res)
}

func getInvoiceBydoc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	var res []byte
	var err error

	authenticate := Authenticate(w, r)
	if authenticate == http.StatusUnauthorized {
		return
	}

	var document int

	document, err = strconv.Atoi(vars["document"])
	if err != nil {
		panic(err)
	}

	invoice, err := GetInvoiceByDoc(document)
	if err != nil {
		panic(err)
	}

	if invoice != (Invoice{}) {
		res, err = json.Marshal(invoice)
		if err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		res, err = json.Marshal(jsonErr{Code: http.StatusNotFound, Text: "Not Found"})
	}

	w.Write(res)
}

func invoiceDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	var res []byte
	var err error

	authenticate := Authenticate(w, r)
	if authenticate == http.StatusUnauthorized {
		return
	}

	var document int

	document, err = strconv.Atoi(vars["document"])
	if err != nil {
		panic(err)
	}

	invoice, err := GetInvoiceByDoc(document)
	if err != nil {
		panic(err)
	}

	if invoice == (Invoice{}) {
		w.WriteHeader(http.StatusNotFound)
		res, err = json.Marshal(jsonErr{Code: http.StatusNotFound, Text: "Not Found"})
	} else {

		deleted, err := DeleteInvoice(document)
		if err != nil {
			panic(err)
		}
		if deleted == "deleted" {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

	}
	w.Write(res)
}
