package utils

import (
	"database/sql"
)

func DBConn() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/student_dosen_pa")
	PanicError(err)
	return db
}