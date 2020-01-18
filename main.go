package main

import (
	"database/sql"
	"fmt"
	"log"

	// Import for the side effect on database/sql
	_ "github.com/mattn/go-sqlite3"

	packr "github.com/gobuffalo/packr/v2"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./migrations"),
	}

	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	n, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
