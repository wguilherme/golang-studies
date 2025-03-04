package main

import (
	"log/slog"
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
	// descomentar para testaar casos acima (exit status 1 ou 0)
	// verificar com echo $?
	// return errors.New("any error")
	return nil
}
