package cli

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_version_command(t *testing.T) {
	t.Run("Should return the provided version", func(t *testing.T) {
		customOut = new(bytes.Buffer)
		cmd := VersionCommand{
			Version: "v0.0.0",
		}
		cmd.Run()
		assert.Equal(t, "Current version: vv0.0.0\n", customOut.(*bytes.Buffer).String(), "Should display the right message")

	})
}
