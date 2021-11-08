#goCLI

![goCLI](assets/gocli.png){ width="100px" align="left" }
A simple and lightweight cli framework for go applications. Get your cli setup in seconds.

![GitHub release (latest by date)](https://img.shields.io/github/v/release/LucasCarioca/gocli)
![GitHub Release Date](https://img.shields.io/github/release-date/LucasCarioca/gocli)
[![Go Reference](https://pkg.go.dev/badge/github.com/LucasCarioca/gocli.svg)](https://pkg.go.dev/github.com/LucasCarioca/gocli)
[![Go Report Card](https://goreportcard.com/badge/github.com/LucasCarioca/gocli)](https://goreportcard.com/report/github.com/LucasCarioca/gocli)

## Getting started

Add the dependency `go get github.com/LucasCarioca/gocli`

Your first cli

```go
package main

import "github.com/LucasCarioca/gocli/cli"

func main() {
    app := cli.NewApp(func() error {
        fmt.Println("Hello world")    	
    })
    app.Run()
}
```

Build and run

```shell
go build main.go -o mycli
mycli
# Hello World
```


