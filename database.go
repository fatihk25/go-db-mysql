package gomysql

import (
	"database/sql"
	"time"
)

func GetConnections() *sql.DB {
	db, err := sql.Open("mysql", "root:mint22**@tcp(localhost:3306)/go_study_db?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
