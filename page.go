package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"time"
)

var (
	tList = []ListNode{
		{Title: "page1", Intro: "page1", View: masterContent},
		{Title: "page1", Intro: "page1", View: Page1},
		{Title: "page2", Intro: "page2", View: Page2},
	}
)

type ListNode struct {
	Title, Intro string
	View         func(w fyne.Window) *fyne.Container
}

func Page1(w fyne.Window) *fyne.Container {

	end := widget.NewMultiLineEntry()
	end.Wrapping = fyne.TextWrapBreak
	
	entryLoremIpsum := widget.NewMultiLineEntry()
	entryLoremIpsum.Wrapping = fyne.TextWrapBreak
	entryLoremIpsum.SetText(loremIpsum)
	entryLoremIpsum.OnChanged = func(s string) {
		end.SetText(s)
	}
	v := container.NewVSplit(
		entryLoremIpsum,
		end,
	)
	v.Offset = 0.5
	return container.NewMax(v)

}

func Page2(w fyne.Window) *fyne.Container {

	return container.NewVBox(
		widget.NewButton("Hi2!", func() {

		}),
	)
}

func masterContent(w fyne.Window) *fyne.Container {
	clock := widget.NewLabel("")

	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)

	hello := widget.NewLabel("Hello Fyne 中文!")

	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

	text1 := canvas.NewText("Hello", green)
	text2 := canvas.NewText("There", green)
	text2.Move(fyne.NewPos(20, 20))
	content := container.NewHBox(text1, text2)

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

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
	entryLoremIpsum := widget.NewMultiLineEntry()
	entryLoremIpsum.SetText(loremIpsum)
	entryLoremIpsum.OnChanged = func(s string) {
		md.ParseMarkdown(s)
	}
	return container.NewVBox(entryLoremIpsum,
		md,
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
		clock,
		content,

		widget.NewCard("nihao", "nihao", canvas.NewText("niaho", green)),
		widget.NewButton("File Open With Filter (.jpg or .png)", func() {
			fd := dialog.NewFileOpen(func(closer fyne.URIReadCloser, err error) {
			}, w)
			fd.Show()
		}),
	)
}
