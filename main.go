package main

import (
	"bigbrother/page"
	"bigbrother/tm"
	_ "embed"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
	"time"
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
	// init config
	LoadSetting()

	a := app.New()

	logLifecycle(a)

	a.Settings().SetTheme(&tm.MyTheme{})

	w := a.NewWindow("bigBrother")
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
	// 桌面系统设置托盘
	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("bigBrother",
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

	var leftList []fyne.CanvasObject
	for index, node := range page.TList {
		callBack := func(index int, node page.ListNode) func() {
			return func() {
				start := time.Now()
				var tmp []fyne.CanvasObject
				if node.V == nil {
					node.V = &[]fyne.CanvasObject{node.View(w)}
				}
				tmp = *node.V
				content.Objects = tmp
				content.Refresh()
				appSetting.Index = index
				appSetting.saveToFile()
				end:= time.Now()
				fmt.Println(end.Sub(start))
			}
		}(index, node)
		leftList = append(leftList, widget.NewButton(node.Title, callBack))
	}
	rectangles := canvas.NewRectangle(color.RGBA{R: 49, G: 48, B: 66})
	left := container.NewVScroll(container.NewVBox(leftList...))

	leftList[appSetting.Index].(*widget.Button).OnTapped()

	leftContent := container.NewMax(
		left,
		rectangles,
	)
	masterContent := container.NewBorder(
		nil, nil, leftContent, nil,
		content,
	)
	w.SetContent(masterContent)

	w.Resize(fyne.NewSize(800, 600))

	w.ShowAndRun()
}
