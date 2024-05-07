package window

import (
	"fmt"
	"go-shoplist/models"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func PageItems(w fyne.Window, l models.List) *widget.Table {

	// var btns []fyne.CanvasObject

	var items models.List_ofItems
	var data [][]string
	items.ListItems(l.ID)

	// Headr
	data = append(data, []string{"id", "id_lista", "Item", "Valor R$", "Quantidade", "Total R$", "dt_inserted", "status"})

	for _, v := range items {
		data = append(data, []string{
			fmt.Sprintf("%v", v.ID),
			fmt.Sprintf("%v", v.Lista),
			v.Item,
			fmt.Sprintf("%v", v.Valor),
			fmt.Sprintf("%v", v.Qtd),
			fmt.Sprintf("%.2f", float64(v.Qtd)*v.Valor),
			v.Dt_insert,
			fmt.Sprintf("%v", v.Status),
		})
	}

	log.Println(items)
	// TODO for para preencher btns -> show list
	// Refazer essa tela
	// Fazer a seleção de itens
	// query, etc
	// form := &widget.Form{
	// 	Items: []*widget.FormItem{
	// 		{Text: "Nome", Widget: widget.NewEntry()},
	// 		{Text: "Descrição", Widget: widget.NewEntry()},
	// 		{Text: "Quantidade", Widget: widget.NewEntry()},
	// 	},
	// 	OnSubmit: func() {},
	// 	OnCancel: func() {},
	// }
	list := widget.NewTable(
		func() (int, int) {
			return len(data), 8
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			// o.(*widget.Label).SetText(items[i.Row][i.Col])
			o.(*widget.Label).SetText(data[i.Row][i.Col])
			// label := o.(*widget.Label)
			// label.SetText(data[i.Row][i.Col])
			if i.Row == 0 {
				o.(*widget.Label).Alignment = fyne.TextAlignCenter
				o.(*widget.Label).TextStyle.Bold = true
				// label.TextStyle.Bold = true
			} else {
				o.(*widget.Label).Alignment = fyne.TextAlignLeading
				o.(*widget.Label).TextStyle.Bold = false
			}
		},
	)

	list.OnSelected = func(id widget.TableCellID) {
		if id.Row > 0 { // Ignore header row
			fmt.Printf("Selected: %s\n", data[id.Row][id.Col])
			// Implement your edit logic here
		}
	}

	// list.OnUnselected = func(id widget.TableCellID) {
	// 	if id.Row > 0 { // Ignore header row
	// 		fmt.Printf("Unselected: %s\n", data[id.Row][id.Col])
	// 	}
	// }

	// list.CreateCell = func() fyne.CanvasObject {
	// 	button := widget.NewButton("Edit", func() {
	// 		// Implement your edit logic here
	// 	})
	// 	return container.NewHBox(button)
	// }

	// // Layout de grade com 2 linhas e 1 coluna
	// gridLayout := layout.NewGridLayoutWithRows(2)

	// // // Adicionar o formulário e a lista de botões ao contêiner de grade
	// gridContainer := container.New(gridLayout,
	// 	// form,
	// 	container.NewVScroll(container.NewVBox(btns...)), // Lista de botões dentro de um contêiner de rolagem vertical
	// )
	// log.Println(gridContainer)
	return list

}
