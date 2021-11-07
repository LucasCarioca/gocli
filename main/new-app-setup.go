package main

import (
	"fmt"
	"github.com/LucasCarioca/gocli/cli"
)

type defaultCommand struct{}

//Run command to run
func (*defaultCommand) Run() error {
	fmt.Println("This is the default command")
	return nil
}

func main() {
	app := cli.NewApp(&defaultCommand{})
	app.Run()
}
