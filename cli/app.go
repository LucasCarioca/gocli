package cli

import "os"

type AppInterface interface {
	AddCommand(name string, command Command)
	Run()
}

type App struct {
	commands map[string]Command
}

func (a *App) AddCommand(name string, command Command) {
	a.commands[name] = command
}

func (a *App) Run() {
	a.commands[os.Args[1]].Run()
}
