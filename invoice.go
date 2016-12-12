package main

import (
	"time"
)

type Invoice struct {
	Document        int       `json:"document,string"`
	Description     string    `json:"description"`
	Amount          float64   `json:"amount"`
	ReferenceMounth int       `json:"reference_mounth"`
	Id_invoice      int       `json:"id_invoice"`
	ReferenceYear   int       `json:"reference_year"`
	CreatedAt       time.Time `json:"created_at"`
	IsActive        byte      `json:"is_active"`
	DesactiveAt     string    `json:"desactive_at"`
}

type Invoices []Invoice
