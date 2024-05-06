package window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Home(w fyne.Window) {

	create := CreateList(w)
	show := ShowList(w)

	tabs := container.NewAppTabs(
		container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
		container.NewTabItem("Nova lista", create),
		container.NewTabItem("Exibir lista", show),
		container.NewTabItem("Tab 3", widget.NewLabel("World!")),
	)

	tabs.SetTabLocation(container.TabLocationBottom)
	w.SetContent(tabs)

}
