/*
	Tartarus is a library for somewhat securely removing files.
*/
package tartarus

import (
	"crypto/rand"
	"errors"
	"os"
	"path/filepath"
)

// isValid checks that the file pointed to by the path
// A: is writable
// B: is regular
// Returns the file size and error (if any)
func isValid(abspath string) (int64, error) {
	fi, err := os.Stat(abspath)
	if err != nil {
		return 0, err
	}
	mode := fi.Mode()
	if _, err := os.OpenFile(abspath, os.O_RDWR, 0666); err != nil {
		return 0, errors.New("can't read/write to file")
	}
	if !mode.IsRegular() {
		return 0, errors.New("path points to directory, pipe file or physical device")
	}
	return fi.Size(), nil
}

// oWrite overwrites the file at the given path 3 times
// with random data
func oWrite(abspath string) error {
	size, err := isValid(abspath)
	if err != nil {
		return err
	}

	f, err := os.Create(abspath)
	// Shouldn't really happen if we checked permissions and everything
	// But nonetheless
	if err != nil {
		return err
	}

	// Proper defer with Close error handling
	defer func() {
		if cerr := f.Close(); cerr != nil {
			if err == nil {
				err = cerr
			}
		}
	}()
	// Writing in increments of 2048 bytes or less, depending on file size
	// (will considerably slow down with huge files), good idea, to implement
	// a helper function which calls shred with different buffer sizes (depending on allowed buffer size)
	inc := int64(2048)

	for i := int64(0); i < size; i += inc {
		bufsize := inc
		if i+inc > size {
			bufsize = size - i
		}
		buf := make([]byte, bufsize)
		if _, err := rand.Read(buf); err != nil {
			return err
		}
		if _, err := f.Write(buf); err != nil {
			return err
		}
		if err := f.Sync(); err != nil {
			return err
		}
	}
	return nil
}

// Shred rewrites data in a file at a given path 3 times.
// After that it deletes the file
func Shred(path string) error {
	abspath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	err = oWrite(abspath)
	if err != nil {
		return err
	}
	err = os.Remove(abspath)
	return err
}
