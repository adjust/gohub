package main

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func TestExecuteShellLogsError(t *testing.T) {
	buf := setupLog()
	executeShell("ls non-existent")

	if !strings.Contains(buf.String(), "Error:") {
		t.Error("Output does not contain error message.")
	}
}

func TestExecuteShellLogsOutput(t *testing.T) {
	buf := setupLog()
	executeShell("ls")

	if !strings.Contains(buf.String(), "Shell output was") {
		t.Error("Command execution failed.")
	}
}

func setupLog() *bytes.Buffer {
	buf := new(bytes.Buffer)
	log.SetOutput(buf)

	return buf
}
