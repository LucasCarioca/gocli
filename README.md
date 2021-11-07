# GoCLI

![GitHub release (latest by date)](https://img.shields.io/github/v/release/LucasCarioca/gocli)
![GitHub Release Date](https://img.shields.io/github/release-date/LucasCarioca/gocli)
![GitHub all releases](https://img.shields.io/github/downloads/LucasCarioca/gocli/total)
[![Go Reference](https://pkg.go.dev/badge/github.com/LucasCarioca/gocli.svg)](https://pkg.go.dev/github.com/LucasCarioca/gocli)

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/LucasCarioca/gocli/CI?label=CI)
[![Coverage Status](https://coveralls.io/repos/github/LucasCarioca/gocli/badge.svg?branch=main)](https://coveralls.io/github/LucasCarioca/gocli?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/LucasCarioca/gocli)](https://goreportcard.com/report/github.com/LucasCarioca/gocli)

[![DeepSource](https://deepsource.io/gh/LucasCarioca/gocli.svg/?label=active+issues&show_trend=true&token=z8O_Knm5SXRE-QTR3IVRiOrP)](https://deepsource.io/gh/LucasCarioca/gocli/?ref=repository-badge)

A lightweight tool for creating simple CLIs in Go.

## Getting Started

Fist you need to add the dependency to your project.

```shell
go get github.com/LucasCarioca/gocli
# or to target a specific version
go get github.com/LucasCarioca/gocli@v0.2.0
```

### Usage

To start you will want to create your first command. Commands are any struct that conform to the `Command` interface.
They require only one method which is the `Run` method and serves as the entry point for the command.

```go
type Command interface {
    Run() error
}
```

Example command implementation

```go
type HelloWorldCommand struct {}

func (*HelloWorldCommand) Run() error {
    fmt.Println("Hello World")
    return nil
}
```

Finally, we will create our application and attach our command to it:

```go
import "github.com/LucasCarioca/gocli/cli"

func main() {
    app := &cli.App{}
    app.AddCommand("hello", &HelloWorldCommand{})
    app.Run()
}
```

you can now validate that your cli is working

```shell
go build main.go -o mycli
mycli hello
# You should then get the output "Hello World"
```

You can also set the command as default using the `NewApp` function

```go
import "github.com/LucasCarioca/gocli/cli"

func main() {
    helloCommand := &HelloWorldCommand{}
    app := &cli.NewApp(helloCommand)
    app.Run()
}
```

Let's test it again

```shell
go build main.go -o mycli
mycli
# You should then get the output "Hello World"
```
