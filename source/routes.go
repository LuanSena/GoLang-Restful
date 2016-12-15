package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"index",
		"GET",
		"/",
		index,
	},
	Route{
		"getAllInvoices",
		"GET",
		"/invoices",
		getAllInvoices,
	},
	Route{
		"invoiceInclude",
		"POST",
		"/invoices",
		invoiceInclude,
	},
	Route{
		"getInvoiceBydoc",
		"GET",
		"/invoices/{document}",
		getInvoiceBydoc,
	},
	Route{
		"invoiceDelete",
		"DELETE",
		"/invoices/{document}",
		invoiceDelete,
	},
}
