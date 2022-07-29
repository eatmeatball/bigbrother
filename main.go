package main

import (
	_ "embed"
	"fmt"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

//go:embed HarmonyOS_Sans_SC_Regular.ttf
var hmTTf []byte

//go:embed resource/logo.png
var logo []byte

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

	w.SetCloseIntercept(func() {
		w.Hide()
	})
	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("MyApp",
			fyne.NewMenuItem("Show", func() {
				w.Show()
			}))
		desk.SetSystemTrayMenu(m)
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
	tLength := len(tList)
	data := make([]string, tLength)
	for i, node := range tList {
		data[i] = node.Title
	}
	content := container.NewMax()
	content.Objects = []fyne.CanvasObject{WelcomePage(w)}

	listLeading := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(data[id])
		},
	)
	title := widget.NewLabel("")
	intro := widget.NewLabel("")
	listLeading.OnSelected = func(id widget.ListItemID) {
		if tLength > id {
			tItem := tList[id]
			title.SetText(tItem.Title)
			title.SetText(tItem.Intro)
			content.Objects = []fyne.CanvasObject{tItem.View(w)}
			content.Refresh()
		}
	}
	listLeading.Select(0)

	masterContent := container.NewHSplit(
		listLeading,
		container.NewBorder(

			container.NewVBox(
				title,
				widget.NewSeparator(),
				intro,
			),

			nil, nil, nil,

			content,
		),
	)
	masterContent.Offset = 0.2
	w.SetContent(masterContent)

	w.Resize(fyne.NewSize(600, 480))

	w.ShowAndRun()
}
