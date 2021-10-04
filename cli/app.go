package cli

import (
	"errors"
	"os"
)

type AppInterface interface {
	AddCommand(name string, command Command)
	Run() error
}

func NewApp(defaultCommand Command) AppInterface {
	app := &App{
		commands: map[string]Command{},
	}
	app.AddCommand("default", defaultCommand)
	return app
}

type App struct {
	commands map[string]Command
}

func (a *App) AddCommand(name string, command Command) {
	a.commands[name] = command
}

func (a *App) Run() error {
	cmd, ok := a.commands[os.Args[1]]
	if ok {
		return cmd.Run()
	}

	cmd, ok = a.commands["default"]
	if ok {
		return cmd.Run()
	}

	return errors.New("could not find command")
}
