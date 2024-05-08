package window

import (
	"fmt"
	"go-shoplist/models"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// TODO
// EDITAR ITEM
func EditItem(w fyne.Window, i int) {

	var item, new_item models.Item
	var bool_status string

	log.Println(i)
	item.GetItem(i)
	log.Println(item)

	iid := widget.NewEntry()
	lista := widget.NewEntry()
	str_item := widget.NewEntry()
	valor := widget.NewEntry()
	qtd := widget.NewEntry()
	dt_insert := widget.NewEntry()
	status := widget.NewEntry()

	iid.SetText(fmt.Sprintf("%v", item.ID))
	lista.SetText(fmt.Sprintf("%v", item.Lista))
	str_item.SetText(item.Item)
	valor.SetText(fmt.Sprintf("%v", item.Valor))
	qtd.SetText(fmt.Sprintf("%v", item.Qtd))
	dt_insert.SetText(item.Dt_insert)
	if item.Status == 0 {
		bool_status = "false"
	} else {
		bool_status = "true"
	}
	status.SetText(bool_status)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Item ID: ", Widget: &iid.DisableableWidget},
			{Text: "Lista ID: ", Widget: &lista.DisableableWidget},
			{Text: "Item: ", Widget: str_item},
			{Text: "Valor R$: ", Widget: valor},
			{Text: "Quantidade: ", Widget: qtd},
			{Text: "Data inserida: ", Widget: &dt_insert.DisableableWidget},
			{Text: "Status: ", Widget: status},
		},
		OnSubmit: func() {
			log.Println("Form submitted")
			intValue64, _ := strconv.ParseInt(iid.Text, 10, 64)
			new_item.ID = int(intValue64)

			intValue64, _ = strconv.ParseInt(lista.Text, 10, 64)
			new_item.Lista = int(intValue64)

			new_item.Item = str_item.Text
			new_item.Valor, _ = strconv.ParseFloat(valor.Text, 64)

			intValue64, _ = strconv.ParseInt(qtd.Text, 10, 64)
			new_item.Qtd = int(intValue64)

			new_item.Dt_insert = dt_insert.Text
			_st := 0
			if status.Text == "true" {
				_st = 1
			}
			new_item.Status = _st
			log.Println(new_item)
			// handlers.Create_handler(newListName.Text)
		},
	}

	w.SetContent(form)

}
