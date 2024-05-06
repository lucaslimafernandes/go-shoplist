package window

import (
	"log"

	"go-shoplist/handlers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func CreateList(w fyne.Window) *widget.Form {

	newListName := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Nova lista: ", Widget: newListName},
		},
		OnSubmit: func() {
			log.Println("Form submitted")
			handlers.Create_handler(newListName.Text)
		},
	}

	return form
	// w.SetContent(form)

}

// form := &widget.Form{
// 	Items: []*widget.FormItem{
// 		{Text: "uname", Widget: uname},
// 		{Text: "password", Widget: paswd},
// 		{Text: "email", Widget: em},
// 	},
// 	OnSubmit: func() {
// 		log.Println("New user form submitted")
// 		status := requests.NewUser(uname.Text, paswd.Text, em.Text)
// 		if !status {
// 			log.Fatalln("NEW USER NOT OK")
// 		}
// 		LoginPage(w, a)
// 	},
// }
