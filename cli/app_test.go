package cli

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

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

	t.Run("Should call command and all the lifecycle hooks", func(t *testing.T) {
		fullCmd := newMockCmdFull()
		app := NewApp(&fullCmd)
		commandCalls, setupCalls, teardownCalls := fullCmd.getCalls()
		assert.NotNil(t, app, "Should return a proper app")
		assert.Equalf(t, 0, commandCalls, "Should not have executed default command")
		assert.Equalf(t, 0, setupCalls, "Should not have executed setup")
		assert.Equalf(t, 0, teardownCalls, "Should not have executed teardown")
		app.Run()
		commandCalls, setupCalls, teardownCalls = fullCmd.getCalls()
		assert.Equalf(t, 1, commandCalls, "Should have executed default command")
		assert.Equalf(t, 1, commandCalls, "Should have executed setup")
		assert.Equalf(t, 1, commandCalls, "Should have executed teardown")
	})

	t.Run("Should throw an error in the setup", func(t *testing.T) {
		fullCmd := newMockCmdFull()
		app := NewApp(&fullCmd)
		expectedError := "setup"
		fullCmd.setError(nil, errors.New(expectedError), nil)
		err := app.Run()
		assert.NotNil(t, app, "Should return a proper app")
		assert.NotNil(t, err, "Should return an error")
		assert.Equalf(t, expectedError, err.Error(), "Should return the proper error message")
	})

	t.Run("Should throw an error in the teardown", func(t *testing.T) {
		fullCmd := newMockCmdFull()
		app := NewApp(&fullCmd)
		expectedError := "teardown"
		fullCmd.setError(nil, nil, errors.New(expectedError))
		err := app.Run()
		assert.NotNil(t, app, "Should return a proper app")
		assert.NotNil(t, err, "Should return an error")
		assert.Equalf(t, expectedError, err.Error(), "Should return the proper error message")
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

	t.Run("Should support functional commands", func(t *testing.T) {
		functionalCommand := func() error {
			cmd.Run()
			return nil
		}
		app := NewApp(functionalCommand)
		app.AddCommand("test", functionalCommand)
		cmd.resetCalls()
		cmd.setError(nil)
		assert.NotNil(t, app, "Should return a proper app")
		assert.Equalf(t, 0, cmd.getCalls(), "Should not have executed the command")
		MockCLICall("app")
		app.Run()
		assert.Equalf(t, 1, cmd.getCalls(), "Should have executed the command")
		cmd.resetCalls()
		assert.Equalf(t, 0, cmd.getCalls(), "Should not have executed the command")
		MockCLICall("app test")
		app.Run()
		assert.Equalf(t, 1, cmd.getCalls(), "Should have executed the command")
	})
}
