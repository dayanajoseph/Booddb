package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	dsn := "root:Djm@1211@tcp(localhost:3306)/booksDB"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}
