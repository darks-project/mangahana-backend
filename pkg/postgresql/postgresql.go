package postgresql

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func Connect() (*sqlx.DB, error) {
	conn, err := sqlx.Connect("pgx", "postgres://pgx_md5:secret@localhost:5432/pgx_test")
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return nil, nil
}
