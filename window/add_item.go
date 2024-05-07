package window

import (
	"go-shoplist/models"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func AddItem(w fyne.Window) *widget.Form {

	var item models.Item
	var listas models.List_ofLists
	var listas_nomes []string
	lni := make(map[string]int)

	listas.ListAll()
	for _, v := range listas {
		listas_nomes = append(listas_nomes, v.Nome)
		lni[v.Nome] = v.ID
	}

	// var selected string
	// listas.ListAll()

	wl := widget.NewSelect(listas_nomes, func(s string) {})
	item_item := widget.NewEntry()
	item_valor := widget.NewEntry()
	item_quantidade := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Qual lista: ", Widget: wl},
			{Text: "Item: ", Widget: item_item},
			{Text: "Valor: ", Widget: item_valor},
			{Text: "Quantidade: ", Widget: item_quantidade},
		},
		OnSubmit: func() {
			log.Println("Form submitted")
			log.Println(wl.Selected)
			item.Lista = lni[wl.Selected]
			item.Item = item_item.Text
			qtd, err := strconv.Atoi(item_quantidade.Text)
			if err != nil {
				log.Println("Digito errado qtd", err)
			}
			item.Qtd = qtd
			val, err := strconv.ParseFloat(item_valor.Text, 64)
			if err != nil {
				log.Println("Digito errado vl", err)
			}
			item.Valor = val
			// handlers.Create_handler(newListName.Text)
			err = item.ItemCreate()
			if err != nil {
				log.Println("form insert item:", err)
			}

			log.Println(item)
		},
	}

	return form

}
