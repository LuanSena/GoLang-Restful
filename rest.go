package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "time"
    //external
    "github.com/gorilla/mux"
)


func main() {
	router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", index)
   	router.HandleFunc("/invoice", invoiceIndex)
   	router.HandleFunc("/invoice/{invoiceDoc}", invoiceShow)

    log.Fatal(http.ListenAndServe(":8080", router))

}

func index(w http.ResponseWriter, r *http.Request) {
	invoice := Invoices{
        Invoice{Document: 1},
        Invoice{Document: 2},
    }
    json.NewEncoder(w).Encode(invoice)
}

func invoiceIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Invoice index")
}

func invoiceShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	invoiceDoc := vars["invoiceDoc"]
	fmt.Fprintln(w, "Numero do documento:", invoiceDoc)	
}