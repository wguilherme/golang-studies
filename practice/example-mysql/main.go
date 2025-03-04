package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		return
	}

	slog.Info("executed successfully")

}

func run() error {

	db, err := sql.Open("mysql", "root:password@/tests")
	if err != nil {
		return err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10) // inclui as conexoes abertas
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		return err
	}

	// Teste de uso do interpolateParams
	// ref: https://github.com/go-sql-driver/mysql?tab=readme-ov-file#interpolateparams
	// e multiStatements
	// ref: https://github.com/go-sql-driver/mysql?tab=readme-ov-file#multistatements

	query := "create table foo (id bigint auto_increment primary key, bar varchar(255));"
	if _, err := db.Exec(query); err != nil {
		return err
	}

	// query = "insert into foo (bar) values ('foobar');"
	// ou como argumento:
	query = "insert into foo (bar) values (?)"
	// uso de placeholder, evita sql injection
	if _, err := db.Exec(query, "foobar"); err != nil {
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
	if err := db.QueryRow(query).Scan(&res.id, &res.bar); err != nil {
		return err
	}

	// verbo em Go "%#+v", que basicamente
	// quer dizer: imprima todas informações do objeto
	// tipo, nome dos campos, pacote que ele pertence, etc...
	fmt.Printf("%#+v\n", res)

	return nil
}
