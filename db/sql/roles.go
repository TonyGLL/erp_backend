package db

import (
	"context"
	"encoding/json"
)

const getRoles = `
SELECT
    JSONB_BUILD_OBJECT(
        'id', r.id,
        'name', r.name,
        'modules', JSONB_AGG(
            JSONB_BUILD_OBJECT(
                'id', m.id,
                'name', m.name
            )
        )
    ) AS role_data
FROM
    erp_schema.roles r
JOIN
    erp_schema.role_modules rm ON r.id = rm.role_id
JOIN
    erp_schema.modules m ON rm.module_id = m.id
GROUP BY
    r.id, r.name
LIMIT $1 OFFSET $2
`

type GetRolesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetRoles(ctx context.Context, arg GetRolesParams) ([]Role, error) {
	rows, err := q.db.QueryContext(ctx, getRoles, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roles := []Role{}
	for rows.Next() {
		var roleData []byte // JSONB data as byte slice
		if err := rows.Scan(&roleData); err != nil {
			return nil, err
		}

		var role Role
		if err := json.Unmarshal(roleData, &role); err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

const CountRoles = `
SELECT COUNT(*) AS total_roles
FROM erp_schema.roles r;
`

func (q *Queries) CountRoles(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, CountRoles)
	var total_roles int64
	err := row.Scan(&total_roles)
	return total_roles, err
}
