package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:Mahik87@@tcp(127.0.0.1:3306)/godb")

	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println("error validating sql.open arguments")
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connectin with db.Ping")
		panic(err.Error())
	}

	insert, err := db.Query("INSERT INTO godb.test (id,firstname,lastname) VALUES ('10','Gani','P');")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Succesful connection to db")

}
