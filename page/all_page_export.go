package page

import "fyne.io/fyne/v2"

func BResource(data []byte) *fyne.StaticResource {
	return &fyne.StaticResource{
		StaticName:    "logo.png",
		StaticContent: data,
	}
}

type ListNode struct {
	Title, Intro string
	View         func(w fyne.Window) *fyne.Container
}

var (
	TList = []ListNode{
		{Title: "Welcome", Intro: "Welcome", View: WelcomePage},
		{Title: "Jwt", Intro: "page1", View: Page1},
		{Title: "UrlEncode", Intro: "UrlEncode", View: urlencode},
		{Title: "page2", Intro: "page2", View: Page2},
		{Title: "page1", Intro: "page1", View: masterContent},
		{Title: "about", Intro: "about", View: about},
	}
)
