goCLI Commands are very simple and can be implemented in two different ways. The simplest is the functional command approach which is simply mapping a function to a command. The more full featured option is to create an implementation for the `Command` interface which can be more complex.


## Functional Commands

```go
type FunctionalCommand func() error
```

The simplest option is to just create a function and map it to a command in the app. You can do this by declaring the function seperately or passing an inline anonymous function.

For example we can create a simple function and use it as the default command:

```go

func MyFunctionalCommand() error {
    fmt.Println("Hello world")
}

func main() {
    app := cli.NewApp(MyFunctionalCommand)
    app.Run()
}
```

Or similarly we could pass it in line as an anonymous function:

```go
func main() {
    app := cli.NewApp(func() error {
        fmt.Println("Hello world")
    })
    app.Run()
}
```

You can also pass as many commands as you want to map:

```go
func main() {
    app := cli.NewApp(func() error {
        fmt.Println("This is default command")
    })
    app.AddCommand("hello", func() error {
        fmt.Println("This is hello command")
    })
    app.Run()
}
```

## Command Type

```go
type Command interface {
	Run() error
}
```

To implement the `Command` interface you need to create a struct with a public `Run` function that returns an `error`. 

For example: 

```go
type ExampleCommand struct{}

func (*ExampleCommand) Run() error {
	fmt.Println("This is the default command")
	return nil
}
```

This approach allows for more complex implementations like the following example using `flags`

```go
type options struct {
	destroy bool
	add     bool
	change  bool
}

type ExampleCommand struct {
    Name: string 
}

func (c *ExampleCommand) getOptions() *options {
	cmd := flag.NewFlagSet(c.Name, flag.ExitOnError)
	greet := cmd.Bool("greet", false, "Should the cli greet the user")
	cmd.Parse(os.Args[2:])
	return &options{
		greet: *greet,
	}
}

func (c *ExampleCommand) Run() error {
	options := c.getOptions()
    if options.greet {
        fmt.Println("Hello World")
    }
	return nil
}
```
