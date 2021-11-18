package cli

//Command interface that cli commands should conform to
type Command interface {
	Run(ctx AppInterface) error
}

//CommandSetup interface for creating commands with setup stage
type CommandSetup interface {
	Setup(ctx AppInterface) error
}

//CommandTeardown interface for creating commands with teardown stage
type CommandTeardown interface {
	Teardown(ctx AppInterface) error
}
