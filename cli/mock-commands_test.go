package cli

type mockCmd struct {
	calls     int
	mockError error
}

func newMockCmd() mockCmd {
	return mockCmd{calls: 0, mockError: nil}
}

func (c *mockCmd) Run(_ AppInterface) error {
	c.calls = c.calls + 1
	return c.mockError
}

func (c *mockCmd) resetCalls() {
	c.calls = 0
}

func (c *mockCmd) getCalls() int {
	return c.calls
}

func (c *mockCmd) setError(mockError error) {
	c.mockError = mockError
}

type mockCmdFull struct {
	commandCalls      int
	setupCalls        int
	teardownCalls     int
	mockCommandError  error
	mockSetupError    error
	mockTeardownError error
}

func newMockCmdFull() mockCmdFull {
	return mockCmdFull{
		commandCalls:      0,
		setupCalls:        0,
		teardownCalls:     0,
		mockCommandError:  nil,
		mockSetupError:    nil,
		mockTeardownError: nil,
	}
}

func (c *mockCmdFull) Setup(_ AppInterface) error {
	c.setupCalls = c.setupCalls + 1
	return c.mockSetupError
}

func (c *mockCmdFull) Teardown(_ AppInterface) error {
	c.teardownCalls = c.teardownCalls + 1
	return c.mockTeardownError
}

func (c *mockCmdFull) Run(_ AppInterface) error {
	c.commandCalls = c.commandCalls + 1
	return c.mockCommandError
}

func (c *mockCmdFull) resetCalls() {
	c.commandCalls = 0
	c.setupCalls = 0
	c.teardownCalls = 0
}

func (c *mockCmdFull) getCalls() (commandCalls, setupCalls, teardownCalls int) {
	return c.commandCalls, c.setupCalls, c.teardownCalls
}

func (c *mockCmdFull) setError(mockCommandError, mockSetupError, mockTeardownError error) {
	c.mockCommandError = mockCommandError
	c.mockSetupError = mockSetupError
	c.mockTeardownError = mockTeardownError
}
