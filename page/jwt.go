package page

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/cast"
)

func Page1(w fyne.Window) *fyne.Container {

	end := widget.NewMultiLineEntry()
	end.Wrapping = fyne.TextWrapBreak
	entryLoremIpsum := widget.NewMultiLineEntry()
	entryLoremIpsum.Wrapping = fyne.TextWrapBreak
	entryLoremIpsum.OnChanged = func(s string) {
		jwtStr := s
		token, _ := jwt.Parse(jwtStr,
			// 防止bug从我做起
			func([]byte) func(token *jwt.Token) (i any, e error) {
				return func(token *jwt.Token) (i any, e error) {
					return "", nil
				}
			}([]byte("")))

		//if err != nil {
		//	end.SetText(cast.ToString(err.Error()))
		//	return
		//}

		if token == nil || token.Claims == nil {
			end.SetText("无法解析")
			return
		}

		//data, _ := json.Marshal(token.Claims)
		data, _ := json.MarshalIndent(token.Claims, "", "  ")
		result := cast.ToString(data)
		end.SetText(result)
	}

	entryLoremIpsum.SetText("")
	v := container.NewVSplit(
		entryLoremIpsum,
		end,
	)
	v.Offset = 0.5
	return container.NewMax(v)

}
