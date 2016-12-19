# GoLang Restful
[![Portugues](https://img.shields.io/badge/README-PT--BR-blue.svg)](https://github.com/LuanSena/GoLang-Restful-Invoice/blob/master/README.md)
[![English](https://img.shields.io/badge/README-ENGLISH-blue.svg)](https://github.com/LuanSena/GoLang-Restful-Invoice/blob/master/README_eng.md)
## Como utilizar:
**é necessario criar o banco de dados e tabela no Mysql,
  script disponivel no arquivo SQL/createDB.sql**

- git clone https://github.com/LuanSena/GoLang-Restful-Invoice.git
- cd source
- configure o arquivo database.go (usuario:senha@/nomeDoBanco)
- go build
- ./GoLang-Restful-Invoice


**o Token de autenticação pode ser localizado no arquivo auth.go**

### Rotas: ###
**https://127.0.0.1:8080/invoices**

- Descrição: Retorna invoices por filtro
- Parametros: Token, orderBy, year, month, limit, offset
- Metodo: GET

**https://127.0.0.1:8080/invoices**

- Descrição: Cria novo invoice
- Parametros: Token, document, description, amount, reference_month, reference_year
- Metodo: POST

**https://127.0.0.1:8080/invoices/{document}**

- Descrição: Retorna invoice por Documento
- Parametros: Token, document
- Metodo: GET

**https://127.0.0.1:8080/invoices/{document}**

- Descrição: Desativa invoice por documento
- Parametros: Token, document
- Metodo: DELETE

**Relacionamento entre arquivos**
![CSCore Logo](https://github.com/LuanSena/LuanSena.github.io/blob/master/img/1481812193122637.PNG)
