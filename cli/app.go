package cli

import (
	"errors"
	"os"
)

//AppInterface basic interface for GoCLI cli applications
type AppInterface interface {
	AddCommand(name string, command Command)
	Run() error
}

//NewApp creates a cli app with a provided default command
func NewApp(defaultCommand Command) AppInterface {
	app := &App{
		commands: map[string]Command{
			"version": &VersionCommand{},
		},
	}
	app.AddCommand("default", defaultCommand)
	return app
}

//App basic structure for managing a cli
type App struct {
	commands map[string]Command
}

//AddCommand adds a command to the cli with a given command name
func (a *App) AddCommand(name string, command Command) {
	if a.commands == nil {
		a.commands = make(map[string]Command)
	}
	a.commands[name] = command
}

//Run executes any command associated with the first argument passed to the application (after the application name itself)
//If no argument is passed it will use the default command
//If no command can be determined based on the argument or default, this will return an error saying that no command could be found
func (a *App) Run() error {
	if len(os.Args) > 1 {
		cmd, ok := a.commands[os.Args[1]]
		if ok {
			return cmd.Run()
		}
	}

	cmd, ok := a.commands["default"]
	if ok {
		return cmd.Run()
	}

	return errors.New("could not find command")
}
