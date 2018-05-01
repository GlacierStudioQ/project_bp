package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// 获取数据库连接
func DbGet() (db *sql.DB) {
	db, _ = sql.Open("mysql", "rache:21431@tcp(localhost:3306)/bipubipu?collation=utf8_general_ci&parseTime=true")
	// TODO 可能要处理返回值
	return
}

// 获取一个表中所有数据
func GetAll(entityTabName string) (result *sql.Rows) {
	result, _ = DbGet().Query("SELECT * FROM " + entityTabName)
	return result
}

// 根据一个字段的值查找数据库数据
func GetBy1Field(entityTabName string, tabfieldName string, value string) (result *sql.Rows) {
	result, _ = DbGet().Query("SELECT * FROM " + entityTabName + " WHERE " + tabfieldName + " = ?", value)
	return result
}