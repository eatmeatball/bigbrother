package page

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/cast"
	"gopkg.in/yaml.v3"
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

func DataBindPage(w fyne.Window) fyne.CanvasObject {

	oData := ""
	bindData := binding.BindString(&oData)

	input := widget.NewMultiLineEntry()
	input.Bind(bindData)
	input.Wrapping = fyne.TextWrapBreak

	show := widget.NewMultiLineEntry()
	show.Bind(bindData)
	show.Wrapping = fyne.TextWrapBreak

	//show.Disable()

	input.SetPlaceHolder("输入json")

	rows := container.NewGridWithRows(1,
		container.NewGridWithColumns(2, input, show),
	)

	return container.NewStack(rows)
}
