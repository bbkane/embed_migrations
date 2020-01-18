// +build ignore

package main

import (
	"log"

    "github.com/bbkane/hello_vfsgen/data"
	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(data.Assets, vfsgen.Options{
        Filename: "data/assets_vfsdata.go",
		PackageName:  "data",
		BuildTags:    "!dev",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
