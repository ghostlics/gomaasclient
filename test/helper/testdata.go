package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Testdata fetches a testdata file.
// The path should be relative to the /test/testdata directory.
func Testdata(path string) (io.ReadCloser, error) {
	path = filepath.Join(RepoBaseDir, "test", "testdata", filepath.FromSlash(path))
	return os.Open(filepath.Clean(path))
}

// TestdataFromJSON fetches a testdata file and decodes it into a type.
// The path should be compatible with Testdata(), and the target should
// be compatible with Golang's json.Unmarshal().
func TestdataFromJSON(path string, target interface{}) error {
	rc, err := Testdata(path)
	if err != nil {
		return err
	}

	defer func() {
		if err := rc.Close(); err != nil {
			fmt.Println("Error when closing:", err)
		}
	}()

	dec := json.NewDecoder(rc)
	dec.DisallowUnknownFields()

	return dec.Decode(target)
}
