package tartarus

import (
	"path/filepath"
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
// func TestShredInvalidPath(t *testing.T) {
// 	// Run the Shred function with a failure in Abs path
// 	dir, _ := os.MkdirTemp("", "") 
//   os.Chdir(dir)                  
//   os.RemoveAll(dir)              
// 	err := Shred("")
// 	orig, _ := os.Getwd()
//   t.Cleanup(func() {
//       os.Chdir(orig)
//   })
// 	if err == nil {
// 		t.Error("Expected error for invalid path")
// 	}
// }


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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Shred(tt.input)
			if got != tt.ans {
				t.Errorf("got an error: %v", got)
			}
			path, _ := filepath.Abs(tt.input)
			if _, err := os.Stat(path); err == nil {
				t.Error("file wasn't deleted")
			}
		})
	}
}

func TestMain(m *testing.M) {

	e := m.Run()

	os.Exit(e)
}