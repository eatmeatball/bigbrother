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
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
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
			return container.NewHBox(widget.NewIcon(theme.SearchReplaceIcon()), widget.NewLabel(""))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[0].(*widget.Icon).SetResource(theme.MenuIcon())
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(data[id])
		},
	)
	//contentHistory := map[int][]fyne.CanvasObject{}
	//listLeading.OnSelected = func(id widget.ListItemID) {
	//	if tLength > id {
	//		tItem := page.TList[id]
	//		title.SetText(tItem.Title)
	//		intro.SetText(tItem.Intro)
	//		var tmp []fyne.CanvasObject
	//		if historyObject, ok := contentHistory[id]; ok {
	//			tmp = historyObject
	//		} else {
	//			contentHistory[id] = []fyne.CanvasObject{tItem.View(w)}
	//			tmp = contentHistory[id]
	//		}
	//		content.Objects = tmp
	//		content.Refresh()
	//		appSetting.Index = id
	//		appSetting.saveToFile()
	//	}
	//}

	//listLeading.Select(page.Index)
	//listLeading.Select(appSetting.Index)

	if false {

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
	} else {

		leftList := []fyne.CanvasObject{}
		for index, node := range page.TList {
			callBack := func(index int, node page.ListNode) func() {
				return func() {
					var tmp []fyne.CanvasObject
					if node.V == nil {
						node.V = &[]fyne.CanvasObject{node.View(w)}
					}
					tmp = *node.V
					content.Objects = tmp
					content.Refresh()
					appSetting.Index = index
					appSetting.saveToFile()
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
	}

	w.Resize(fyne.NewSize(800, 600))

	w.ShowAndRun()
}
