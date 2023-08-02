package page

import (
	"encoding/base64"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/cast"
)

func Base64Page(w fyne.Window) fyne.CanvasObject {

	show := widget.NewMultiLineEntry()
	show.Wrapping = fyne.TextWrapBreak

	show.SetText(`eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoidGhoIiwiZXhwIjoxNjYwNzM0OTMwLCJuYmYiOjE2NjAxMzAxMzB9.tnTQYaaM7FfJhah9veVNS0OSkh1q4hKtP1UUUgESCQY`)

	base64DecodeButton := widget.NewButton("decode", func() {
		data, _ := base64.StdEncoding.DecodeString(show.Text)
		var result string
		//fmt.Println(err, data)
		//if err != nil {
		//	return
		//}
		result = cast.ToString(data)
		if len(data) == 0 {
			return
		}
		show.SetText(result)
	})
	base64EncodeButton := widget.NewButton("encode", func() {
		data := base64.StdEncoding.EncodeToString([]byte(show.Text))
		show.SetText(data)
	})
	masterContent := container.NewBorder(
		container.NewGridWithRows(1,
			container.NewGridWithColumns(2, base64DecodeButton, base64EncodeButton),
		), nil, nil, nil,
		show,
	)
	return masterContent

}
