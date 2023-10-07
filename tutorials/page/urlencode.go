package page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"net/url"
)

func UrlEncodePage(w fyne.Window) fyne.CanvasObject {
	show := widget.NewMultiLineEntry()
	show.Wrapping = fyne.TextWrapBreak
	//show.Disable()
	input := widget.NewMultiLineEntry()
	input.Wrapping = fyne.TextWrapBreak
	input.OnChanged = func(s string) {
		enEscapeUrl, _ := url.QueryUnescape(s)
		show.SetText("解码:" + enEscapeUrl)
	}

	input.SetText("%E6%B1%89%E5%AD%97")
	v := container.NewVSplit(
		input,
		show,
	)
	v.Offset = 0.5
	return container.NewStack(v)

}
