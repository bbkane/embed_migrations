// +build dev

package data

import "net/http"

// Migrations contains SQL migrations
var Migrations http.FileSystem = http.Dir("migrations")
