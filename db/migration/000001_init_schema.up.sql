-- Create schema
CREATE SCHEMA erp_schema;

-- Tables creation
CREATE TABLE erp_schema.users (
  id SERIAL PRIMARY KEY,
  role_id integer NOT NULL,
  name varchar(30) NOT NULL,
  first_last_name varchar(30) NOT NULL,
  second_last_name varchar(30) NOT NULL,
  email varchar UNIQUE NOT NULL,
  age integer NOT NULL,
  phone varchar(15) NOT NULL,
  username varchar(20) UNIQUE NOT NULL,
  avatar text NOT NULL,
  cellphone_verification boolean NOT NULL DEFAULT false,
  salary float NOT NULL,
  deleted boolean NOT NULL DEFAULT false,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE erp_schema.passwords (
  id SERIAL PRIMARY KEY,
  user_id integer,
  value varchar(255),
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE erp_schema.addresses (
  id SERIAL PRIMARY KEY,
  user_id integer,
  coords varchar NOT NULL,
  street varchar(30) NOT NULL,
  ext varchar(30) NOT NULL,
  city varchar(30) NOT NULL,
  state varchar(30) NOT NULL,
  zip_code varchar(5) NOT NULL,
  country varchar(30) NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE erp_schema.roles (
  id SERIAL PRIMARY KEY,
  name varchar NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE erp_schema.modules (
  id SERIAL PRIMARY KEY,
  name varchar NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE erp_schema.role_modules (
  id SERIAL PRIMARY KEY,
  role_id integer,
  module_id integer,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE erp_schema.transfers (
  id SERIAL PRIMARY KEY,
  token varchar NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE erp_schema.invoice (
  id SERIAL PRIMARY KEY,
  transaction_id integer,
  user_id integer,
  total float NOT NULL,
  subtotal float NOT NULL,
  shipping float NOT NULL,
  taxes float NOT NULL,
  description text NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE erp_schema.invoice_products (
  id SERIAL PRIMARY KEY,
  product_id integer,
  invoice_id integer,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE erp_schema.products (
  id SERIAL PRIMARY KEY,
  name varchar(100) NOT NULL,
  description varchar(255),
  stock integer NOT NULL,
  price float NOT NULL,
  regular_price float NOT NULL,
  weight float NOT NULL,
  unit varchar(3) NOT NULL,
  branch_id integer,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE erp_schema.branches (
  id SERIAL PRIMARY KEY,
  name varchar(100) NOT NULL,
  description varchar(255),
  coords varchar NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE erp_schema.transactions (
  id SERIAL PRIMARY KEY,
  tx_from integer NOT NULL,
  quantity integer NOT NULL,
  tx_to integer NOT NULL
);

CREATE TABLE erp_schema.payroll_periods (
  id SERIAL PRIMARY KEY,
  start_date date NOT NULL,
  end_date date NOT NULL,
  status varchar(20) NOT NULL DEFAULT 'pending',
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE erp_schema.payrolls (
  id SERIAL PRIMARY KEY,
  user_id integer,
  payroll_period_id integer,
  gross_salary float NOT NULL,
  net_salary float NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE erp_schema.deductions (
  id SERIAL PRIMARY KEY,
  name varchar(100) NOT NULL,
  description text,
  amount float NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE erp_schema.bonuses (
  id SERIAL PRIMARY KEY,
  name varchar(100) NOT NULL,
  description text,
  amount float NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE erp_schema.payroll_deductions (
  id SERIAL PRIMARY KEY,
  payroll_id integer,
  deduction_id integer,
  amount float NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE erp_schema.payroll_bonuses (
  id SERIAL PRIMARY KEY,
  payroll_id integer,
  bonus_id integer,
  amount float NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

-- Add foreign keys
ALTER TABLE erp_schema.users
  ADD CONSTRAINT fk_role_id FOREIGN KEY (role_id) REFERENCES erp_schema.roles (id);

ALTER TABLE erp_schema.passwords
  ADD CONSTRAINT fk_password_user_id FOREIGN KEY (user_id) REFERENCES erp_schema.users (id);

ALTER TABLE erp_schema.addresses
  ADD CONSTRAINT fk_address_user_id FOREIGN KEY (user_id) REFERENCES erp_schema.users (id);

ALTER TABLE erp_schema.role_modules
  ADD CONSTRAINT fk_role_module_role_id FOREIGN KEY (role_id) REFERENCES erp_schema.roles (id),
  ADD CONSTRAINT fk_role_module_module_id FOREIGN KEY (module_id) REFERENCES erp_schema.modules (id);

ALTER TABLE erp_schema.invoice
  ADD CONSTRAINT fk_invoice_transaction_id FOREIGN KEY (transaction_id) REFERENCES erp_schema.transfers (id),
  ADD CONSTRAINT fk_invoice_user_id FOREIGN KEY (user_id) REFERENCES erp_schema.users (id);

ALTER TABLE erp_schema.invoice_products
  ADD CONSTRAINT fk_invoice_product_product_id FOREIGN KEY (product_id) REFERENCES erp_schema.products (id),
  ADD CONSTRAINT fk_invoice_product_invoice_id FOREIGN KEY (invoice_id) REFERENCES erp_schema.invoice (id);

ALTER TABLE erp_schema.products
  ADD CONSTRAINT fk_product_branch_id FOREIGN KEY (branch_id) REFERENCES erp_schema.branches (id);

ALTER TABLE erp_schema.transactions
  ADD CONSTRAINT fk_transaction_from FOREIGN KEY ("tx_from") REFERENCES erp_schema.branches (id),
  ADD CONSTRAINT fk_transaction_to FOREIGN KEY ("tx_to") REFERENCES erp_schema.branches (id);

ALTER TABLE erp_schema.payrolls
  ADD CONSTRAINT fk_payroll_user_id FOREIGN KEY (user_id) REFERENCES erp_schema.users (id),
  ADD CONSTRAINT fk_payroll_period_id FOREIGN KEY (payroll_period_id) REFERENCES erp_schema.payroll_periods (id);

ALTER TABLE erp_schema.payroll_deductions
  ADD CONSTRAINT fk_payroll_deduction_payroll_id FOREIGN KEY (payroll_id) REFERENCES erp_schema.payrolls (id),
  ADD CONSTRAINT fk_payroll_deduction_deduction_id FOREIGN KEY (deduction_id) REFERENCES erp_schema.deductions (id);

ALTER TABLE erp_schema.payroll_bonuses
  ADD CONSTRAINT fk_payroll_bonus_payroll_id FOREIGN KEY (payroll_id) REFERENCES erp_schema.payrolls (id),
  ADD CONSTRAINT fk_payroll_bonus_bonus_id FOREIGN KEY (bonus_id) REFERENCES erp_schema.bonuses (id);
