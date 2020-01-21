// +build dev

package data

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
	"path/filepath"
)

// https://github.com/gobuffalo/here/blob/f233917fd86cadcfa7396b89ffda281b7268fdaa/here.go#L29
func run(n string, args ...string) ([]byte, error) {
	c := exec.Command(n, args...)

	bb := &bytes.Buffer{}
	ebb := &bytes.Buffer{}
	c.Stdout = bb
	c.Stderr = ebb
	err := c.Run()
	if err != nil {
		return nil, fmt.Errorf("%s: %s", err, ebb)
	}

	return bb.Bytes(), nil
}

func projectRootDir() string {
	b, err := run("go", "env", "GOMOD")
	if err != nil {
		// go env shoudl return an empty string if modules not used
		panic(fmt.Sprintf("This should never happen: %v", err))
	}
	return filepath.Dir(string(b))
}

// Migrations contains SQL migrations
var Migrations http.FileSystem = http.Dir(filepath.Join(projectRootDir(), "data/migrations"))
