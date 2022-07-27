package main

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

//go:embed HarmonyOS_Sans_SC_Regular.ttf
var hmTTf []byte

func main() {
	a := app.New()
	a.Settings().SetTheme(&myTheme{})
	w := a.NewWindow("Hello")
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
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))
	w.Resize(fyne.NewSize(480, 480))

	w.ShowAndRun()
}

type myTheme struct{}

var _ fyne.Theme = (*myTheme)(nil)

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return color.White
		}
		return color.Black
	}

	return theme.DefaultTheme().Color(name, variant)
}
func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {

	return theme.DefaultTheme().Icon(name)
}

func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	return &fyne.StaticResource{
		StaticName:    "HarmonyOS_Sans_SC_Regular.ttf",
		StaticContent: hmTTf,
	}
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
