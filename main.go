package main

import (
	"log"

	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.NewWithID("roc-app-1")
	ismob := a.Driver().Device().IsMobile()
	log.Println("isMobile:", ismob)
	// w := a.NewWindow("Form Widget")

}
