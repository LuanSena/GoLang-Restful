create database desafio;
use desafio;
CREATE TABLE invoices (
	id_invoice INT UNIQUE AUTO_INCREMENT,
	document VARCHAR(14) NOT NULL,
	description VARCHAR(256) NOT NULL,
	amount DECIMAL(16,2) NOT NULL,
	reference_month INT NOT NULL,
	reference_year INT NOT NULL,
	created_at DATETIME NOT NULL,
	is_active TINYINT NOT NULL,
	desactive_at DATETIME,
	PRIMARY KEY (document)
);
