package main

import (
	"database/sql"
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

func GetInvoiceByDoc(document int) (Invoice, error) {
	var invoice Invoice
	stmt, err := db.Prepare("select * from invoices where document=?")
	if err != nil {
		panic(err)
	}

	err = stmt.QueryRow(document).Scan(&invoice.Document, &invoice.Description, &invoice.Amount,
		&invoice.ReferenceMounth, &invoice.ReferenceYear,
		&invoice.CreatedAt, &invoice.IsActive, &invoice.DesactiveAt)
	if err == sql.ErrNoRows {
		return invoice, err
	}
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	return invoice, err
}

func CreateInvoice(invoice *Invoice) (*Invoice, error) {
	var sql = "insert into invoices set document=?, description=?, amount=?,"
	sql += " reference_month=?, reference_year=?, created_at=NOW(), is_active=1"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(invoice.Document, invoice.Description, invoice.Amount,
		invoice.ReferenceMounth, invoice.ReferenceYear)
	if err != nil {
		return invoice, err
	}

	defer stmt.Close()

	return invoice, err
}

func DeleteInvoice(document int) (string, error) {
	stmt, err := db.Prepare("update invoices set is_active=0, desactive_at=NOW() where document = ?")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(document)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	return "De Actived", err
}
