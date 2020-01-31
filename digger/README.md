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
// dig into $HOME
root := digger.Dig(digger.Home())
// pseudo tree bash
println(root.Tree())
// build honeypot of digged folder
target := "./honeypot"
// init honeypot
digger.Build(target)
// build honeypot
root.Build(target, 0)
// dig and print honeypot
temp := digger.Dig(target)
println(temp.Tree())
// purge honeypot
digger.Purge(target)
```
## Git ##
```bash
$ git clone https://github.com/h4ck32m4n/gotools.git
```
