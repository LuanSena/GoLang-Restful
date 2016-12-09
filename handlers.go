package main

import(
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to invoice RESFUL Api")
}

func invoiceIndex(w http.ResponseWriter, r *http.Request) {
		invoice := Invoices{
        Invoice{Document: 1},
        Invoice{Document: 2},
    }
    err := json.NewEncoder(w).Encode(invoice)
    if err != nil{
    	panic(err)
    }

}

func invoiceShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	invoiceDoc := vars["invoiceDoc"]
	fmt.Fprintln(w, "Numero do documento:", invoiceDoc)	
}