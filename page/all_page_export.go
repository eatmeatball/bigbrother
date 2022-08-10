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
	V            *[]fyne.CanvasObject
}

var (
	TList = []ListNode{
		{Title: "Welcome", Intro: "Welcome", View: WelcomePage},
		{Title: "Jwt解析", Intro: "jwt解析", View: JwtParsePage},
		{Title: "UrlEncode", Intro: "UrlEncode", View: urlEncodePage},
		{Title: "时间转化", Intro: "timePage", View: timePage},
		{Title: "json2yaml", Intro: "json2yaml", View: json2yaml},
		{Title: "clockPage", Intro: "clockPage", View: clockPage},
		{Title: "timeButton", Intro: "time", View: getTimeButton},
		{Title: "other", Intro: "other", View: masterContent},
		{Title: "about", Intro: "about", View: about},
		//{Title: "table", Intro: "table", View: makeTableTab},
	}
	Index = 3
)
