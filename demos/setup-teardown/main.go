package main

import (
	"fmt"
	"github.com/LucasCarioca/gocli/cli"
)

type myCommand struct{}

//Setup function to run before running the command
func (*myCommand) Setup() error {
	fmt.Println("Setting up")
	return nil
}

//Run command to run
func (*myCommand) Run() error {
	fmt.Println("Running the command")
	return nil
}

//Teardown function to run after running the command
func (*myCommand) Teardown() error {
	fmt.Println("Tearing down")
	return nil
}

func main() {
	app := cli.NewApp(func() error {
		fmt.Println("default command")
		return nil
	})
	app.AddCommand("test", &myCommand{})
	app.Run()
}
