package page

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"time"
)

func getTimeButton(w fyne.Window) *fyne.Container {

	clock := widget.NewLabel("")
	clock.SetText("")
	clock.Alignment = fyne.TextAlignCenter
	go func() {
		for range time.Tick(time.Second) {
			now := time.Now()
			text := fmt.Sprintf("Time: %v\nUnix:%v",
				now.Format("2006-01-02 15:04:05"),
				now.Unix(),
			)
			clock.SetText(text)
		}
	}()
	path := widget.NewLabel("")
	return container.NewVBox(
		clock,
		path,
		widget.NewButton("OpenFile", func() {
			fd := dialog.NewFileOpen(func(closer fyne.URIReadCloser, err error) {
				path.SetText(closer.URI().String())
			}, w)

			fd.Show()
		}),
		widget.NewButton("OpenDir", func() {

			fd := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
				path.SetText(uri.String())
			}, w)
			fd.Show()
		}),
	)
}
