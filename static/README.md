# gotools - static #
```bash
$ go get github.com/h4ck32m4n/gotools/static
```
## TODO ##
- 



## [static](main.go) ##
```bash
static
├── go.mod
├── main.go
└── README.md
```
### Usage ###
```golang
package main

import "github.com/h4ck32m4n/gotools/static"

func Header() string {
	header := "<!doctype html>\n<html>\n<head>\n"
	header += "<meta name=\"viewport\" content=\"width=device-width\"/>\n"
	header += "<meta charset=\"utf-8\">\n<title>" + "Title" + "</title>\n"
	header += "<link href=\"main.css\" rel=\"stylesheet\">\n</head>\n"
	return header
}

func Body() string {
	body := "<body>\n<h1>" + "h1" + "</h1>\n"
	body += "<hr><h2>" + "h2" + "</h2><hr>"
	return body
}

func Footer() string {
	footer := "</body>\n</html>"
	return footer
}

func CSS() string {
	css := "body { width: device-width; }\n"
	css += "h1 { width: device-width; font: 500 80px/1.5 Helvetica, Verdana, sans-serif; margin: 0; padding: 0; }\n"
	css += "h2 { width: device-width; font: 500 50px/1.5 Helvetica, Verdana, sans-serif; margin: 0; padding: 0; }\n"
	return css
}


func main() {
    static.Make("/path/to/folder/", CSS(), Header(), Body(), Footer()).Build()
}
```
## Git ##
```bash
$ git clone https://github.com/h4ck32m4n/gotools.git
```
