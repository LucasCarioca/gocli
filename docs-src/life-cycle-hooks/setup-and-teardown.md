## Setup and Teardown

```go
type CommandSetup interface {
	Setup(ctx AppInterface) error
}

type CommandTeardown interface {
	Teardown(ctx AppInterface) error
}
```


If you would like to organize your command into steps you can also use the provided `CommandSetup` and `CommandTeardown` types.
If you command conforms to these interfaces you can add functions to be executed before and after the main `Run` function is executed.

Both interfaces are optional and can be combined.

For example

```go
type myCommand struct{}

//Setup function to run before running the command
func (*myCommand) Setup(_ cli.AppInterface) error {
	fmt.Println("Setting up")
	return nil
}

//Run command to run
func (*myCommand) Run(_ cli.AppInterface) error {
	fmt.Println("Running the command")
	return nil
}

//Teardown function to run after running the command
func (*myCommand) Teardown(_ cli.AppInterface) error {
	fmt.Println("Tearing down")
	return nil
}
```

The above example creates a command that implements `Command`, `CommandSetup`, and `CommandTeardown` interfaces. This allows it to use both lifecycle hooks. 