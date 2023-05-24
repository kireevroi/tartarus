# tartarus
A go library that implements a file shredding helper function

## How to run tests
Tests are run with ``go test -coverprofile cover.out -v``

Coverage is checked with ``go tool cover -html=cover.out``

## Test cases
### Implemented test cases:
- Valid files with read/write permissions
- Valid file without permissions
- Valid file path to file that doesn't exist
- Invalid file path
- Path to irregular file
- Path to directory
- Intermediate helper function rewrite test
- Unimplemented:
	- filepath.Abs() failure (e.g. working directory is faulty as well as file path)
	- file.Close() failure (e.g. file doesn't close)
	- os.Create failure (even after checking permissions and everything)
	- file.Write() failure (suddenly can't write or fail to write)
	- file.Sync() failure (suddenly can't sync or fail to sync)

## Use cases
- Semi-secure data deletion (it becomes difficult to recover original content)
- Fulfill company requirements for sensitive data protection

### Advantages:
- Conventient and reusable
- Overwriting multiple times makes it more challenging to recover
- Works for any file types (and sizes if buffer modification helper function is implemented)

### Disadvantages:
- Can't guarantee complete data erasure, physical destruction required
- Time-consuming for large files
- Permanent deletion (backup before running, if there is a slight chance the data may be neede)
- File system specific errors aren't handled (yet)
