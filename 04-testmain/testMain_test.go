package testMain

import (
	"os"
	"testing"
)

// TestA demonstrates a simple test case.
func TestA(t *testing.T) {
	t.Log("Test A run")
}

// TestB demonstrates another simple test case.
func TestB(t *testing.T) {
	t.Log("Test B run")
}

// TestMain provides setup and teardown for the test suite.
func TestMain(m *testing.M) {
	// TODO: add any required setup code here.
	exitVal := m.Run()
	if exitVal == 0 {
		// TODO: add any required teardown code here.
	}
	os.Exit(exitVal)
}
