package test

import (
	"bytes"
	"hello-world/pkg/greetings"
	"io"
	"os"
	"testing"
)

func TestRunAll(t *testing.T) {
	var tests = []struct {
		name     string
		expected string
		testFunc func()
	}{
		{"TestSayHello", "Hello world!", func() { greetings.SayHello() }},
		{"TestSayHelloWithCustomName", "Hello tester!", func() { greetings.SayHelloCustomized("tester", false) }},
		{"TestSayHelloWithCustomNameAndNewLine", "Hello tester!\n", func() { greetings.SayHelloCustomized("tester", true) }},
	}

	for _, tt := range tests {
		RunTest(t, tt.testFunc, tt.expected)
	}
}

func RunTest(t *testing.T, testedFunc func(), expected string) {
	// Save the original stdout.
	originalStdout := os.Stdout
	defer func() { os.Stdout = originalStdout }() // Restore at the end.

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	os.Stdout = w // Redirect stdout to the pipe's write end.

	testedFunc()

	w.Close() // Close the write end to signal completion.

	// Read the captured output.
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Errorf("failed to read output: %v", err)
	}

	// Assert on the result.
	got := buf.String()

	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}
