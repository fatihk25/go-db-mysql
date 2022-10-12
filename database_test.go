package gomysql

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:mint22**@tcp(localhost:3306)/go_study_db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

}
