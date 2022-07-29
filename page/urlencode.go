package page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"net/url"
)

func urlencode(w fyne.Window) *fyne.Container {
	end := widget.NewMultiLineEntry()
	end.Wrapping = fyne.TextWrapBreak
	entryLoremIpsum := widget.NewMultiLineEntry()
	entryLoremIpsum.Wrapping = fyne.TextWrapBreak
	entryLoremIpsum.OnChanged = func(s string) {
		enEscapeUrl, _ := url.QueryUnescape(s)
		end.SetText("解码:" + enEscapeUrl)
	}

	entryLoremIpsum.SetText("")
	v := container.NewVSplit(
		entryLoremIpsum,
		end,
	)
	v.Offset = 0.5
	return container.NewMax(v)

}
