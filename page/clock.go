package page

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"time"
)

func clockPage(w fyne.Window) *fyne.Container {
	clock := widget.NewLabel("")
	clock.SetText("")
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
	return container.NewVBox(clock)
}
