package pathsanitize

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

const TEST_FILE = "/test-lcov.info"

func TestPathSanitize(t *testing.T) {
	tmpDir := t.TempDir()
	copyTestFile(tmpDir)

	trimWord := "projects"
	pscmd := CreateCommand()
	buffer := bytes.NewBufferString("")
	pscmd.SetOut(buffer)
	pscmd.SetArgs([]string{"-t", trimWord, "--lcov", tmpDir + TEST_FILE})
	err := pscmd.Execute()
	if err != nil {
		t.Fatal(err)
		return
	}
	_, err = io.ReadAll(buffer)
	if err != nil {
		t.Fatal(err)
		return
	}

	testFile, _ := os.ReadFile(tmpDir + TEST_FILE)
	assert.NotContains(t, string(testFile), "\\")
	lines := strings.Split(string(testFile), "\n")
	for _, line := range lines {
		if strings.Contains(line, trimWord) {
			assert.Equal(t, strings.Index(line, trimWord), 3)
		}
	}
}

func copyTestFile(tmpDir string) {
	input, err := os.ReadFile("../../test_resources/lcov.info")
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = os.WriteFile(tmpDir+TEST_FILE, input, 0600)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
