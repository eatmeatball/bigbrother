package page

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/cast"
)

func JwtParsePage(w fyne.Window) fyne.CanvasObject {

	show := widget.NewMultiLineEntry()
	show.Wrapping = fyne.TextWrapBreak
	//show.Disable()
	input := widget.NewMultiLineEntry()
	input.Wrapping = fyne.TextWrapBreak
	input.OnChanged = func(s string) {
		jwtStr := s
		token, _ := jwt.Parse(jwtStr,
			// 防止bug从我做起
			func([]byte) func(token *jwt.Token) (i any, e error) {
				return func(token *jwt.Token) (i any, e error) {
					return "", nil
				}
			}([]byte("")))

		//if err != nil {
		//	show.SetText(cast.ToString(err.Error()))
		//	return
		//}

		if token == nil || token.Claims == nil {
			show.SetText("无法解析")
			return
		}

		//data, _ := json.Marshal(token.Claims)
		data, _ := json.MarshalIndent(token.Claims, "", "  ")
		result := cast.ToString(data)
		show.SetText(result)
	}

	input.SetText(`eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoidGhoIiwiZXhwIjoxNjYwNzM0OTMwLCJuYmYiOjE2NjAxMzAxMzB9.tnTQYaaM7FfJhah9veVNS0OSkh1q4hKtP1UUUgESCQY`)
	v := container.NewVSplit(
		input,
		show,
	)
	v.Offset = 0.5
	return container.NewMax(v)

}
