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
	ID        int
	Lista     int
	Item      string
	Valor     float64
	Qtd       int
	Dt_insert string
	Status    int
}

type List_ofItems []Item

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

func (i *Item) ItemCreate() error {

	time_now := time.Now()
	fmt_time_now := time_now.Format("02/01/2006 15:04")

	r, err := Db.Exec("INSERT INTO item (lista, item, valor, qtd, dt_insert, status) VALUES (?, ?, ?, ?, ?, ?)",
		i.Lista, i.Item, i.Valor, i.Qtd, fmt_time_now, 0,
	)
	if err != nil {
		log.Println("Erro insert item:", err)
	}

	lastrowid, _ := r.LastInsertId()
	log.Println("inserted:", lastrowid)

	return err

}

func (li *List_ofItems) ListItems(l_id int) error {

	rows, err := Db.Query("SELECT id, lista, item, valor, qtd, dt_insert, status FROM item where lista = ?", l_id)
	if err != nil {
		log.Println("erro ao consultar itens da lista:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item Item

		err = rows.Scan(&item.ID, &item.Lista, &item.Item, &item.Valor, &item.Qtd, &item.Dt_insert, &item.Status)
		if err != nil {
			log.Println("err scan rows:", err)
		}

		*li = append(*li, item)
	}

	return rows.Err()

}
