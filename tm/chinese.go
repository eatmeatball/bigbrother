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
	if n == theme.ColorNameBackground {
		if v == theme.VariantLight {
			return color.White
		}
		return color.Black
	}
	switch n {
	case theme.ColorNameBackground:
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameButton:
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
