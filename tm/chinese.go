package tm

import (
	"bigbrother/resource"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type MyTheme struct{}

var _ fyne.Theme = (*MyTheme)(nil)

func (m MyTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	//return theme.DefaultTheme().Color(n, v)
	if false {
		switch n {
		case theme.ColorNameBackground:
			return color.NRGBA{R: 0x24, G: 0x27, B: 0x2e, A: 0xff}
		case theme.ColorNameButton:
			return color.NRGBA{R: 0x24, G: 0x27, B: 0x2e, A: 0xff}
		case theme.ColorNameDisabledButton:
			return color.NRGBA{R: 0x26, G: 0x26, B: 0x26, A: 0xff}
		case theme.ColorNameDisabled:
			return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x42}
		case theme.ColorNameError:
			return color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff}
		case theme.ColorNameFocus:
			return color.NRGBA{R: 0x21, G: 0x96, B: 0xf3, A: 0x7f}
		case theme.ColorNameForeground:
			return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
		case theme.ColorNameHover:
			return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xf}
		case theme.ColorNameInputBackground:
			return color.NRGBA{R: 0x24, G: 0x27, B: 0x2e, A: 0x0}
		case theme.ColorNamePlaceHolder:
			return color.NRGBA{R: 0xb2, G: 0xb2, B: 0xb2, A: 0xff}
		case theme.ColorNamePressed:
			return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x66}
		case theme.ColorNamePrimary:
			return color.NRGBA{R: 0x42, G: 0x6c, B: 0x79, A: 0xff}
		case theme.ColorNameScrollBar:
			return color.NRGBA{R: 0x0, G: 0x0, B: 0x0, A: 0x26}
		case theme.ColorNameShadow:
			return color.NRGBA{R: 0x0, G: 0x0, B: 0x0, A: 0x66}
		default:
			return theme.DefaultTheme().Color(n, v)
		}
	}
	switch n {
	case theme.ColorNameBackground:
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameButton:
		return BlueGray300
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameDisabledButton:
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameDisabled:
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameError:
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameFocus:
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameForeground:
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameHover:
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameInputBackground:
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNamePlaceHolder:
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNamePressed:
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNamePrimary:
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameScrollBar:
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameSelection:
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameShadow:
		return theme.DefaultTheme().Color(n, v)
	default:
		return theme.DefaultTheme().Color(n, v)
	}
}
func (m MyTheme) Icon(name fyne.ThemeIconName) fyne.Resource {

	return theme.DefaultTheme().Icon(name)
}

func (m MyTheme) Font(style fyne.TextStyle) fyne.Resource {
	//return theme.DefaultTheme().Font(style)
	return &fyne.StaticResource{
		StaticName:    "HarmonyOS_Sans_SC_Regular.ttf",
		StaticContent: resource.HMttf,
	}
}

func (m MyTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
