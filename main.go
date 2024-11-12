package main

import (
	"log/slog"
	"os"

	"github.com/ei-sugimoto/datapuppy/cmd"
	"github.com/ei-sugimoto/datapuppy/internal/infra"
)

func main() {
	slog.Info("Starting the application...")

	db, err := infra.Connect()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	err = infra.Migrate(db)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	cmd.Serve()
}
