package handlers

import (
	"go-shoplist/models"
	"log"
)

func Create_handler(name string) {
	// var a models.List
	a := models.List{Nome: name}
	a.Create()
	log.Println(a)
}
