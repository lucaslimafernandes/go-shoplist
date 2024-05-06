package window

import (
	"go-shoplist/models"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ShowList(w fyne.Window) *widget.Form {

	var lists models.List_ofLists

	err := lists.ListAll()
	if err != nil {
		log.Println("err")
	}
	log.Println(lists)
	// var btns []*widget.Button
	var btns []fyne.CanvasObject
	for _, v := range lists {
		temp := widget.NewButton(v.Nome, func() { log.Println("clicou", v.Nome) })
		btns = append(btns, temp)
	}

	log.Println(btns)

	// b := widget.NewButton(lists[0].Nome, func() { log.Println("clicou", lists[0].Nome) })

	// scrollContainer := container.NewVScroll(
	// 	container.NewVBox(btns...),
	// )

	// form := &widget.Form{
	// 	Items: []*widget.FormItem{
	// 		{Text: "", Widget: scrollContainer},
	// 	},
	// }

	gridLayout := layout.NewGridLayoutWithColumns(1) // Define o layout de grade com uma coluna
	// gridLayout := layout.NewGridLayoutWithRows(10) // Define o layout de grade com uma coluna

	// Cria um contêiner de grade para os botões
	gridContainer := container.New(gridLayout,
		btns...,
	)

	// Cria um contêiner de rolagem vertical para a grade
	// scrollContainer := container.NewVScroll(
	// 	gridContainer,
	// )
	// Cria um formulário com o contêiner de rolagem
	// form := &widget.Form{
	// 	Content:  scrollContainer,
	// 	OnSubmit: func() {},
	// 	OnCancel: func() {},
	// }

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "", Widget: gridContainer}, //ou scrollContainer
		},
	}

	return form

}

// func ShowList(w fyne.Window) *fyne.Container {
// 	// Criar um contêiner para armazenar os botões das listas
// 	scrollContainer := container.NewVScroll()

// 	// Carregar as listas
// 	var lists models.List_ofLists
// 	err := lists.ListAll()
// 	if err != nil {
// 		log.Println("Erro ao carregar as listas:", err)
// 		return scrollContainer // Retorna o contêiner vazio em caso de erro
// 	}

// 	// Criar um botão para cada lista e adicionar ao contêiner
// 	for _, list := range lists {
// 		// Captura o valor de list para uso dentro da função de callback
// 		list := list
// 		b := widget.NewButton(list.Nome, func() { log.Println("Clicou", list.Nome) })
// 		scrollContainer.Add(b)
// 	}

// 	return scrollContainer
// }
