package main

import (
	"database/sql"
	"fmt"

	// Import for the side effect on database/sql
	_ "github.com/mattn/go-sqlite3"

	packr "github.com/gobuffalo/packr/v2"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	// migrations := &migrate.MemoryMigrationSource{
	// 	Migrations: []*migrate.Migration{
	// 		&migrate.Migration{
	// 			Id:   "123",
	// 			Up:   []string{"CREATE TABLE people (id int)"},
	// 			Down: []string{"DROP TABLE people"},
	// 		},
	// 	},
	// }

	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./migrations"),
	}

	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		// Handle errors!
	}

	n, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)
	if err != nil {
		// Handle errors!
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
