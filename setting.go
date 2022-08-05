package main

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2/app"
	"github.com/spf13/cast"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Setting struct {
	Index int
}

var SettingPath string

var appSetting Setting

func LoadSetting() error {
	s := &app.SettingsSchema{}
	path := s.StoragePath()
	SettingPath = path
	file, err := os.Open(path) // #nosec
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(filepath.Dir(path), 0700)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	decode := json.NewDecoder(file)

	return decode.Decode(&appSetting)
}

func (s *Setting) saveToFile() error {
	path := SettingPath
	err := os.MkdirAll(filepath.Dir(path), 0700)
	if err != nil { // this is not an exists error according to docs
		return err
	}

	data, err := json.Marshal(&s)
	fmt.Println(cast.ToString(data), err)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, data, 0644)
}
