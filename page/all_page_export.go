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
		{Title: "UrlEncode", Intro: "UrlEncode", View: makeTableTab},
		{Title: "timeButton", Intro: "time", View: getTimeButton},
		{Title: "clockPage", Intro: "clockPage", View: clockPage},
		{Title: "other", Intro: "other", View: masterContent},
		{Title: "about", Intro: "about", View: about},
	}
)
