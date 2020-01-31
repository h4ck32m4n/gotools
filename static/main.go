package static

import "os"

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
	html := CreateFile(s.Path + "/" + "index" + ".html")
	WriteFile(html, s.Content["header"])
	WriteFile(html, s.Content["body"])
	WriteFile(html, s.Content["footer"])
	CloseFile(html)
	css := CreateFile(s.Path + "/" + "main" + ".css")
	WriteFile(css, s.Content["css"])
	CloseFile(css)
}

func CreateFile(path string) *os.File {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return file
}

func WriteFile(file *os.File, data string) {
	_, err := file.WriteString(data)
	if err != nil {
		panic(err)
	}
}

func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		panic(err)
	}
}
