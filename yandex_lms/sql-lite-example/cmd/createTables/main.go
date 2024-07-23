package main

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx := context.TODO()

	db, err := sql.Open("sqlite3", "store.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	if err = createTables(ctx, db); err != nil {
		panic(err)
	}
}

func createTables(ctx context.Context, db *sql.DB) error {
	const (
		usersTable = `
			CREATE TABLE IF NOT EXISTS users(
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				name TEXT,
				balance INTEGER NOT NULL CHECK(balance >= 0)
			);`
		expressionsTable = `
			CREATE TABLE IF NOT EXISTS expressions(
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				expression TEXT NOT NULL,
				user_id INTEGER NOT NULL,
				FOREIGN KEY (user_id) REFERENCES expressions (id)
			);`
	)

	if _, err := db.ExecContext(ctx, usersTable); err != nil {
		return err
	}
	if _, err := db.ExecContext(ctx, expressionsTable); err != nil {
		return err
	}

	return nil
}
