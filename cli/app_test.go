package cli

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockCmd struct {
	calls     int
	mockError error
}

func newMockCmd() mockCmd {
	return mockCmd{calls: 0, mockError: nil}
}

func (c *mockCmd) Run() error {
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

func Test_app(t *testing.T) {
	cmd := newMockCmd()
	t.Run("Should create an app with a default command", func(t *testing.T) {
		app := NewApp(&cmd)
		cmd.resetCalls()
		cmd.setError(nil)
		assert.NotNil(t, app, "Should return a proper app")
		assert.Equalf(t, 0, cmd.getCalls(), "Should not have executed default command")
	})

	t.Run("Should create an app without a default command and throw an error when called", func(t *testing.T) {
		app := App{}
		err := app.Run()
		expectedError := "could not find command"
		assert.NotNil(t, app, "Should return a proper app")
		assert.NotNil(t, err, "Should return an error")
		assert.Equalf(t, expectedError, err.Error(), "Should return the proper error message")
	})

	t.Run("Should call default command if run without arguments", func(t *testing.T) {
		app := NewApp(&cmd)
		cmd.resetCalls()
		cmd.setError(nil)
		assert.NotNil(t, app, "Should return a proper app")
		assert.Equalf(t, 0, cmd.getCalls(), "Should not have executed default command")
		app.Run()
		assert.Equalf(t, 1, cmd.getCalls(), "Should have executed default command")
	})

	t.Run("Should run the proper command based on the command line arguments", func(t *testing.T) {
		newCmd := newMockCmd()
		app := NewApp(&cmd)
		app.AddCommand("test", &newCmd)
		newCmd.resetCalls()
		cmd.resetCalls()
		cmd.setError(nil)
		newCmd.setError(nil)
		assert.NotNil(t, app, "Should return a proper app")
		assert.Equalf(t, 0, cmd.getCalls(), "Should not have executed default command")
		assert.Equalf(t, 0, newCmd.getCalls(), "Should not have executed test command")
		MockCLICall("app test")
		app.Run()
		assert.Equalf(t, 0, cmd.getCalls(), "Should not have executed default command")
		assert.Equalf(t, 1, newCmd.getCalls(), "Should have executed test command")
	})
}
