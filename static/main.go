package static

import (
	"os"
	"github.com/h4ck32m4n/gotools/digger"
)

type Static struct {
	Path    string
	Content map[string]string
}

func Create(path string) *Static {
	return &Static{Path: path}
}

func Make(path string, css string, header string, body string, footer string) *Static {
	static := Create(path)
	static.Content = make(map[string]string)
	static.Content["css"] = css
	static.Content["header"] = header
	static.Content["body"] = body
	static.Content["footer"] = footer
	return static
}

func (s *Static) Build() {
	digger.TouchEcho(s.Path + "/" + "index" + ".html", s.Content["header"] + s.Content["body"] + s.Content["footer"])
	digger.TouchEcho(s.Path + "/" + "main" + ".css", s.Content["css"])
}