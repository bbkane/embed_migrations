package projectroot

import (
	"bytes"
	"fmt"
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

// ProjectRoot returns the root directory if go modules is on, else an empty
// string It can panic if 'go env GOMOD' fails, which I don't think ever
// happens
func ProjectRootDir() string {
	b, err := run("go", "env", "GOMOD")
	if err != nil {
		// go env shoudl return an empty string if modules not used
		panic(fmt.Sprintf("This should never happen: %v", err))
	}
	return filepath.Dir(string(b))
}

