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
	switch os.Args[1] {
	case "-d", "-dir", "-g", "-gen", "-m", "-model":
		log.Print(os.Args[0] + "!")
	default:
	}
	err := dogen.Run(os.Args, os.Stdout, os.Stderr)
	if err != nil && err != flag.ErrHelp && err != dogen.ShowVersion {
		log.Fatal(err)
	}
	if err == flag.ErrHelp || err == dogen.ShowVersion {
		os.Exit(0)
	}
	log.Print("1/2 complete")
	if err := exec.Command("go", "generate", "./...").Run(); err != nil {
		log.Fatal(err)
	}
	log.Print("2/2 complete")
	log.Println(os.Args[0] + " ʕ◔ϖ◔ʔ ")
}
