package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// insertToDatabase()
	// queryDatabase()
	// SafeFromSqlInjection()
	// AutoIncrementQuery()
	// PrepareStmt()
	TransactionalDbOperation()
}

func TransactionalDbOperation() {
	db := getConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO comments (email, comment) VALUES (?,?)"
	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	for i := 10; i < 20; i++ {
		email := fmt.Sprintf("emailbaru%v@gmail.com", i)
		comment := fmt.Sprintf("komen ke-%v", i)

		_, err := tx.ExecContext(ctx, query, email, comment)

		if err != nil {
			panic(err)
		}
	}

	// you can either commit or rollback
	// tx.Commit()
	err = tx.Rollback()
	if err != nil {
		panic(err)
	}

	fmt.Println("BERHASIL")
}

// multiple insert
func PrepareStmt() {
	db := getConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	stmt, err := db.PrepareContext(ctx, query)

	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := fmt.Sprintf("reno%v@gmail.com", i)
		comment := fmt.Sprintf("ini komentar ke-%v", i)

		_, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

	}

	fmt.Sprintln("SUKSES")
}

func AutoIncrementQuery() {
	db := getConnection()
	defer db.Close()

	ctx := context.Background()

	email := "reno@gmail.com"
	comment := "tes komentar"

	query := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	result, err := db.ExecContext(ctx, query, email, comment)

	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Printf("sukses menambahkan data dengan id = %v\n", id)

}

func SafeFromSqlInjection() {
	db := getConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	query := "SELECT username FROM user WHERE username = ? AND password = ?"
	rows, err := db.QueryContext(ctx, query, username, password)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		fmt.Println("berhasil")
	} else {
		fmt.Println("gagal")
	}

}

func queryDatabase() {
	db := getConnection()
	defer db.Close()

	context := context.Background()

	query := "SELECT * FROM customer"
	rows, err := db.QueryContext(context, query)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// print data
	// while rows has next
	for rows.Next() {
		var id, name string

		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println(id)
		fmt.Println(name)
	}

}

func insertToDatabase() {
	// id := "budi"
	// name := "BUDI"
	// query := fmt.Sprintf("INSERT INTO customer VALUES('%v', '%v')", id, name)
	// fmt.Println(query)

	db := getConnection()
	defer db.Close()

	// context
	context := context.Background()

	// exec context can be use either INSERT, UPDATE, or DELETE
	query := "INSERT INTO customer VALUES ('budi', 'budi')"
	_, err := db.ExecContext(context, query)

	if err != nil {
		panic(err)
	}

	fmt.Println("Berhasil tambah data")
}

func getConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang_db")

	if err != nil {
		fmt.Println(err)
		panic("Connection Failed")
	}

	// database pooling
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(time.Second * 5)
	db.SetConnMaxLifetime(time.Hour * 1)

	// defer db.Close()

	return db
}
