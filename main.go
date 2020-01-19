//go:generate go run -tags=dev generate_migrations.go

package main

import (
	"database/sql"
	"fmt"
	"log"

	// Import for the side effect on database/sql
	_ "github.com/mattn/go-sqlite3"

	migrate "github.com/rubenv/sql-migrate"

    "github.com/bbkane/embed_migrations/data"
)

func main() {
	migrationSource := &migrate.HttpFileSystemMigrationSource{
		FileSystem: data.Migrations,
	}

    migrationList, err := migrationSource.FindMigrations()
	if err != nil {
		log.Fatal(err)
	}
    fmt.Printf("migrationList: %v\n", migrationList)

	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	n, err := migrate.Exec(db, "sqlite3", migrationSource, migrate.Up)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
