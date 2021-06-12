package repoCoffe

import (
	"context"

	coffeDomain "github.com/rayzalzero/go-sukha/src/domain/coffe"
	Error "github.com/rayzalzero/go-sukha/src/pkg/error"
)

func (m *repoCoffe) GetListCoffe(ctx context.Context, offset, limit int, search, sort string) (res []coffeDomain.ListCoffe, count int, err error) {
	query := `SELECT name, type FROM coffe`

	rows, err := m.Conn.QueryContext(ctx, query)
	if err != nil {
		Error.Error(err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		d := coffeDomain.ListCoffe{}
		err = rows.Scan(
			&d.Name,
			&d.Type,
		)
		if err != nil {
			Error.Error(err)
			return
		}
		res = append(res, d)
	}

	query = `SELECT COUNT(*) FROM coffe`
	err = m.Conn.QueryRow(query).Scan(&count)
	return
}
