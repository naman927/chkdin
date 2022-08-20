package modals

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func SetUpDb() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/demo?parseTime=true")
	if err != nil {
		return nil, err
	}
	return db, nil
}
