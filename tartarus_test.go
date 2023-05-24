package tartarus

import (
	"os"
	"testing"
)

func TestShredDirectory(t *testing.T) {
	// Run the Shred function with an invalid file path
	err := Shred("testfiles")
	if err == nil {
		t.Error("Expected error for irregular file")
	}
}
func TestShredInvalidFile(t *testing.T) {
	// Run the Shred function with an invalid file path
	err := Shred("")
	if err == nil {
		t.Error("Expected error for invalid file path")
	}
}
func TestShredNonExistentFile(t *testing.T) {
	// Run the Shred function with a non-existent file
	err := Shred("ne.txt")
	if err == nil {
		t.Error("Expected error for non existing file")
	}
}

func TestMain(m *testing.M) {

	e := m.Run()

	os.Exit(e)
}