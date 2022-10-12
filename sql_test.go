package gomysql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSQL(t *testing.T) {
	db := GetConnections()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer (ID, NAME) VALUES('joko', 'JOKO')"

	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySQL(t *testing.T) {
	db := GetConnections()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer"

	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
	}

	defer rows.Close()
}

func TestQueryDataTypeSQL(t *testing.T) {
	db := GetConnections()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"

	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var createdAt time.Time
		var birthDate sql.NullTime
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("============")
		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
		if email.Valid {
			fmt.Println("EMail: ", email.String)
		}
		fmt.Println("Balance: ", balance)
		fmt.Println("Rating: ", rating)
		if birthDate.Valid {
			fmt.Println("Birth date: ", birthDate.Time)
		}
		fmt.Println("Married: ", married)
		fmt.Println("Created at: ", createdAt)

	}

	defer rows.Close()
}

func TestSQLInjection(t *testing.T) {
	db := GetConnections()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "admin"

	script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"

	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success login", username)
	} else {
		fmt.Println("Gagal login")
	}

	defer rows.Close()
}

func TestSQLInjectionSafe(t *testing.T) {
	db := GetConnections()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"

	rows, err := db.QueryContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success login", username)
	} else {
		fmt.Println("Gagal login")
	}

	defer rows.Close()
}

func TestExecSQLParam(t *testing.T) {
	db := GetConnections()
	defer db.Close()

	ctx := context.Background()

	username := "johan"
	password := "johan"

	script := "INSERT INTO user (username, password) VALUES(?, ?)"

	_, err := db.ExecContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnections()
	defer db.Close()

	ctx := context.Background()

	email := "johan@mail.com"
	comment := "Test comment 2"

	script := "INSERT INTO comments (email, comment) VALUES(?, ?)"

	result, err := db.ExecContext(ctx, script, email, comment)

	if err != nil {
		panic(err)
	}

	insertID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comments with id", insertID)
}

func TestPrepareStament(t *testing.T) {
	db := GetConnections()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comments (email, comment) VALUES(?, ?)"
	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "johan " + strconv.Itoa(i) + " @gmail.com"
		comment := "ini comment ke " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)

		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("COmment id ", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnections()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments (email, comment) VALUES(?, ?)"

	//do trans
	for i := 0; i < 10; i++ {
		email := "johan " + strconv.Itoa(i) + " @gmail.com"
		comment := "ini comment ke " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)

		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("COmment id ", id)
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}

}
