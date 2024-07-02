package db

import (
	"context"
	"time"
)

const getUser = `
SELECT
    u.id,
    u.name,
    u.first_last_name,
    u.second_last_name,
    u.email,
    u.age,
    u.phone,
    u.username,
    u.avatar,
    u.cellphone_verification,
    u.salary,
    u.deleted,
    u.created_at,
    u.updated_at,
    r.id AS role_id,
    r.name AS role_name,
    ut.id AS user_type_id,
    ut.type AS user_type_name
FROM erp_schema.users u
LEFT JOIN erp_schema.roles r ON u.role_id = r.id
LEFT JOIN erp_schema.user_types ut ON u.user_type_id = ut.id
WHERE u.id = $1 AND u.deleted IS NOT TRUE LIMIT 1
`

type GetUserRow struct {
	ID                    int32     `json:"id"`
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
	RoleID                int32     `json:"role_id"`
	RoleName              string    `json:"role_name"`
	UserTypeID            int32     `json:"user_type_id"`
	UserTypeName          string    `json:"user_type_name"`
}

func (q *Queries) GetUser(ctx context.Context, id int32) (GetUserRow, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var user GetUserRow
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.FirstLastName,
		&user.SecondLastName,
		&user.Email,
		&user.Age,
		&user.Phone,
		&user.Username,
		&user.Avatar,
		&user.CellphoneVerification,
		&user.Salary,
		&user.Deleted,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.RoleID,
		&user.RoleName,
		&user.UserTypeID,
		&user.UserTypeName,
	)
	return user, err
}

const getUsers = `
SELECT
    u.id,
    u.name,
    u.first_last_name,
    u.second_last_name,
    u.email,
    u.age,
    u.phone,
    u.username,
    u.avatar,
    u.cellphone_verification,
    u.salary,
    u.deleted,
    u.created_at,
    u.updated_at,
    r.id AS role_id,
    r.name AS role_name,
    ut.id AS user_type_id,
    ut.type AS user_type_name
FROM erp_schema.users u
LEFT JOIN erp_schema.roles r ON u.role_id = r.id
LEFT JOIN erp_schema.user_types ut ON u.user_type_id = ut.id
WHERE u.deleted IS NOT TRUE
ORDER BY u.id
LIMIT $1 OFFSET $2
`

type GetUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetUsers(ctx context.Context, arg GetUsersParams) ([]GetUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetUserRow{}
	for rows.Next() {
		var user GetUserRow
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.FirstLastName,
			&user.SecondLastName,
			&user.Email,
			&user.Age,
			&user.Phone,
			&user.Username,
			&user.Avatar,
			&user.CellphoneVerification,
			&user.Salary,
			&user.Deleted,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.RoleID,
			&user.RoleName,
			&user.UserTypeID,
			&user.UserTypeName,
		); err != nil {
			return nil, err
		}
		items = append(items, user)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const CountUsers = `
SELECT COUNT(*) AS total_users
FROM erp_schema.users u
WHERE u.deleted IS NOT TRUE;
`

func (q *Queries) CountUsers(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, CountUsers)
	var total_accounts int64
	err := row.Scan(&total_accounts)
	return total_accounts, err
}

const createUser = `
INSERT INTO erp_schema.users (role_id, user_type_id, name, first_last_name, second_last_name, email, age, phone, username, avatar, salary) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id
`

type CreateUserParams struct {
	RoleID         int    `json:"role_id"`
	UserTypeID     int    `json:"user_type_id"`
	Name           string `json:"name"`
	FirstLastName  string `json:"first_last_name"`
	SecondLastName string `json:"second_last_name"`
	Email          string `json:"email"`
	Age            int    `json:"age"`
	Phone          string `json:"phone"`
	Username       string `json:"username"`
	Avatar         string `json:"avatar"`
	Salary         int    `json:"salary"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.RoleID, arg.UserTypeID, arg.Name, arg.FirstLastName, arg.SecondLastName, arg.Email, arg.Age, arg.Phone, arg.Username, arg.Avatar, arg.Salary)
	var user User
	err := row.Scan(
		&user.ID,
	)
	return user.ID, err
}

const createPassword = `
INSERT INTO erp_schema.passwords (user_id, value)
VALUES ($1, $2)
`

type CreatePasswordParams struct {
	UserID   int32  `json:"user_id"`
	Password string `json:"password"`
}

func (q *Queries) CreatePassword(ctx context.Context, arg CreatePasswordParams) error {
	_, err := q.db.ExecContext(ctx, createPassword, arg.UserID, arg.Password)
	return err
}

const updateUser = `
UPDATE erp_schema.users u
SET name = $2, first_last_name = $3, second_last_name = $4, age = $5, avatar = $6, salary = $7
WHERE u.id = $1
`

type UpdateUserParams struct {
	ID             int32  `json:"id"`
	Name           string `json:"name"`
	FirstLastName  string `json:"first_last_name"`
	SecondLastName string `json:"second_last_name"`
	Age            int    `json:"age"`
	Avatar         string `json:"avatar"`
	Salary         int    `json:"salary"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser, arg.ID, arg.Name, arg.FirstLastName, arg.SecondLastName, arg.Age, arg.Avatar, arg.Salary)
	return err
}

const deleteUser = `
UPDATE erp_schema.users u
SET deleted = TRUE
WHERE u.id = $1
`

type DeleteUserParams struct {
	ID int32 `json:"id"`
}

func (q *Queries) DeleteUser(ctx context.Context, arg DeleteUserParams) error {
	_, err := q.db.ExecContext(ctx, deleteUser, arg.ID)
	return err
}

const getUsersForDownload = `
SELECT
    u.id,
    u.name,
    u.first_last_name,
    u.second_last_name,
    u.email,
    u.age,
    u.phone,
    u.username,
    u.avatar,
    u.cellphone_verification,
    u.salary,
    u.deleted,
    u.created_at,
    u.updated_at,
    r.id AS role_id,
    r.name AS role_name,
    ut.id AS user_type_id,
    ut.type AS user_type_name
FROM erp_schema.users u
LEFT JOIN erp_schema.roles r ON u.role_id = r.id
LEFT JOIN erp_schema.user_types ut ON u.user_type_id = ut.id
WHERE u.deleted IS NOT TRUE
ORDER BY u.id
`

func (q *Queries) GetUsersForDownload(ctx context.Context) ([]GetUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getUsersForDownload)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetUserRow{}
	for rows.Next() {
		var user GetUserRow
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.FirstLastName,
			&user.SecondLastName,
			&user.Email,
			&user.Age,
			&user.Phone,
			&user.Username,
			&user.Avatar,
			&user.CellphoneVerification,
			&user.Salary,
			&user.Deleted,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.RoleID,
			&user.RoleName,
			&user.UserTypeID,
			&user.UserTypeName,
		); err != nil {
			return nil, err
		}
		items = append(items, user)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
