# GoLang Restful
[![Portugues](https://img.shields.io/badge/README-PT--BR-blue.svg)](https://github.com/LuanSena/GoLang-Restful-Invoice/blob/master/README.md)
[![English](https://img.shields.io/badge/README-ENGLISH-blue.svg)](https://github.com/LuanSena/GoLang-Restful-Invoice/blob/master/README_eng.md)
## How to use:
**It's necessary to create the database and table in Mysql,
  Script available in file createDB.sql**

- git clone https://github.com/LuanSena/GoLang-Restful-Invoice.git
- cd GoLang-Restful-Invoice
- configure the database.go file (user:pass@/dataBaseName)
- go build
- ./GoLang-Restful-Invoice


**The authentication token can be find in the file auth.go**

### Routes: ###
**https://127.0.0.1:8080/invoices**

- Description: Return invoices by filters
- Parameters: Token, orderBy, year, month, limit, offset
- Method: GET

**https://127.0.0.1:8080/invoices**

- Description: Creates a new invoice
- Parameters: Token, document, description, amount, reference_month, reference_year
- Method: POST

**https://127.0.0.1:8080/invoices/{document}**

- Description: Return an invoice by Document
- Parameters: Token, document
- Method: GET

**https://127.0.0.1:8080/invoices/{document}**

- Description: Deactivates an invoice by document
- Parameters: Token, document
- Method: DELETE

**Relacionamento entre arquivos**
![CSCore Logo](https://github.com/LuanSena/LuanSena.github.io/blob/master/img/1481812193122637.PNG)
