package main

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type (
	User struct {
		ID      int64
		Name    string
		Balance int64
	}

	Expression struct {
		ID         int64
		Expression string
		UserID     int64
	}
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

	user := &User{
		Name:    "Petr",
		Balance: 200,
	}
	userID, err := insertUser(ctx, db, user)
	if err != nil {
		panic(err)
	}

	expression := &Expression{
		Expression: "2+2",
		UserID:     userID,
	}
	expressionID, err := insertExpression(ctx, db, expression)
	if err != nil {
		panic(err)
	}
	expression.ID = expressionID
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

func insertUser(ctx context.Context, db *sql.DB, user *User) (int64, error) {
	var q = `
		INSERT INTO users (name, balance) values ($1, $2)
	`

	result, err := db.ExecContext(ctx, q, user.Name, user.Balance)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func insertExpression(
	ctx context.Context,
	db *sql.DB,
	expression *Expression,
) (int64, error) {
	var q = `
		INSERT INTO expressions (expression, user_id) values ($1, $2)
	`

	result, err := db.ExecContext(ctx, q, expression.Expression, expression.UserID)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
