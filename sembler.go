package sembler

import (
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"

	"github.com/fabulousduck/sembler/lexer"
)

/*
Sembler is a mounter struct for all functions in the sembler lib
*/
type Sembler struct {
}

/*
NewSembler returns a new instance of a sembler.

All functions that a user might need are mounted on this
*/
func NewSembler() *Sembler {
	return new(Sembler)
}

/*
Compile takes a file and generates a binary with that name containing the assembled 6502 assembly
*/
func (s *Sembler) Compile(filename string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	l := lexer.NewLexer(filename, string(file))
	l.Lex()

	spew.Dump(l)
}
