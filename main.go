package main

import "github.com/h4ck32m4n/gotools/digger"

func main() {
	root := digger.Dig(digger.Home())
	println(root.Tree())
}
