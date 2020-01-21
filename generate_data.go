// +build ignore

package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"log"

    "github.com/bbkane/embed_migrations/data"
	"github.com/shurcooL/vfsgen"
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


func main() {
	err := vfsgen.Generate(data.Migrations, vfsgen.Options{
        Filename: filepath.Join(projectRootDir(), "data/migrations_vfsdata.go"),
		PackageName:  "data",
		BuildTags:    "!dev",
		VariableName: "Migrations",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
