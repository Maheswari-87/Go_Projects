package main

import (
	"fmt"
	"log"
	"net/http"
	"workspace/dbProject/controller"
	"workspace/dbProject/model"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("serving.......")
	log.Fatal(http.ListenAndServe(":3007", mux))

}
