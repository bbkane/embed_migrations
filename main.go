//go:generate go run -tags=dev generate_data.go

package main

import (
	"database/sql"
	"embed"
	"fmt"
	"os"

	migrate "github.com/rubenv/sql-migrate"
	"gopkg.in/gorp.v1"
	_ "modernc.org/sqlite"
)

//go:embed migrations
var migrations embed.FS

func run() error {
	// HACK to use modernc/sqlite instead of mattn's CGO version
	migrate.MigrationDialects["sqlite"] = gorp.SqliteDialect{}

	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: migrations,
		Root:       "migrations",
	}

	dbPath := "tmp.db"

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return fmt.Errorf("open error: %s: %w", dbPath, err)
	}

	n, err := migrate.Exec(db, "sqlite", migrations, migrate.Up)
	if err != nil {

		return fmt.Errorf("migrate error: %s: %w", dbPath, err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
