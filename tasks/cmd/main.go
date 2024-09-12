package cmd

import (
	"flag"
	"os"

	"github.com/mulatinho/golabs/tasks/core"
)

func Start() {
	useDebug := flag.Bool("debug", false, "display debug messages in output")

	if *useDebug {
		os.Setenv("DEBUG", "true")
	}

	core.Start()
}
