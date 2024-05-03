package database

import (
	"database/sql"
	"fmt"
	"log"
)

func DBSet() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ecommerce?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Susscessfully connected to MYSQL !!")
	return db
}

var DB *sql.DB = DBSet()
