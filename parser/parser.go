package parser

import (
	"strings"

	"github.com/fabulousduck/sembler/lexer"
	"github.com/fabulousduck/sembler/parser/mode"
	"github.com/fabulousduck/sembler/parser/node"
)

/*
Parser struct
structure on which all paring functions can be called
*/
type Parser struct {
	ParsedNodes []*node.Node
}

/*
NewParser returns a new parser structure pointer
*/
func NewParser() *Parser {
	return new(Parser)
}

/*
Parse takes a set of lexed lines and turns them into nodes
these nodes can then be made into opcodes
*/
func (p *Parser) Parse(lines []lexer.Line) {
	for _, line := range lines {
		nodes := ParseMBI(&line, GetInstructionMode(&line))
		p.ParsedNodes = append(p.ParsedNodes, nodes)
	}
}

/*
ParseMBI parses an MBI line into an opcode node
*/
func ParseMBI(line *lexer.Line, mode *mode.Mode) *node.Node {
	switch mode.Name {
	case "implied":
		return ParseImplied(line)
	case "immidiate":
		return ParseImmidiate(line)
	case "indirect":
		return ParseIndirect(line, mode.Variable)
	case "absolute":
		return ParseAbsolute(line, mode.Variable)
	case "zeroPage":
		return ParseZeroPage(line, mode.Variable)
	}

	return node.NewNode()
}

/*
GetInstructionMode gets the mode in which the line was written
*/
func GetInstructionMode(line *lexer.Line) *mode.Mode {
	mode := mode.NewMode()

	//1 is the index at which mode is defined

	modeIndentifierChar := line.Tokens[1].Type
	operationValue := line.Tokens[2].Value

	//check if its an implied instruction like BRK
	if len(line.Tokens) == 1 {
		mode.Name = "implied"
		mode.Variable = ""
		return mode
	}

	//final character for non direct operations is the last one
	XYNonDirectLocation := strings.ToLower(line.Tokens[len(line.Tokens)-1].Value)

	//check for x or y variables
	if XYNonDirectLocation == "x" || XYNonDirectLocation == ")" {
		mode.Variable = "x"
	}

	if XYNonDirectLocation == "y" {
		mode.Variable = "y"
	}

	//only immidiate mode starts with a #
	if modeIndentifierChar == "hashtag" {
		mode.Name = "immidiate"
		mode.Variable = ""
		return mode
	}

	//if it is encapsulated, we can assume it is an indirect operation
	if modeIndentifierChar == "left_paren" {
		mode.Name = "indirect"
		return mode
	}

	//if the value of the operation is 4 characters long,
	//it is assumed it is an absolute operation
	if len(operationValue) == 4 {
		mode.Name = "absolute"
		return mode
	}

	mode.Name = "zeroPage"
	return mode

}

/*
FindInt looks for an integer value in a parsed line
2nd return value indicatest whether it has been found or not
1 means found
0 means not found
*/
func FindInt(l *lexer.Line) (lexer.Token, int) {
	for _, value := range l.Tokens {
		if value.Type == "integer" {
			return value, 1
		}
	}

	return l.Tokens[0], 0
}

func (p *Parser) validateSyntax() {

}
