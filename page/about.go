package page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func about(w fyne.Window) *fyne.Container {
	text := `
# about
## What's it？

It's a gui tool with go.

---

* Item1 in _three_ segments
* Item2
* Item3
---

* Item1 in _three_ segments
* Item2
* Item3
---

* Item1 in _three_ segments
* Item2
* Item3

`
	about := widget.NewRichTextFromMarkdown(text)
	about.Wrapping = fyne.TextWrapWord
	//about.Scroll = container.ScrollBoth

	for i := range about.Segments {
		if seg, ok := about.Segments[i].(*widget.TextSegment); ok {
			seg.Style.Alignment = fyne.TextAlignCenter
		}
		if seg, ok := about.Segments[i].(*widget.HyperlinkSegment); ok {
			seg.Alignment = fyne.TextAlignCenter
		}
	}
	return container.NewVBox(
		about,
	)
}
