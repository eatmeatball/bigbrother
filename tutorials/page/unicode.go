package page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/leancodebox/goose/unicode"
)

func UnicodePage(w fyne.Window) fyne.CanvasObject {
	show := widget.NewMultiLineEntry()
	show.Wrapping = fyne.TextWrapBreak
	//show.Disable()
	input := widget.NewMultiLineEntry()
	input.Wrapping = fyne.TextWrapBreak
	input.OnChanged = func(s string) {
		enEscapeUrl := unicode.Decode(s)
		show.SetText("解码:" + enEscapeUrl)
	}

	input.SetText("\\u4e07\\u82b1\\u7b52")
	v := container.NewVSplit(
		input,
		show,
	)
	v.Offset = 0.5
	return container.NewStack(v)

}
