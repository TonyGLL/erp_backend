CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "role_id" integer NOT NULL REFERENCES "roles" ("id"),
  "user_type_id" integer NOT NULL REFERENCES "user_types" ("id"),
  "name" varchar(30) NOT NULL,
  "first_last_name" varchar(30) NOT NULL,
  "second_last_name" varchar(30) NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "age" integer NOT NULL,
  "phone" varchar(15) NOT NULL,
  "username" varchar(20) UNIQUE NOT NULL,
  "avatar" text NOT NULL,
  "cellphone_verification" boolean NOT NULL DEFAULT false,
  "salary" float NOT NULL,
  "deleted" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now()
);

CREATE TABLE "user_types" (
  "id" SERIAL PRIMARY KEY,
  "type" varchar(20) UNIQUE NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now()
);

CREATE TABLE "passwords" (
  "id" SERIAL PRIMARY KEY,
  "user_id" integer REFERENCES "users" ("id"),
  "value" varchar(255),
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now()
);

CREATE TABLE "addresses" (
  "id" SERIAL PRIMARY KEY,
  "user_id" integer REFERENCES "users" ("id"),
  "coords" varchar NOT NULL,
  "street" varchar(30) NOT NULL,
  "ext" varchar(30) NOT NULL,
  "city" varchar(30) NOT NULL,
  "state" varchar(30) NOT NULL,
  "zip_code" varchar(5) NOT NULL,
  "country" varchar(30) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now()
);

CREATE TABLE "roles" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now()
);

CREATE TABLE "modules" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now()
);

CREATE TABLE "role_modules" (
  "id" SERIAL PRIMARY KEY,
  "role_id" integer REFERENCES "roles" ("id"),
  "module_id" integer REFERENCES "modules" ("id"),
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now()
);

CREATE TABLE "transfers" (
  "id" SERIAL PRIMARY KEY,
  "token" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now()
);

CREATE TABLE "invoice" (
  "id" SERIAL PRIMARY KEY,
  "transaction_id" integer REFERENCES "transfers" ("id"),
  "user_id" integer REFERENCES "users" ("id"),
  "total" float NOT NULL,
  "subtotal" float NOT NULL,
  "shipping" float NOT NULL,
  "taxes" float NOT NULL,
  "description" text NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now()
);

CREATE TABLE "invoice_products" (
  "id" SERIAL PRIMARY KEY,
  "product_id" integer REFERENCES "products" ("id"),
  "invoice_id" integer REFERENCES "invoice" ("id"),
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now()
);

CREATE TABLE "products" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "description" varchar(255),
  "stock" integer NOT NULL,
  "price" float NOT NULL,
  "regular_price" float NOT NULL,
  "weight" float NOT NULL,
  "unit" varchar(3) NOT NULL,
  "branch_id" integer REFERENCES "branches" ("id"),
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now()
);

CREATE TABLE "branches" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "description" varchar(255),
  "coords" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now()
);

CREATE TABLE "transactions" (
  "id" SERIAL PRIMARY KEY,
  "from" integer REFERENCES "branches" ("id"),
  "quantity" integer NOT NULL,
  "to" integer REFERENCES "branches" ("id")
);

CREATE TABLE "payroll_periods" (
  "id" SERIAL PRIMARY KEY,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "status" varchar(20) NOT NULL DEFAULT 'pending',
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now()
);

CREATE TABLE "payrolls" (
  "id" SERIAL PRIMARY KEY,
  "user_id" integer REFERENCES "users" ("id"),
  "payroll_period_id" integer REFERENCES "payroll_periods" ("id"),
  "gross_salary" float NOT NULL,
  "net_salary" float NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now()
);

CREATE TABLE "deductions" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "description" text,
  "amount" float NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now()
);

CREATE TABLE "bonuses" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "description" text,
  "amount" float NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now()
);

CREATE TABLE "payroll_deductions" (
  "id" SERIAL PRIMARY KEY,
  "payroll_id" integer REFERENCES "payrolls" ("id"),
  "deduction_id" integer REFERENCES "deductions" ("id"),
  "amount" float NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now()
);

CREATE TABLE "payroll_bonuses" (
  "id" SERIAL PRIMARY KEY,
  "payroll_id" integer REFERENCES "payrolls" ("id"),
  "bonus_id" integer REFERENCES "bonuses" ("id"),
  "amount" float NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now()
);
