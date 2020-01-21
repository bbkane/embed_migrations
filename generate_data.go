// +build ignore

package main

import (
	"log"
	"path/filepath"

	"github.com/bbkane/embed_migrations/clienthtml"
	"github.com/bbkane/embed_migrations/migrations"
	"github.com/bbkane/embed_migrations/projectroot"
	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(migrations.Migrations, vfsgen.Options{
		Filename:     filepath.Join(projectroot.ProjectRootDir(), "migrations/migrations_vfsdata.go"),
		PackageName:  "migrations",
		BuildTags:    "!dev",
		VariableName: "Migrations",
	})
	if err != nil {
		log.Fatalln(err)
	}
	err = vfsgen.Generate(clienthtml.ClientHtml, vfsgen.Options{
		Filename:     filepath.Join(projectroot.ProjectRootDir(), "clienthtml/clienthtml_vfsdata.go"),
		PackageName:  "clienthtml",
		BuildTags:    "!dev",
		VariableName: "ClientHtml",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
