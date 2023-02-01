package page

import (
	"bigbrother/resource"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"net/url"
)

func WelcomePage(w fyne.Window) *fyne.Container {
	logoObj := canvas.NewImageFromResource(BResource(resource.Logo))
	logoObj.FillMode = canvas.ImageFillContain
	logoObj.SetMinSize(fyne.NewSize(256, 256))
	linkT := widget.NewHyperlink("thh9.github.io", parseURL("https://thh9.github.io/"))

	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("Welcome to the BigBrother toolkit app", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		logoObj,
		container.NewHBox(
			linkT,
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
