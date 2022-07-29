package page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Page2(w fyne.Window) *fyne.Container {
	return container.NewVBox(
		widget.NewButton("Hi2!", func() {

		}),
	)
}
