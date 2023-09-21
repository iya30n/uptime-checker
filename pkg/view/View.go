package view

import (
	"os"
	"strings"
	"uptime/pkg/logger"
)

type View struct {
	Path string
	Data map[string]string
}

func (vt *View) Render() string {
	view := getFile(vt.Path)

	for k, v := range vt.Data {
		view = strings.Replace(view, k, v, -1)
	}

	return view
}

func getFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		logger.Error(err.Error())
		return ""
	}

	return string(content)
}