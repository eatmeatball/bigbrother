package main

import (
	_ "embed"
	"fmt"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"image/color"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

//go:embed HarmonyOS_Sans_SC_Regular.ttf
var hmTTf []byte

const (
	loremIpsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque quis consectetur nisi. Suspendisse id interdum felis.
Sed egestas eget tellus eu pharetra. Praesent pulvinar sed massa id placerat. Etiam sem libero, semper vitae consequat ut, volutpat id mi.
Mauris volutpat pellentesque convallis. Curabitur rutrum venenatis orci nec ornare. Maecenas quis pellentesque neque.
Aliquam consectetur dapibus nulla, id maximus odio ultrices ac. Sed luctus at felis sed faucibus. Cras leo augue, congue in velit ut, mattis rhoncus lectus.

Praesent viverra, mauris ut ullamcorper semper, leo urna auctor lectus, vitae vehicula mi leo quis lorem.
Nullam condimentum, massa at tempor feugiat, metus enim lobortis velit, eget suscipit eros ipsum quis tellus. Aenean fermentum diam vel felis dictum semper.
Duis nisl orci, tincidunt ut leo quis, luctus vehicula diam. Sed velit justo, congue id augue eu, euismod dapibus lacus. Proin sit amet imperdiet sapien.
Mauris erat urna, fermentum et quam rhoncus, fringilla consequat ante. Vivamus consectetur molestie odio, ac rutrum erat finibus a.
Suspendisse id maximus felis. Sed mauris odio, mattis eget mi eu, consequat tempus purus.`
)

func logLifecycle(a fyne.App) {
	a.Lifecycle().SetOnStarted(func() {
		log.Println("Lifecycle: Started")
	})
	a.Lifecycle().SetOnStopped(func() {
		log.Println("Lifecycle: Stopped")
	})
	a.Lifecycle().SetOnEnteredForeground(func() {
		log.Println("Lifecycle: Entered Foreground")
	})
	a.Lifecycle().SetOnExitedForeground(func() {
		log.Println("Lifecycle: Exited Foreground")
	})
}

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

func main() {
	str := binding.NewString()
	_ = str.Set("Initial value")

	a := app.New()
	logLifecycle(a)
	a.Settings().SetTheme(&myTheme{})
	w := a.NewWindow("Hello")
	w.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("Edit", fyne.NewMenuItem("edit", func() {
			fmt.Println("edit")
		}),
		),
	))
	w.SetMaster()

	clock := widget.NewLabel("")
	w.SetContent(clock)
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("MyApp",
			fyne.NewMenuItem("Show", func() {
				w.Show()
			}))
		desk.SetSystemTrayMenu(m)
	}
	w.SetCloseIntercept(func() {
		w.Hide()
	})
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

	entryLoremIpsum := widget.NewMultiLineEntry()
	entryLoremIpsum.SetText(loremIpsum)
	entryLoremIpsum.OnChanged = func(s string) {
		md.ParseMarkdown(s)
	}

	w.SetContent(container.NewGridWithRows(3,
		layout.NewSpacer(),
		container.NewGridWithColumns(3,
			layout.NewSpacer(),
			container.NewGridWithColumns(1,
				widget.NewButton("New Project", func() {
				}),
			),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	))
	masterContent := container.NewHSplit(
		container.NewVBox(hello),
		container.NewBorder(
			container.NewVBox(widget.NewLabel("Component name"), widget.NewSeparator(), widget.NewLabel("intro")),
			nil, nil, nil,
			container.NewVBox(entryLoremIpsum,
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
			),
		),
	)

	w.SetContent(masterContent)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	w.Resize(fyne.NewSize(480, 480))

	w.ShowAndRun()
}
