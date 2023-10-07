package page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/cast"
	"time"
)

func TimePage(w fyne.Window) fyne.CanvasObject {

	t1 := tractionInput("时间戳转化日期", func(s string) string {
		return time.Unix(cast.ToInt64(s), 0).Format("2006-01-02 15:04:05")
	})

	t2 := tractionInput("日期转化时间戳", func(s string) string {
		var LOC, _ = time.LoadLocation("Asia/Shanghai")
		timeTemplate := "2006-01-02 15:04:05"
		tim, _ := time.ParseInLocation(timeTemplate, s, LOC)
		return cast.ToString(tim.Unix())
	})
	page := container.NewVBox(t1, t2,
		widget.NewSeparator(), getNowTime())
	return container.NewStack(page)

}

func tractionInput(desc string, traction func(s string) string) *fyne.Container {
	// init
	input := widget.NewEntry()
	descShow := widget.NewLabel(desc)
	show := widget.NewEntry()
	// set style
	descShow.Alignment = fyne.TextAlignCenter
	//show.Disable()
	// set event
	input.OnChanged = func(s string) {
		show.SetText(traction(s))
	}

	rows := container.NewGridWithRows(1,
		container.NewGridWithColumns(3, input, descShow, show),
	)
	return rows
}

func getNowTime() *fyne.Container {
	show1 := widget.NewEntry()
	show2 := widget.NewEntry()
	button := widget.NewButton("获取最新时间", func() {
		show1.SetText(time.Now().Format("2006-01-02 15:04:05"))
		show2.SetText(cast.ToString(time.Now().Unix()))
	})
	rows := container.NewGridWithRows(1,
		container.NewGridWithColumns(3, show1, show2, button),
	)
	return rows
}
