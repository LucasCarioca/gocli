package main

import (
	"fmt"
	"github.com/LucasCarioca/gocli/cli"
)

type helloCommand struct {}

//Run command to run
func (*helloCommand) Run() error {
	fmt.Println("Hello World")
	return nil
}


func main() {
	app := &cli.App{}
	app.AddCommand("hello", &helloCommand{})
	app.Run()
}