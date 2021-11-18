## Command Type

```go
type Command interface {
	Run(ctx AppInterface) error
}
```

To implement the `Command` interface you need to create a struct with a public `Run` function that returns an `error`.

For example:

```go
type ExampleCommand struct{}

func (*ExampleCommand) Run(_ cli.AppInterface) error {
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

func (c *ExampleCommand) Run(_ cli.AppInterface) error {
	options := c.getOptions()
    if options.greet {
        fmt.Println("Hello World")
    }
	return nil
}
```