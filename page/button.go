package page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func getTimeButton(w fyne.Window) *fyne.Container {

	text := widget.NewLabel("xxxx")
	return container.NewVBox(
		text,
		widget.NewButton("GetNow", func() {
			text.SetText("now")
		}),
		widget.NewButton("File Open With Filter (.jpg or .png)", func() {
			fd := dialog.NewFileOpen(func(closer fyne.URIReadCloser, err error) {
			}, w)
			fd.Show()
		}),
	)
}
