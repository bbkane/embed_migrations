// +build ignore

package main

import (
	"path/filepath"
	"log"

    "github.com/bbkane/embed_migrations/data"
    "github.com/bbkane/embed_migrations/projectroot"
	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(data.Migrations, vfsgen.Options{
        Filename: filepath.Join(projectroot.ProjectRootDir(), "data/migrations_vfsdata.go"),
		PackageName:  "data",
		BuildTags:    "!dev",
		VariableName: "Migrations",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
