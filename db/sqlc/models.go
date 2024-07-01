package db

import (
	"time"
)

type Address struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"user_id"`
	Coords    string    `json:"coords"`
	Street    string    `json:"street"`
	Ext       string    `json:"ext"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	ZipCode   string    `json:"zip_code"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Bonuse struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Branch struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Coords      string    `json:"coords"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Deduction struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Invoice struct {
	ID            int32     `json:"id"`
	TransactionID int32     `json:"transaction_id"`
	UserID        int32     `json:"user_id"`
	Total         float64   `json:"total"`
	Subtotal      float64   `json:"subtotal"`
	Shipping      float64   `json:"shipping"`
	Taxes         float64   `json:"taxes"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type InvoiceProduct struct {
	ID        int32     `json:"id"`
	ProductID int32     `json:"product_id"`
	InvoiceID int32     `json:"invoice_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Module struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Password struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"user_id"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Payroll struct {
	ID              int32     `json:"id"`
	UserID          int32     `json:"user_id"`
	PayrollPeriodID int32     `json:"payroll_period_id"`
	GrossSalary     float64   `json:"gross_salary"`
	NetSalary       float64   `json:"net_salary"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type PayrollBonuses struct {
	ID        int32     `json:"id"`
	PayrollID int32     `json:"payroll_id"`
	BonusID   int32     `json:"bonus_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PayrollDeduction struct {
	ID          int32     `json:"id"`
	PayrollID   int32     `json:"payroll_id"`
	DeductionID int32     `json:"deduction_id"`
	Amount      float64   `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PayrollPeriod struct {
	ID        int32     `json:"id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Product struct {
	ID           int32     `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Stock        int32     `json:"stock"`
	Price        float64   `json:"price"`
	RegularPrice float64   `json:"regular_price"`
	Weight       float64   `json:"weight"`
	Unit         string    `json:"unit"`
	BranchID     int32     `json:"branch_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Role struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RoleModule struct {
	ID        int32     `json:"id"`
	RoleID    int32     `json:"role_id"`
	ModuleID  int32     `json:"module_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Transaction struct {
	ID       int32 `json:"id"`
	From     int32 `json:"from"`
	Quantity int32 `json:"quantity"`
	To       int32 `json:"to"`
}

type Transfer struct {
	ID        int32     `json:"id"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID                    int32     `json:"id"`
	UserTypeID            int32     `json:"user_type_id"`
	Name                  string    `json:"name"`
	FirstLastName         string    `json:"first_last_name"`
	SecondLastName        string    `json:"second_last_name"`
	Email                 string    `json:"email"`
	Age                   int32     `json:"age"`
	Phone                 string    `json:"phone"`
	Username              string    `json:"username"`
	Avatar                string    `json:"avatar"`
	CellphoneVerification bool      `json:"cellphone_verification"`
	Salary                float64   `json:"salary"`
	Deleted               bool      `json:"deleted"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	Role                  Role      `json:"role"`
	UserType              UserType  `json:"user_type"`
}

type UserType struct {
	ID        int32     `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
