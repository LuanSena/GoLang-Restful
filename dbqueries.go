package main

import (
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllInvoices(params map[string][]string) (Invoices, error) {
	orderBy := strings.Join(params["orderBy"], "")
	year := strings.Join(params["year"], "")
	month := strings.Join(params["month"], "")
	limit := strings.Join(params["limit"], "")
	offset := strings.Join(params["offset"], "")

	var invoice Invoice
	var invoices Invoices

	var sql = "select * from invoices"
	//Seleção por ano/mes
	if year != "" {
		sql += "  where reference_year = " + year
		if month != "" {
			sql += "  and reference_month = " + month
		}
	} else {
		if month != "" {
			sql += " where reference_month = " + month
		}
	}

	//Ordenação da seleção
	if orderBy != "" {
		sql += " order by " + orderBy
	}
	if limit != "" {
		sql += " limit " + limit
		if offset != "" {
			sql += " offset " + offset
		}
	}
	//Coloca o resultado em ROWS
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	//Popula o slice Invoices com o retorno do BD
	for rows.Next() {
		err = rows.Scan(&invoice.Document, &invoice.Description, &invoice.Amount,
			&invoice.ReferenceMounth, &invoice.ReferenceYear,
			&invoice.CreatedAt, &invoice.IsActive, &invoice.DesactiveAt, &invoice.Id_invoice)
		invoices = append(invoices, invoice)
		if err != nil {
			panic(err)
		}
	}

	defer rows.Close()

	return invoices, err
}
