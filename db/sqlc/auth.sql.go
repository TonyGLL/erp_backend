package db

import "context"

const getUserPassword = `
SELECT p.value AS password FROM erp_schema.passwords p
INNER JOIN erp_schema.users u ON p.user_id = u.id
WHERE u.username = $1 AND u.deleted IS NOT TRUE LIMIT 1
`

func (q *Queries) GetUserPassword(ctx context.Context, username string) (string, error) {
	row := q.db.QueryRowContext(ctx, getUserPassword, username)
	var password string
	err := row.Scan(&password)
	return password, err
}
