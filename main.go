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
	log.Print(os.Args[0] + "!")
	if err := dogen.Run(os.Args, os.Stdout, os.Stderr); err != nil && err != flag.ErrHelp && err != dogen.ShowVersion {
		log.Fatal(err)
	}
	log.Print("1/2 complete")
	if err := exec.Command("go", "generate", "./...").Run(); err != nil {
		log.Fatal(err)
	}
	log.Print("2/2 complete")
	log.Println(os.Args[0] + " ʕ◔ϖ◔ʔ ")
}
