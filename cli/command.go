package cli

//Command interface that cli commands should conform to
type Command interface {
	Run() error
}

//CommandSetup interface for creating commands with setup stage
type CommandSetup interface {
	Setup() error
}

//CommandTeardown interface for creating commands with teardown stage
type CommandTeardown interface {
	Teardown() error
}
