package page

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/cast"
	"gopkg.in/yaml.v3"
	"strings"
)

func Json2yaml(w fyne.Window) fyne.CanvasObject {

	input := widget.NewMultiLineEntry()
	show := widget.NewMultiLineEntry()

	input.Wrapping = fyne.TextWrapBreak
	show.Wrapping = fyne.TextWrapBreak

	//show.Disable()

	input.SetPlaceHolder("输入json")
	input.OnChanged = func(s string) {
		var obj any
		_ = json.Unmarshal([]byte(s), &obj)
		aa, _ := yaml.Marshal(obj)
		show.SetText(cast.ToString(aa))
	}

	input.SetText(`{
	"object":{
		"name":"xxx",
		"sex":1,
		"success":true
	},
	"list":[
		1,2,3,4
	],
	"listObject":[
		{"name":"name1","age":12},
		{"name":"name1","age":12},
		{"name":"name1","age":12},
		{"name":"name1","age":12}
	]
}`)

	rows := container.NewGridWithRows(1,
		container.NewGridWithColumns(2, input, show),
	)

	return container.NewStack(rows)

}

func LowerUpper(w fyne.Window) fyne.CanvasObject {

	inputBig := binding.NewString()
	inputLower := binding.NewString()

	lower := widget.NewMultiLineEntry()
	lower.Bind(inputBig)
	lower.OnChanged = func(s string) {
		_ = inputLower.Set(strings.ToUpper(s))
	}
	lower.Wrapping = fyne.TextWrapBreak

	upper := widget.NewMultiLineEntry()
	upper.Bind(inputLower)
	upper.OnChanged = func(s string) {
		_ = inputBig.Set(strings.ToLower(s))
	}
	upper.Wrapping = fyne.TextWrapBreak

	//show.Disable()

	lower.SetPlaceHolder("消息字母")

	rows := container.NewGridWithRows(1,
		container.NewGridWithColumns(2, lower, upper),
	)

	return container.NewStack(rows)
}
