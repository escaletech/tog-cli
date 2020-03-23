package main

import (
	"log"
	"os"

	"github.com/escaletech/tog-cli/internal/command"
	"github.com/escaletech/tog-cli/internal/config"
)

var Version = ""

func main() {
	cstore, err := config.New()
	if err != nil {
		log.Fatal("failed to initialize config store: ", err)
	}

	if err := command.Execute(Version, cstore); err != nil {
		os.Exit(1)
	}
}
