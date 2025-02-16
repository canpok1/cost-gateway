package db

import (
	"database/sql"
	"fmt"
)

func Open(host string, port int, user, password, database string) (*sql.DB, error) {
	return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, password, host, port, database))
}
