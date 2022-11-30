package page

import (
	"encoding/base64"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/cast"
)

func Base64Page(w fyne.Window) *fyne.Container {

	show := widget.NewMultiLineEntry()
	show.Wrapping = fyne.TextWrapBreak
	input := widget.NewMultiLineEntry()
	input.Wrapping = fyne.TextWrapBreak
	input.OnChanged = func(s string) {
		data, err := base64.StdEncoding.DecodeString(s)
		var result string
		if err != nil {
			result = "无法解析"
		}
		result = cast.ToString(data)
		show.SetText(result)
	}

	input.SetText(`eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoidGhoIiwiZXhwIjoxNjYwNzM0OTMwLCJuYmYiOjE2NjAxMzAxMzB9.tnTQYaaM7FfJhah9veVNS0OSkh1q4hKtP1UUUgESCQY`)
	v := container.NewVSplit(
		input,
		show,
	)
	v.Offset = 0.5
	return container.NewMax(v)

}
