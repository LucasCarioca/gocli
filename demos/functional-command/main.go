package main

import (
	"fmt"
	"github.com/LucasCarioca/gocli/cli"
)

func functionalCommand(_ cli.AppInterface) error {
	fmt.Println("This is the default functional command")
	return nil
}

func functionalCommand2(_ cli.AppInterface) error {
	fmt.Println("This is another functional command")
	return nil
}

func main() {
	app := cli.NewApp(functionalCommand)
	app.AddCommand("hello", functionalCommand2)
	app.AddCommand("inline", func() error {
		fmt.Println("This command is created inline")
		return nil
	})
	app.Run()
}
