package cli

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_test_helper(t *testing.T) {
	t.Run("Should mock cli call with arguments", func(t *testing.T) {
		MockCLICall("app command")
		assert.Equal(t, "app", os.Args[0], "Should find the right app called")
		assert.Equal(t, "command", os.Args[1], "Should find the right command passed")
	})

	t.Run("Should support go flags", func(t *testing.T) {
		MockCLICall("app -t expectedValue")
		actualValue := flag.String("t", "", "example")
		flag.Parse()
		assert.Equal(t, "app", os.Args[0], "Should find the right app called")
		assert.Equal(t, "expectedValue", *actualValue, "Should be able to read the flag correctly")
	})

	t.Run("Should support go flags and commands", func(t *testing.T) {
		MockCLICall("app command -t expectedValue")
		cmd := flag.NewFlagSet("command", flag.ExitOnError)
		actualValue := cmd.String("t", "", "example")
		cmd.Parse(os.Args[2:])
		assert.Equal(t, "app", os.Args[0], "Should find the right app called")
		assert.Equal(t, "command", os.Args[1], "Should find the right command passed")
		assert.Equal(t, "expectedValue", *actualValue, "Should be able to read the flag correctly")
	})
}
