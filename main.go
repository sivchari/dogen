package main

import (
	"dogen/gen"
	"flag"
	"log"
	"os"

	"github.com/kyokomi/emoji"
)

func main() {
	log.SetFlags(0)
	err := gen.Run(os.Args, os.Stdout, os.Stderr)
	if err != nil && err != flag.ErrHelp && err != gen.ShowVersion {
		log.Println(err)
		os.Exit(2)
	}
	emoji := emoji.Sprint(":lollipop:")
	log.Println(os.Args[0] + " ʕ◔ϖ◔ʔ " + emoji)
	os.Exit(0)
}
