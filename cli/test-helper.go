package cli

import (
	"os"
	"strings"
)

//MockCLICall mock arguments passed to the CLI for tests
func MockCLICall(cmd string) {
	os.Args = strings.Split(cmd, " ")
}
