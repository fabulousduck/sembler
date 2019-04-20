package lexer

/*
Token contains data about a single element of grammer
*/
type Token struct {
	Value, Type string
	Line, Col   int
}

/*
NewToken returns a new token at a given index
*/
func NewToken(line int, col int, value string) *Token {
	t := new(Token)
	t.Line = line
	t.Col = col
	t.Type = determineType(value)
	t.Value = value
	return t
}
