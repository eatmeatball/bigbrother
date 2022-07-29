package main

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/cast"
	"image/color"
	"net/url"
	"time"
)

func BResource(data []byte) *fyne.StaticResource {
	return &fyne.StaticResource{
		StaticName:    "logo.png",
		StaticContent: data,
	}
}

var (
	tList = []ListNode{
		{Title: "Welcome", Intro: "Welcome", View: WelcomePage},
		{Title: "Jwt", Intro: "page1", View: Page1},
		{Title: "page2", Intro: "page2", View: Page2},
		{Title: "page1", Intro: "page1", View: masterContent},
	}
)

type ListNode struct {
	Title, Intro string
	View         func(w fyne.Window) *fyne.Container
}

func WelcomePage(w fyne.Window) *fyne.Container {
	logoObj := canvas.NewImageFromResource(BResource(logo))
	logoObj.FillMode = canvas.ImageFillContain
	logoObj.SetMinSize(fyne.NewSize(256, 256))
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("Welcome to the Fyne toolkit demo app", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		logoObj,
		container.NewHBox(
			widget.NewHyperlink("fyne.io", parseURL("https://fyne.io/")),
			widget.NewLabel("-"),
			widget.NewHyperlink("documentation", parseURL("https://developer.fyne.io/")),
			widget.NewLabel("-"),
			widget.NewHyperlink("sponsor", parseURL("https://fyne.io/sponsor/")),
		),
		widget.NewLabel(""), // balance the header on the tutorial screen we leave blank on this content
	))
}

func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return link
}

func Page1(w fyne.Window) *fyne.Container {

	end := widget.NewMultiLineEntry()
	end.Wrapping = fyne.TextWrapBreak

	entryLoremIpsum := widget.NewMultiLineEntry()
	entryLoremIpsum.Wrapping = fyne.TextWrapBreak
	entryLoremIpsum.SetText(loremIpsum)
	entryLoremIpsum.OnChanged = func(s string) {
		jwtStr := s
		token, _ := jwt.Parse(jwtStr,
			// 防止bug从我做起
			func([]byte) func(token *jwt.Token) (i any, e error) {
				return func(token *jwt.Token) (i any, e error) {
					return "", nil
				}
			}([]byte("")))

		//if err != nil {
		//	end.SetText(cast.ToString(err.Error()))
		//	return
		//}

		if token == nil || token.Claims == nil {
			end.SetText("无法解析")
			return
		}

		//data, _ := json.Marshal(token.Claims)
		data, _ := json.MarshalIndent(token.Claims, "", "  ")
		result := cast.ToString(data)
		end.SetText(result)
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
