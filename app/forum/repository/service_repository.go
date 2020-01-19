package repository

import (
	"github.com/moguchev/BD-Forum/pkg/sql_queries"
)

func (r *Repository) Clear() error {
	_, err := r.DbConn.Exec(sql_queries.Truncate)

	if err != nil {
		return err
	}

	return nil
}
