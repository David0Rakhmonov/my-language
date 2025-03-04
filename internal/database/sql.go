package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectSQL подключается к SQL-базе данных
func ConnectSQL(driver, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}

// Query выполняет SQL-запрос
func Query(db *sql.DB, query string) (*sql.Rows, error) {
	return db.Query(query)
}
