package page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"time"
)

func getTimeButton(w fyne.Window) *fyne.Container {
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text := canvas.NewText("hi", green)
	text.Alignment = fyne.TextAlignTrailing
	text.TextStyle = fyne.TextStyle{Italic: true}
	return container.NewVBox(
		text,
		widget.NewButton("GetNow", func() {
			text.Text = time.Now().Format("2006-01-02 15:04:05")
		}),
		widget.NewButton("File Open With Filter (.jpg or .png)", func() {
			fd := dialog.NewFileOpen(func(closer fyne.URIReadCloser, err error) {
			}, w)
			fd.Show()
		}),
	)
}
