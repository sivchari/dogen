package main

import (
	"flag"
	"log"
	"os"
	"os/exec"

	"github.com/sivchari/dogen/dogen"
)

func main() {
	log.SetFlags(0)
	if err := dogen.Run(os.Args, os.Stdout, os.Stderr); err != nil && err != flag.ErrHelp && err != dogen.ShowVersion {
		log.Fatal(err)
	}
	if err := exec.Command("go", "generate", "./...").Run(); err != nil {
		log.Fatal(err)
	}
	log.Println(os.Args[0] + " ʕ◔ϖ◔ʔ ")
}
