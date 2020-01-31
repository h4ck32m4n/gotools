# gotools - digger #
```bash
$ go get github.com/h4ck32m4n/gotools/digger
```
## TODO ##
- 

## [digger](main.go) ##
```bash
digger
├── go.mod
├── main.go
└── README.md
```
### Usage ###
```golang
package main

import "github.com/h4ck32m4n/gotools/digger"

func main() {
    // dig and print $HOME
    root := digger.Dig(digger.Home())
    println(root.Tree())
    
    // init honeypot
    target := "./honeypot"
    digger.Build(target)
    
    // build honeypot
    root.Build(target, 0)
    temp := digger.Dig(target)
    println(temp.Tree())
    
    // purge honeypot
    digger.Purge(target)
    temp := digger.Dig(target)
    println(temp.Tree())
}
```
## Git ##
```bash
$ git clone https://github.com/h4ck32m4n/gotools.git
```
