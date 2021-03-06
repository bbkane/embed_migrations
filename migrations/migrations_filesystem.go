// +build dev

package migrations

import (
	"net/http"
	"path/filepath"

    "github.com/bbkane/embed_migrations/projectroot"
)

// Migrations contains SQL migrations
var Migrations http.FileSystem = http.Dir(filepath.Join(projectroot.ProjectRootDir(), "migrations/migrations"))
