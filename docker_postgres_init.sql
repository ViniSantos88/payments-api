CREATE TABLE IF NOT EXISTS accounts (
	account_id serial PRIMARY KEY,
	document_number VARCHAR ( 50 ) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS operations_types (
	operation_type_id serial PRIMARY KEY,
	description0 VARCHAR ( 50 ) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS transactions (
	transaction_id serial PRIMARY KEY,
  account_id INT NOT NULL,
  operation_type_id INT NOT NULL,
  amount NUMERIC(5,2) NOT NULL,
	event_date TIMESTAMP,
  FOREIGN KEY (account_id)
      REFERENCES accounts (account_id),
  FOREIGN KEY (operation_type_id)
      REFERENCES operations_types (operation_type_id)    
);

insert into operations_types (operation_type_id, description0) VALUES (1,'COMPRA A VISTA');
insert into operations_types (operation_type_id, description0) VALUES (2,'COMPRA PARCELADA');
insert into operations_types (operation_type_id, description0) VALUES (3,'SAQUE');
insert into operations_types (operation_type_id, description0) VALUES (4,'PAGAMENTO');

commit;