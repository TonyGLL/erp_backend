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
    ut.type AS user_type_name,
	ad.*
FROM users u
INNER JOIN roles r ON u.role_id = r.id
INNER JOIN user_types ut ON u.user_type_id = ut.id
INNER JOIN addresses ad ON u.id = ad.user_id
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
	Addresses             []Address `json:"addresses"`
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
		&user.Addresses,
	)
	return user, err
}

const getUserPassword = `
SELECT p.value AS password FROM passwords p
INNER JOIN users u ON p.user_id = u.id
WHERE u.username = $1 LIMIT 1
`

func (q *Queries) GetUserPassword(ctx context.Context, username string) (string, error) {
	row := q.db.QueryRowContext(ctx, getUserPassword, username)
	var password string
	err := row.Scan(&password)
	return password, err
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
FROM users u
JOIN roles r ON u.role_id = r.id
JOIN user_types ut ON u.user_type_id = ut.id
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
FROM users;
`

func (q *Queries) CountUsers(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, CountUsers)
	var total_accounts int64
	err := row.Scan(&total_accounts)
	return total_accounts, err
}
