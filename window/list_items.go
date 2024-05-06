package window

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func PageItems(w fyne.Window) {

	var btns []fyne.CanvasObject
	// TODO for para preencher btns -> show list
	// Refazer essa tela
	// Fazer a seleção de itens
	// query, etc
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Nome", Widget: widget.NewEntry()},
			{Text: "Descrição", Widget: widget.NewEntry()},
			{Text: "Quantidade", Widget: widget.NewEntry()},
		},
		OnSubmit: func() {},
		OnCancel: func() {},
	}

	// Layout de grade com 2 linhas e 1 coluna
	gridLayout := layout.NewGridLayoutWithRows(2)

	// Adicionar o formulário e a lista de botões ao contêiner de grade
	gridContainer := container.New(gridLayout,
		form,
		container.NewVScroll(container.NewVBox(btns...)), // Lista de botões dentro de um contêiner de rolagem vertical
	)
	log.Println(gridContainer)
	// return gridContainer
}
