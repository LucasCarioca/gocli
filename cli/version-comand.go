package cli

import (
	"fmt"
	"io"
	"os"
)

var (
	customOut io.Writer = os.Stdout
)

//VersionCommand used to retrieve the current version of the terraform cli
type VersionCommand struct {
	Version string
}

//Run executes the command
func (c *VersionCommand) Run(_ AppInterface) error {
	fmt.Fprintf(customOut, "Current version: v%s\n", c.Version)
	return nil
}
