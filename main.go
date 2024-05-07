package main

import (
	"fmt"
	"go-shoplist/models"
	"go-shoplist/window"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {

	err := models.InitDB("db.sqlite3")
	if err != nil {
		log.Fatalln("initdb:", err)
	}

	a := app.NewWithID("go-shoplist")
	ismob := a.Driver().Device().IsMobile()
	log.Println("isMobile:", ismob)
	w := a.NewWindow("Form Widget")

	if ismob {
		// w.Resize(fyne.NewSize(screenSize.Width, screenSize.Height))
		w.SetFullScreen(true)
	} else {
		w.Resize(fyne.NewSize(500, 400))
		// w.SetFullScreen(true)
		// w.Resize(fyne.NewSize(screenSize.Width, screenSize.Height))
	}

	// lucas_time()
	l()

	// window.CreateList(w)
	window.Home(w)

	w.ShowAndRun()

}

func lucas_time() {
	a := time.Now()

	b := a.Format("02/01/2006 15:04") // Usando o layout correto

	fmt.Println(b)

	// 02: Representa o dia do mês com zero à esquerda, se necessário, para ter dois dígitos (por exemplo, "01", "02", ..., "31").
	// 01: Representa o mês com zero à esquerda, se necessário, para ter dois dígitos (por exemplo, "01" para janeiro, "02" para fevereiro, ..., "12" para dezembro).
	// 2006: Representa o ano com quatro dígitos (por exemplo, "2006", "1999", "2023", etc.).
}

func l() {

	var item models.Item
	var listas models.List_ofLists
	var listas_nomes []string
	listas.ListAll()

	for _, n := range listas {
		fmt.Println("add", n.Nome)
		listas_nomes = append(listas_nomes, n.Nome)
	}

	log.Println(listas)
	log.Println(listas_nomes)

	item.Item = "nome"
	log.Println(item)

}
