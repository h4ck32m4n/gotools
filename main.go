package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/h4ck32m4n/gotools/digger"
	"github.com/h4ck32m4n/gotools/static"
)

func main() {
	diggerCmd := flag.NewFlagSet("digger", flag.ExitOnError)
	diggerSilent := diggerCmd.Bool("silent", false, "silent")
	diggerPurge := diggerCmd.Bool("purge", false, "purge folder /!\\ warning")
	diggerEcho := diggerCmd.Bool("echo", true, "write path in touched file")
	diggerFolder := diggerCmd.String("folder", digger.Home(), "folder to dig")
	diggerTarget := diggerCmd.String("target", "", "target where to clone digged folder")

	staticCmd := flag.NewFlagSet("static", flag.ExitOnError)
	staticSilent := staticCmd.Bool("silent", false, "silent")
	staticFolder := staticCmd.String("folder", digger.Home(), "build static in folder")

	if len(os.Args) < 2 {
		fmt.Println("expected 'digger' or 'static' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "digger":
		diggerCmd.Parse(os.Args[2:])
		if !*diggerSilent {
			fmt.Println("subcommand 'digger'")
			fmt.Println("  silent:", *diggerSilent)
			fmt.Println("  echo:", *diggerEcho)
			fmt.Println("  purge:", *diggerPurge)
			fmt.Println("  folder:", *diggerFolder)
			fmt.Println("  target:", *diggerTarget)
			fmt.Println("  tail:", diggerCmd.Args())
		}
		root := digger.Dig(*diggerFolder)
		if !*diggerSilent {
			println(root.Tree())
		}
		if *diggerTarget != "" {
			digger.Build(*diggerTarget)
			root.Build(*diggerTarget, 0)
			if !*diggerSilent {
				temp := digger.Dig(*diggerTarget)
				println(temp.Tree())
			}
		}
		if *diggerPurge {
			digger.Purge(*diggerFolder)
		}
	case "static":
		staticCmd.Parse(os.Args[2:])
		if !*diggerSilent {
			fmt.Println("subcommand 'static'")
			fmt.Println("  verbose:", *staticSilent)
			fmt.Println("  folder:", *staticFolder)
			fmt.Println("  tail:", staticCmd.Args())
		}
		header := "<!doctype html>\n<html>\n<head>\n"
		header += "<meta name=\"viewport\" content=\"width=device-width\"/>\n"
		header += "<meta charset=\"utf-8\">\n<title>" + "Title" + "</title>\n"
		header += "<link href=\"main.css\" rel=\"stylesheet\">\n</head>\n"
		body := "<body>\n<h1>" + "h1" + "</h1>\n"
		body += "<hr><h2>" + "h2" + "</h2><hr>"
		footer := "</body>\n</html>"
		css := "body { width: device-width; }\n"
		css += "h1 { width: device-width; font: 500 80px/1.5 Helvetica, Verdana, sans-serif; margin: 0; padding: 0; }\n"
		css += "h2 { width: device-width; font: 500 50px/1.5 Helvetica, Verdana, sans-serif; margin: 0; padding: 0; }\n"
		static.Make(*staticFolder, css, header, body, footer).Build()
		temp := digger.Dig(*staticFolder)
		if !*staticSilent {
			println(temp.Tree())
		}
	default:
		fmt.Println("expected 'digger' or 'static' subcommands")
		os.Exit(1)
	}

}
