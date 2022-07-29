package main

import (
	"bigbrother/page"
	"bigbrother/tm"
	_ "embed"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
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

func main() {
	a := app.New()
	logLifecycle(a)
	a.Settings().SetTheme(&tm.MyTheme{})
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
	tLength := len(page.TList)
	data := make([]string, tLength)
	for i, node := range page.TList {
		data[i] = node.Title
	}
	content := container.NewMax()
	title := widget.NewLabel("")
	intro := widget.NewLabel("")
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
	listLeading.OnSelected = func(id widget.ListItemID) {
		if tLength > id {
			tItem := page.TList[id]
			title.SetText(tItem.Title)
			intro.SetText(tItem.Intro)
			content.Objects = []fyne.CanvasObject{tItem.View(w)}
			content.Refresh()
		}
	}

	listLeading.Select(2)

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
