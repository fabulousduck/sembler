package lexer

import (
	"bytes"
	"fmt"
	"os"
)

/*
Line is a set of tokens for a single line of assembly

we can do this because assembly is all one line instructions
*/
type Line struct {
	Tokens                  []Token
	Raw                     string
	CurrentIndex, LineIndex int
}

/*
NewLine returns a new line pointer
*/
func NewLine(raw string, lineIndex int) *Line {
	l := new(Line)
	l.Raw = raw
	l.LineIndex = lineIndex
	return l
}

func (l *Line) currentChar() string {
	return string(l.Raw[l.CurrentIndex])
}

/*
Lex turns a line into a set of tokens
*/
func (l *Line) Lex() {
	for l.CurrentIndex < len(l.Raw) {
		currentToken := NewToken(l.LineIndex, l.CurrentIndex, l.currentChar())
		switch currentToken.Type {
		case "character":
			currentToken.Value = l.peekTypeN("character")
		case "integer":
			currentToken.Value = l.peekTypeN("integer")
		case "comment":
			l.readComment()
			l.advance()
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
			//TODO proper errors
			fmt.Printf("fuck off error")
			os.Exit(65)
		case "newline":
			break
		case "ignoreable":
			l.advance()
			continue
		}

		l.Tokens = append(l.Tokens, *currentToken)
	}

}

func (l *Line) peekTypeN(typeName string) string {
	var currentString bytes.Buffer

	for t := determineType(l.currentChar()); t == typeName; t = determineType(l.currentChar()) {
		char := l.currentChar()

		//we do this to avoid index out of range errors
		if l.CurrentIndex+1 >= len(l.Raw) {

			currentString.WriteString(char)
			l.advance()

			return currentString.String()
		}
		currentString.WriteString(char)
		l.advance()
	}

	return currentString.String()
}

func (l *Line) advance() {
	l.CurrentIndex++
}

func (l *Line) readComment() {
	l.CurrentIndex++
	for t := determineType(l.currentChar()); t != "newline"; t = determineType(l.currentChar()) {
		l.CurrentIndex++
	}
}

func (l *Line) tagKeywords() {
	for i, token := range l.Tokens {
		if token.Type == "character" {
			l.Tokens[i].Type = getKeyword(&token)
		}
	}
}
