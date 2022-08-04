package page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/cast"
	"time"
)

func timePage(w fyne.Window) *fyne.Container {
	show := widget.NewMultiLineEntry()
	show.Wrapping = fyne.TextWrapBreak
	show.Disable()
	input := widget.NewMultiLineEntry()
	input.Wrapping = fyne.TextWrapBreak
	input.OnChanged = func(s string) {
		iTime := time.Unix(cast.ToInt64(s), 0)
		show.SetText("日期:" + iTime.Format("2006-01-02 15:04:05"))
	}

	input.SetText("")
	v := container.NewVSplit(
		input,
		show,
	)
	v.Offset = 0.5
	return container.NewMax(v)

}
