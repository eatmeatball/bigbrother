package page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"time"
)

func clockPage(w fyne.Window) *fyne.Container {
	clock := widget.NewLabel("")
	formatted := time.Now().Format("Time: 15:04:05")
	clock.SetText(formatted)
	go func() {
		for range time.Tick(time.Second) {
			formatted := time.Now().Format("Time: 15:04:05")
			clock.SetText(formatted)
		}
	}()
	return container.NewVBox(clock)
}
