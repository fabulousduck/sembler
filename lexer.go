package sembler

import (
	"bytes"
	"fmt"
	"os"

	"github.com/fabulousduck/smol/errors"
)

//Token contains all info about a specific token from syntax
type token struct {
	Value, Type string
	Line, Col   int
}

//Lexer contains all the info needed for the lexer to generate a set of usable tokens
type lexer struct {
	Tokens                                []token
	currentIndex, currentLine, currentCol int
	FileName, Program                     string
}

//NewLexer creates a new instance of a lexer stuct
func newLexer(filename string, program string) *lexer {
	l := new(lexer)
	l.Program = program
	l.FileName = filename
	return l
}

func newToken(line int, col int, value string) *token {
	t := new(token)
	t.Line = line
	t.Col = col
	t.Type = determineType(value)
	t.Value = value
	return t
}

//Lex takes a sourcecode string and transforms it into usable tokens to build an AST with
func (l *lexer) Lex() {

	for l.currentIndex < len(l.Program) {
		currTok := newToken(l.currentLine, l.currentCol, l.currentChar())
		switch currTok.Type {
		case "character":
			currTok.Value = l.peekTypeN("character")
		case "integer":
			currTok.Value = l.peekTypeN("integer")
		case "comment":
			l.readComment()
			l.advance()
			l.currentCol = 0
			continue
		case "dollar":
			fallthrough
		case "comma":
			fallthrough
		case "left_bracket":
			fallthrough
		case "hashtag":
			fallthrough
		case "right_bracket":
			fallthrough
		case "double_dot":
			fallthrough
		case "semicolon":
			l.advance()
		case "undefined_symbol":
			errors.Report(l.currentLine, l.FileName, "undefined symbol used")
			os.Exit(65)
		case "newline":
			l.currentCol = 0
			l.advance()
			continue
		case "ignoreable":
			l.advance()
			continue
		}

		l.Tokens = append(l.Tokens, *currTok)

	}
	l.tagKeywords()
}

func (l *lexer) advance() {
	l.currentCol++
	l.currentIndex++
}

func (l *lexer) readComment() {
	l.currentIndex++
	for t := determineType(l.currentChar()); t != "newline"; t = determineType(l.currentChar()) {
		l.currentIndex++
	}
}

func (l *lexer) peekTypeN(typeName string) string {
	var currentString bytes.Buffer

	for t := determineType(l.currentChar()); t == typeName; t = determineType(l.currentChar()) {
		char := l.currentChar()

		//we do this to avoid index out of range errors
		if l.currentIndex+1 >= len(l.Program) {

			currentString.WriteString(char)
			l.advance()

			return currentString.String()
		}
		currentString.WriteString(char)
		l.advance()
	}

	return currentString.String()
}

func (l *lexer) currentChar() string {
	return string(l.Program[l.currentIndex])
}

func (l *lexer) tagKeywords() {
	for i, token := range l.Tokens {
		if token.Type == "character" {
			l.Tokens[i].Type = getKeyword(&token)
		}
	}
}

//ThrowSemanticError can be used when an error occurs while generating an AST and not at interpret time
func ThrowSemanticError(token *token, expected []string, filename string) {
	errors.Report(
		token.Line,
		filename,
		fmt.Sprintf("expected one of [%s]. got %s",
			errors.ConcatVariables(expected, ", "),
			token.Type),
	)
}
