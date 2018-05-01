package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DbGet() (db *sql.DB) {
	db, _ = sql.Open("mysql", "rache:21431@tcp(localhost:3306)/bipubipu?collation=utf8_general_ci&parseTime=true")
	// TODO 可能要处理返回值
	return
}
