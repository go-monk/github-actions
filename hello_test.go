package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestMain_output(t *testing.T) {
	cmd := exec.Command(os.Args[0], "-test.run=^$")
	// Build and run the binary to capture output
	cmd = exec.Command("go", "run", "hello.go")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("failed to run hello.go: %v", err)
	}
	got := strings.TrimSpace(string(out))
	want := "Hello world!"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
