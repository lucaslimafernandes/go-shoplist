package models

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

// func oi() {
// 	fmt.Println("oi")
// }

func InitDB(databasePath string) error {

	_, err := os.Stat(databasePath)
	if os.IsNotExist(err) {
		log.Println("NOOT EXISTS")
		create_db(databasePath)
	}
	// var err error
	Db, err = sql.Open("sqlite3", databasePath)
	if err != nil {
		return err
	}

	return nil

}

func create_db(databasePath string) {

	file, err := os.Create(databasePath)
	if err != nil {
		log.Fatalln("Error to create db file:", err)
	}
	defer file.Close()

	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		log.Fatalln("Erro to open db file:", err)
	}
	defer db.Close()

	l, err := db.Exec(`CREATE TABLE lista (
		id integer PRIMARY KEY AUTOINCREMENT UNIQUE,
		nome text,
		completa int,
		dt_created text   
	);`)

	log.Println(l, err)

	db.Exec(`CREATE TABLE item (
		id integer PRIMARY KEY AUTOINCREMENT UNIQUE,
		lista int,
		item text,
		valor real,
		qtd int,
		dt_insert text,
		status int
	);`)

}
