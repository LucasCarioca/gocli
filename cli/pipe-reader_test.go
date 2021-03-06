package cli

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func Test_pipe_reader(t *testing.T) {
	temporaryFile, err := ioutil.TempFile("", "example")
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	oldStdin := os.Stdin
	os.Stdin = temporaryFile
	defer func() {
		os.Remove(temporaryFile.Name())
		os.Stdin = oldStdin
	}()
	setPipeContent := func(content string) {
		if _, err := temporaryFile.Write([]byte(content)); err != nil {
			assert.FailNow(t, err.Error())
		}
		if _, err := temporaryFile.Seek(0, 0); err != nil {
			assert.FailNow(t, err.Error())
		}
	}

	t.Run("Should support reading piped input", func(t *testing.T) {
		setPipeContent("expectedPipedData")
		s, err := ReadPipe()
		assert.Nil(t, err, "Should not throw an error")
		assert.Equal(t, "expectedPipedData", s, "Should read the right piped data")
	})

}

func Test_pipe_reader_errors(t *testing.T) {
	t.Run("Should return an error if not piped input provided", func(t *testing.T) {
		s, err := ReadPipe()
		expectedError := "the command is intended to work with pipes"
		assert.NotNil(t, err, "Should not throw an error")
		assert.Equal(t, expectedError, err.Error(), "Should return the proper error")
		assert.Equal(t, "", s, "Should return an empty string")
	})
}

func Test_pipe_reader_file_errors(t *testing.T) {

	os.Stdin = nil
	defer func() {
		oldStdin := os.Stdin
		os.Stdin = oldStdin
	}()

	t.Run("Should support reading piped input", func(t *testing.T) {
		s, err := ReadPipe()
		expectedError := "invalid argument"
		assert.NotNil(t, err, "Should not throw an error")
		assert.Equal(t, expectedError, err.Error(), "Should return the proper error")
		assert.Equal(t, "", s, "Should return an empty string")
	})

}
