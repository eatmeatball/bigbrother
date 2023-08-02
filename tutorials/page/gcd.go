package page

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/cast"
)

func GcdPage(_ fyne.Window) fyne.CanvasObject {
	// init
	a := widget.NewEntry()
	b := widget.NewEntry()
	show := widget.NewEntry()
	a.SetText("1")
	b.SetText("1")
	// set style
	//show.Disable()
	// set event
	a.OnChanged = func(s string) {
		aNumber := cast.ToInt64(a.Text)
		bNumber := cast.ToInt64(b.Text)
		if aNumber == 0 || bNumber == 0 {
			show.SetText("0")
		} else {
			mcd := _gcd(aNumber, bNumber)
			show.SetText(fmt.Sprintf("%v/%v", aNumber/mcd, bNumber/mcd))
		}

	}
	b.OnChanged = func(s string) {
		aNumber := cast.ToInt64(a.Text)
		bNumber := cast.ToInt64(b.Text)
		if aNumber == 0 || bNumber == 0 {
			show.SetText("0")
		} else {
			mcd := _gcd(aNumber, bNumber)
			show.SetText(fmt.Sprintf("%v/%v", aNumber/mcd, bNumber/mcd))
		}
	}

	rows := container.NewGridWithRows(1,
		container.NewGridWithColumns(3, a, b, show),
	)

	return rows

}

func _gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return _gcd(b, a%b)
}
