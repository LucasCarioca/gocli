package cli

//FunctionalCommand is any simple function that returns an optional error
type FunctionalCommand func() error

//FunctionalCommandWrapper an adapter to support running functional commands
type FunctionalCommandWrapper struct {
	Command FunctionalCommand
}

//Run executes the wrapped functional command
func (c *FunctionalCommandWrapper) Run() error {
	return c.Command()
}
