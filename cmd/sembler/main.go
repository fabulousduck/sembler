package main

import (
	"flag"

	"github.com/fabulousduck/sembler"
)

func main() {
	s := sembler.NewSembler()

	filenamePtr := flag.String("file", "", "input file for the interpreter")
	flag.Parse()

	if *filenamePtr != "" {
		s.Compile(*filenamePtr)
	}
}
