package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DBSet() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/threadsapplication?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Susscessfully connected to MYSQL !!")
	fmt.Printf("Connected to")
	return db
}

var Client *sql.DB = DBSet()
