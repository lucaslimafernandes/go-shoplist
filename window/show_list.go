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
		temp := widget.NewButton(v.Nome, func() {
			log.Println("clicou", v.Nome)
			res, ed_call := PageItems(w, v)
			btn_home := widget.NewButton("Voltar", func() { Home(w) })
			// btn_edit := widget.NewButton("Editar", func() {
			// 	res.OnSelected = func(id widget.TableCellID) {
			// 		log.Printf("Selected: %v\n", id.Row)
			// 		if id.Row > 0 { // Ignore header row
			// 			// fmt.Printf("Selected: %s\n", data[id.Row][id.Col])
			// 			fmt.Printf("Selected: %v\n", id.Row)
			// 			// Implement your edit logic here
			// 		}
			// 	}
			// })
			btn_edit := widget.NewButton("Editar", func() { ed_call() })
			// Cria um contêiner de grade para os botões
			// gridLayout := layout.NewGridLayoutWithRows(1) // Define o layout de grade com uma coluna
			gridLayout := layout.NewGridLayoutWithColumns(1) // Define o layout de grade com uma coluna

			gridContainer := container.New(gridLayout,
				res,
				btn_home,
				btn_edit,
			)
			// w.SetContent(res)
			w.SetContent(gridContainer)
		})
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
