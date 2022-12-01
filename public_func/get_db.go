package public_func

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Getdb() (*sql.DB, error) {
	var dns = "root:114514@tcp(127.0.0.1:3306)/test"
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}
	return db, nil
}
