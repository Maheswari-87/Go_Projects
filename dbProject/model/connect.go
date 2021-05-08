package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:Mahik87@@tcp(localhost:3306)/mysql")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection established")
	con = db
	return db
}
