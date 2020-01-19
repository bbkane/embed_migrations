// +build ignore

package main

import (
	"log"

    "github.com/bbkane/embed_migrations/data"
	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(data.Migrations, vfsgen.Options{
        Filename: "data/migrations_vfsdata.go",
		PackageName:  "data",
		BuildTags:    "!dev",
		VariableName: "Migrations",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
