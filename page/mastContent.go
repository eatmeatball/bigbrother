package page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func masterContent(w fyne.Window) *fyne.Container {

	md := widget.NewRichTextFromMarkdown(`
# RichText Heading

## A Sub Heading

---

* Item1 in _three_ segments
* Item2
* Item3

Normal **Bold** *Italic* [Link](https://fyne.io/) and some ` + "`Code`" + `.
This styled row should also wrap as expected, but only *when required*.

> An interesting quote here, most likely sharing some very interesting wisdom.`)

	md.Wrapping = fyne.TextWrapBreak

	entryLoremIpsum := widget.NewMultiLineEntry()
	entryLoremIpsum.SetText("loremIpsum")
	entryLoremIpsum.OnChanged = func(s string) {
		md.ParseMarkdown(s)
	}
	return container.NewVBox(entryLoremIpsum,
		entryLoremIpsum,
		md)
}
