package main

import (
	"fmt"
	"log"
	"os"

	"github.com/escaletech/tog-cli/internal/color"
	"github.com/escaletech/tog-cli/internal/command"
	"github.com/escaletech/tog-cli/internal/config"
)

var Version = ""

func main() {
	cstore, err := config.New()
	if err != nil {
		log.Fatal("failed to initialize config store: ", err)
	}

	if cmd, err := command.Execute(Version, cstore); err != nil {
		fmt.Printf("%v %v\n\n", color.Error("Error:"), err)
		cmd.Usage()
		os.Exit(1)
	}
}
