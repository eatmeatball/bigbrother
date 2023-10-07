package page

import (
	"bigbrother/arms/str"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Camel2snake(w fyne.Window) fyne.CanvasObject {

	input := widget.NewMultiLineEntry()
	show := widget.NewMultiLineEntry()

	input.Wrapping = fyne.TextWrapBreak
	show.Wrapping = fyne.TextWrapBreak

	//show.Disable()

	input.SetPlaceHolder("输入json")
	input.OnChanged = func(s string) {
		show.SetText(str.Snake(s))
	}

	input.SetText(`abandonType
okType
TypeList
`)

	rows := container.NewGridWithRows(1,
		container.NewGridWithColumns(2, input, show),
	)

	return container.NewStack(rows)

}
