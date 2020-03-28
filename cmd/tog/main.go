package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/escaletech/tog-cli/internal/color"
	"github.com/escaletech/tog-cli/internal/command"
	"github.com/escaletech/tog-cli/internal/config"
)

var Version = "dev"
var BuildDate = time.Now().Format("2006-01-02")

func main() {
	cstore, err := config.New()
	if err != nil {
		log.Fatal("failed to initialize config store: ", err)
	}

	command.SetMetadata(Version, BuildDate)

	if cmd, err := command.Execute(cstore); err != nil {
		fmt.Printf("%v %v\n\n", color.Error("Error:"), err)
		cmd.Usage()
		os.Exit(1)
	}
}
