package cli

import (
	"errors"
	"os"
)

//AppInterface basic interface for GoCLI cli applications
type AppInterface interface {
	AddCommand(name string, command interface{})
	Run() error
}

//NewApp creates a cli app with a provided default command
func NewApp(defaultCommand interface{}) AppInterface {
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
	state map[string]interface{}
}

//AddCommand adds a command to the cli with a given command name
//can handle both regular commands (Command type) or functional commands (FunctionalCommand type)
func (a *App) AddCommand(name string, command interface{}) {
	if a.commands == nil {
		a.commands = make(map[string]Command)
	}
	switch c := command.(type) {
	case Command:
		a.commands[name] = c
	case func(ctx AppInterface) error:
		a.commands[name] = &FunctionalCommandWrapper{c}
	}
}

//Run executes any command associated with the first argument passed to the application (after the application name itself)
//If no argument is passed it will use the default command
//If no command can be determined based on the argument or default, this will return an error saying that no command could be found
func (a *App) Run() error {
	cmd, err := a.findCommand()
	if err != nil {
		return err
	}

	s, ok := cmd.(CommandSetup)
	if ok {
		err := s.Setup(a)
		if err != nil {
			return err
		}
	}

	c, ok := cmd.(Command)
	if ok {
		err := c.Run(a)
		if err != nil {
			return err
		}
	}

	t, ok := cmd.(CommandTeardown)
	if ok {
		err := t.Teardown(a)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) findCommand() (interface{}, error) {
	if len(os.Args) > 1 {
		cmd, ok := a.commands[os.Args[1]]
		if ok {
			return cmd, nil
		}
	}
	cmd, ok := a.commands["default"]
	if ok {
		return cmd, nil
	}
	return nil, errors.New("could not find command")
}
