package lexer

import (
	"fmt"
	"strings"

	"github.com/fabulousduck/smol/errors"
)

/*
Lexer is a wrapper for all the lines and keeps track of global lexing indexes
*/
type Lexer struct {
	Lines             []Line
	currentLine       int
	FileName, Program string
}

/*
NewLexer returns a new lexer struct upon which lexer functions can be called such as Lex()
*/
func NewLexer(filename string, program string) *Lexer {
	l := new(Lexer)
	l.Program = program
	l.FileName = filename
	return l
}

//Lex takes a sourcecode string and transforms it into usable lines of tokens
func (l *Lexer) Lex() {

	lines := strings.Split(l.Program, "\n")

	for l.currentLine < len(lines) {
		currentLine := NewLine(lines[l.currentLine], l.currentLine)
		currentLine.Lex()
		currentLine.tagKeywords()
		l.Lines = append(l.Lines, *currentLine)
		l.currentLine++

	}
}

//ThrowSemanticError can be used when an error occurs while generating an AST and not at interpret time
func ThrowSemanticError(token *Token, expected []string, filename string) {
	errors.Report(
		token.Line,
		filename,
		fmt.Sprintf("expected one of [%s]. got %s",
			errors.ConcatVariables(expected, ", "),
			token.Type),
	)
}
