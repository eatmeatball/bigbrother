package page

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"time"
)

func ClockPage(_ fyne.Window) fyne.CanvasObject {
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

	head := widget.NewLabelWithStyle("Program Summary", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	c := container.NewVBox(
		widget.NewButton("time", func() {

		}),
		widget.NewButton("time", func() {

		}),
		widget.NewButton("time", func() {

		}),
	)

	card := widget.NewCard("", "", c)

	rectangles := canvas.NewRectangle(color.RGBA{R: 49, G: 48, B: 66})
	tmp := container.NewStack(
		container.NewVBox(head, card),
		rectangles,
	)

	return container.NewVBox(clock, tmp)
}
