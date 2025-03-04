package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		return
	}

	slog.Info("executed successfully")

}

func run() error {

	urlExample := "postgres://pg:password@localhost:5432/tests"
	db, err := pgxpool.New(context.Background(), urlExample)
	if err != nil {
		return err
	}

	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	queries := New(db)
	ctx := context.Background()

	authors, err := queries.ListAuthors(context.Background())
	if err != nil {
		return err
	}
	fmt.Println(authors)

	author, err := queries.CreateAuthor(ctx, CreateAuthorParams{
		Name: "any name",
		Bio:  pgtype.Text{String: "any bio", Valid: true},
	})

	if err != nil {
		return err
	}
	fmt.Println(author)

	/* output
	[]
	{1 any name {any bio true}}
	2025/03/04 17:26:10 INFO executed successfully
	*/

	return nil
}
