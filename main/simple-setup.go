package main

import (
	"fmt"
	"github.com/LucasCarioca/gocli/cli"
)

type simpleCommand struct {}

//Run command to run
func (*simpleCommand) Run() error {
	fmt.Println("Hello World")
	return nil
}


func main() {
	app := &cli.App{}
	app.AddCommand("hello", &simpleCommand{})
	app.Run()
}