package tartarus

import (
	"bytes"
	"os"
	"path/filepath"
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
func TestShredInvalidPermissions(t *testing.T) {
	err := Shred("tests/nopermission.txt")
	if err == nil {
		t.Error("Expected error for no permissions")
	}
}
func TestShredPipe(t *testing.T) {
	// Run the Shred function with a non-existent file
	err := Shred(os.DevNull)
	if err == nil {
		t.Error("Expected error for irregular file")
	}
}
func TestShredInvalidPath(t *testing.T) {
	// Run the Shred function with a failure in Abs path
	orig, _ := os.Getwd()
	dir, _ := os.MkdirTemp("", "") 
  os.Chdir(dir)                  
  os.RemoveAll(dir)              
	err := Shred("")
  t.Cleanup(func() {
      os.Chdir(orig)
  })
	if err == nil {
		t.Error("Expected error for invalid path")
	}
}


func TestShredValid(t *testing.T) {
	var tests = []struct {
		name string
		input string
		ans error
	}{
		{
			"Simple text file",
			"tests/randomfile.txt",
			nil,
		},
		{
			"Big file test",
			"tests/bigfile",
			nil,
		},
		{
			"Unicode text file",
			"tests/unicode",
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Shred(tt.input)
			if got != tt.ans {
				t.Errorf("got an error: %v", got)
			}
			path, _ := filepath.Abs(tt.input)
			if _, err := os.Stat(path); !os.IsNotExist(err) {
				t.Error("file wasn't deleted")
			}
		})
	}
}

func TestOWrite(t *testing.T) {
	file, err := os.Open("./tests/testscramble")
	if err != nil {
		t.Error("couldn't open initial file")
	}

	buffer := make([]byte, 1024)
	n, _ := file.Read(buffer)
	file.Close()

	got := oWrite("./tests/testscramble")
	if got != nil {
		t.Errorf("got an error: %v", got)
	}

	file2, err := os.Open("./tests/testscramble")
	if err != nil {
		t.Error("couldn't open scrambled file")
	}

	buffer_new := make([]byte, 1024)

	n_new, _ := file2.Read(buffer_new)
	file2.Close()
	
	if n_new != n {
		t.Error("file scrambled incorrectly (size problem)")
	}
	if bytes.Equal(buffer, buffer_new) {
		t.Error("file didn't scramble at all")
	}
}

func TestMain(m *testing.M) {

	e := m.Run()

	os.Exit(e)
}