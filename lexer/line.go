package lexer

import (
	"bytes"
	"fmt"
	"os"

	"github.com/fabulousduck/proto/src/types"
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
			l.Tokens = append(l.Tokens, *currentToken)
			return
		case "left_paren":
			fallthrough
		case "right_paren":
			fallthrough
		case "dollar":
			fallthrough
		case "comma":
			fallthrough
		case "plus":
			fallthrough
		case "dash":
			fallthrough
		case "slash":
			fallthrough
		case "star":
			fallthrough
		case "hashtag":
			fallthrough
		case "left_bracket":
			fallthrough
		case "right_bracket":
			fallthrough
		case "double_dot":
			l.Advance()
		case "undefined_symbol":
			//TODO proper errors
			fmt.Printf("fuck off error")
			os.Exit(65)
		case "newline":
			break
		case "ignoreable":
			l.Advance()
			continue
		default:
			fmt.Printf("unknown character %s", currentToken.Type)
		}

		l.Tokens = append(l.Tokens, *currentToken)
	}
	return
}

func (l *Line) peekTypeN(typeName string) string {
	var currentString bytes.Buffer

	for t := determineType(l.currentChar()); t == typeName; t = determineType(l.currentChar()) {
		char := l.currentChar()

		//we do this to avoid index out of range errors
		if l.CurrentIndex+1 >= len(l.Raw) {

			currentString.WriteString(char)
			l.Advance()

			return currentString.String()
		}
		currentString.WriteString(char)
		l.Advance()
	}

	return currentString.String()
}

/*
Advance moves the currentIndex of the line one forward
*/
func (l *Line) Advance() {
	l.CurrentIndex++
}

/*
Eol returns if there are no more tokens to be read
*/
func (l *Line) Eol() bool {
	return l.CurrentIndex+1 == len(l.Tokens)
}

/*
HasSingleChar checks all tokens for an occurance of char
*/
func (l *Line) HasSingleChar(char string) bool {
	for _, token := range l.Tokens {
		if token.Value == char {
			return true
		}
	}
	return false
}

/*
Expect checks if the NEXT token is of a given set of types.
If not, it will throw a syntax error
*/
func (l *Line) Expect(expectedValues []string) {
	nextToken := l.NextToken()
	if !types.Contains(nextToken.Type, expectedValues) {
		//TODO proper errors
		fmt.Printf("\nsyntax error: unexpected %s\nexpected one of : ", nextToken.Value)
		for _, val := range expectedValues {
			fmt.Printf("%s, ", val)
		}
		fmt.Printf("\nin line: %s\n\n", l.Raw)
		// ThrowSemanticError(&nextToken, expectedValues, p.Filename)
		os.Exit(65)
	}
}

/*
ExpectSequence allows you to expect a sequence of types
*/
func (l *Line) ExpectSequence(expectedValues [][]string) {
	for _, expectedSubValues := range expectedValues {
		l.Expect(expectedSubValues)
		l.Advance()
	}
}

/*
CurrentToken returns the token at the currentIndex
*/
func (l *Line) CurrentToken() Token {
	return l.Tokens[l.CurrentIndex]
}

/*
NextToken returns the token at the next index above currentIndex
*/
func (l *Line) NextToken() Token {
	return l.Tokens[l.CurrentIndex+1]
}

func (l *Line) tagKeywords() {
	for i, token := range l.Tokens {
		if token.Type == "character" {
			l.Tokens[i].Type = GetKeyword(&token)
		}
	}
}
