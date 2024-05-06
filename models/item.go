package models

import (
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type List struct {
	ID         int
	Nome       string
	Completa   bool
	Dt_created string
}

type List_ofLists []List

type Item struct {
	id        int
	lista     int
	item      string
	valor     float64
	qtd       int
	dt_insert string
	status    int
}

type List_ofItems []Item

// var Db *sql.DB

// func oi() {
// 	fmt.Println("oi")
// }

// func InitDB(databasePath string) error {

// 	_, err := os.Stat(databasePath)
// 	if os.IsNotExist(err) {
// 		log.Println("NOOT EXISTS")
// 	}
// 	// var err error
// 	Db, err = sql.Open("sqlite3", databasePath)
// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }

func (l *List) Create() (int, error) {

	time_now := time.Now()
	fmt_time_now := time_now.Format("02/01/2006 15:04")

	r, err := Db.Exec("INSERT INTO lista (nome, completa, dt_created) VALUES (?, 0, ?);", l.Nome, fmt_time_now)
	if err != nil {
		log.Println("List Insert Err:", err)
		return 0, err
	}

	lastrowid, _ := r.LastInsertId()
	return int(lastrowid), err

}

func (l *List_ofLists) ListAll() error {

	rows, err := Db.Query("SELECT id, nome, completa, dt_created from lista where completa = 0;")
	if err != nil {
		log.Println("Err no select:", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var il List

		err = rows.Scan(&il.ID, &il.Nome, &il.Completa, &il.Dt_created)
		if err != nil {
			return err
		}

		*l = append(*l, il)

	}
	return rows.Err()

}
