package files_test

import (
	"fmt"
	"testing"

	"github.com/files/pkg/files"
)

type ExpectedListRestult struct {
	files []string
	err   error
}

func TestRunAll(t *testing.T) {
	var tests = []struct {
		name     string
		expected ExpectedListRestult
		testFunc func() ([]string, error)
	}{
		{"TestNonExistingDir", ExpectedListRestult{files: []string{}, err: fmt.Errorf("directory does not exist: ./NonExistingDir")}, func() ([]string, error) { return files.ListFiles("./NonExistingDir") }},
		{"TestEmptyDir", ExpectedListRestult{files: []string{}, err: nil}, func() ([]string, error) { return files.ListFiles("../../test/resources/EmptyDir") }},
		{name: "TestDirWithFiles", expected: ExpectedListRestult{files: []string{"../../test/resources/file1.txt", "../../test/resources/subdir/file2.txt"}, err: nil}, testFunc: func() ([]string, error) { return files.ListFiles("../../test/resources/") }},
	}

	for _, tt := range tests {
		RunTest(t, tt.testFunc, tt.expected)
	}
}

func RunTest(t *testing.T, testedFunc func() ([]string, error), expected ExpectedListRestult) {

	files, err := testedFunc()
	if expected.err != nil {
		if err == nil || err.Error() != expected.err.Error() {
			t.Errorf("Expected error: %v, got: %v", expected.err, err)
			return
		}
	} else if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}
	if len(files) != len(expected.files) {
		t.Errorf("Expected %d files, got %d files", len(expected.files), len(files))
		return
	}
	fileMap := make(map[string]bool)
	for _, f := range files {
		fileMap[f] = true
	}
	for _, ef := range expected.files {
		if !fileMap[ef] {
			t.Errorf("Expected file %s not found in result", ef)
		}

	}
}
