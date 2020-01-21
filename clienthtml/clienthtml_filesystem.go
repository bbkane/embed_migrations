// +build dev

package clienthtml

import (
	"net/http"
	"path/filepath"

    "github.com/bbkane/embed_migrations/projectroot"
)

// Migrations contains SQL migrations
var ClientHtml http.FileSystem = http.Dir(filepath.Join(projectroot.ProjectRootDir(), "clienthtml/clienthtml"))
