//go:generate go run -tags=dev generate_assets.go

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

    "github.com/bbkane/hello_vfsgen/data"
)

// similar to https://github.com/rakyll/statik/blob/74b078749dc19d4fead4faf52bf8f2b88b9e7a74/fs/fs_test.go#L257
func testReadDir(fs http.FileSystem) {

	file, err := fs.Open("/")
	if err != nil {
		log.Fatal(err)
	}

	// This returns an empty list
	files, err := file.Readdir(0)
	if err != nil {
		log.Fatal(err)
	}

	// https://github.com/rakyll/statik/blob/74b078749dc19d4fead4faf52bf8f2b88b9e7a74/fs/fs.go#L164
    // says directory listing is disabled. It seems to work in tests

    // prints an empty list
	fmt.Printf("readdir: %v\n", files)
}

func testReadFile(httpFS http.FileSystem) {
	// Access individual files by their paths.
	r, err := httpFS.Open("/migrations/2019_10_12.1.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	contents, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

    // indicates that there is a file
	fmt.Println(string(contents))
}

func main() {

    // check that at least one file made it
	testReadFile(data.Assets)

    // yet still no files are shown in the directory
	testReadDir(data.Assets)
}
