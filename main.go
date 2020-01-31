package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/h4ck32m4n/gotools/digger"
)

func main() {
	diggerCmd := flag.NewFlagSet("digger", flag.ExitOnError)
	diggerSilent := diggerCmd.Bool("silent", false, "silent")
	diggerPurge := diggerCmd.Bool("purge", false, "purge folder /!\\ warning")
	diggerFolder := diggerCmd.String("folder", digger.Home(), "folder to dig")
	diggerTarget := diggerCmd.String("target", "", "target where to clone digged folder")

	staticCmd := flag.NewFlagSet("static", flag.ExitOnError)
	staticSilent := staticCmd.Bool("silent", false, "silent")
	staticBuild := staticCmd.Bool("build", false, "build static")

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
			temp := digger.Dig(*diggerTarget)
			if !*diggerSilent {
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
			fmt.Println("  build:", *staticBuild)
			fmt.Println("  tail:", staticCmd.Args())
		}
	default:
		fmt.Println("expected 'digger' or 'static' subcommands")
		os.Exit(1)
	}

}
