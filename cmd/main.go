package main

import (
	"flag"
	"log"
	"os"

	"github.com/sivchari/dogen.git/dogen"
)

func main() {
	log.SetFlags(0)
	err := dogen.Run(os.Args, os.Stdout, os.Stderr)
	if err != nil && err != flag.ErrHelp && err != dogen.ShowVersion {
		log.Println(err)
		os.Exit(2)
	}
	log.Println(os.Args[0] + " ʕ◔ϖ◔ʔ ")
	os.Exit(1)
}
