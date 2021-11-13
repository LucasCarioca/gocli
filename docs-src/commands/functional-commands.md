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
