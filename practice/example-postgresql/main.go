package main

import (
	"context"
	"fmt"
	"log/slog"

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

	// exemplo com ENV:
	// conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	urlExample := "postgres://pg:password@localhost:5432/tests"
	// dessa forma o DB é apenas uma conexão e não um pool
	/*
		db, err := pgx.Connect(context.Background(), urlExample)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
		defer db.Close(context.Background())
	*/

	// exemplo com pool de conexão:

	db, err := pgxpool.New(context.Background(), urlExample)
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	query := "create table foo (id bigserial primary key, bar varchar(255));"

	if _, err := db.Exec(context.Background(), query); err != nil {
		return err
	}

	// query = "insert into foo (bar) values ('foobar');"
	// ou como argumento:
	// neste caso, o placeholder muda com postgres, ao invés de ?
	// usamos $<n> de acordo com a ordem dos placeholders
	query = "insert into foo (bar) values ($1)"
	// uso de placeholder, evita sql injection
	if _, err := db.Exec(context.Background(), query, "foobar"); err != nil {
		return err
	}

	query = "select * from foo limit 1"

	type foobar struct {
		id  int64
		bar string
	}

	var res foobar
	// sempre que fazemos scan, temos que passar um ponteiro
	// pois o scan deve ser capaz de modificar o valor de res
	// é operação de mutação
	if err := db.QueryRow(context.Background(), query).Scan(&res.id, &res.bar); err != nil {
		return err
	}

	// verbo em Go "%#+v", que basicamente
	// quer dizer: imprima todas informações do objeto
	// tipo, nome dos campos, pacote que ele pertence, etc...
	fmt.Printf("%#+v\n", res)

	return nil
}
