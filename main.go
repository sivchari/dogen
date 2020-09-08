package main

import (
	"dogen/gen"
	"flag"
	"os"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	err := gen.Run(os.Args, os.Stdout, os.Stderr)
	if err != nil && err != flag.ErrHelp && err != gen.ShowVersion {
		logger.Info("dogen: logging")
		os.Exit(2)
	}
}
