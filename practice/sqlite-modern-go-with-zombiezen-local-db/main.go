package main

import (
	"fmt"
	"log/slog"

	"zombiezen.com/go/sqlite"
	"zombiezen.com/go/sqlite/sqlitex"
)

func main() {

	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)

		// aqui temos 2 opções:
		// return, vai sair com exit status 0
		// ou os.Exit(1), vai sair com exit status 1
		return
	}
	slog.Info("all systems offline")
}

func run() error {
	// ...

	// Open an in-memory database.
	conn, err := sqlite.OpenConn("./test.db", sqlite.OpenReadWrite|sqlite.OpenCreate)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Execute a query.
	err = sqlitex.ExecuteTransient(conn, "SELECT 'hello, world';", &sqlitex.ExecOptions{
		ResultFunc: func(stmt *sqlite.Stmt) error {
			fmt.Println(stmt.ColumnText(0))
			return nil
		},
	})
	if err != nil {
		return err
	}

	return nil
}
