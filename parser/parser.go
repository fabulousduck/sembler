package parser

import (
	"fmt"
	"os"

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
	Labels      []*Label
	CurrentByte int //we use currentByte for when we need to keep track of where we found a label so we can later jump back to it
}

/*
Label is a struct containing info about labels in the code

Labels can be used to define memory adresses of the line so instructions
line BNE or JSR can jump to it
*/
type Label struct {
	Name string
	Pos  int
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
func (p *Parser) Parse(lines *[]lexer.Line) {
	for _, line := range *lines {
		mode := mode.GetInstructionMode(&line)

		//we have to reset the cursor to 0 on the line
		line.CurrentIndex = 0

		nodes := p.ParseLine(&line, mode)
		p.ParsedNodes = append(p.ParsedNodes, nodes)
	}
}

func (p *Parser) createLabel(line *lexer.Line) {
	label := new(Label)
	label.Name = line.Tokens[0].Value
	label.Pos = p.CurrentByte
	p.Labels = append(p.Labels, label)
	line.Advance()
}

func (p *Parser) getLabelByName(name string) *Label {
	for _, label := range p.Labels {
		if label.Name == name {
			return label
		}
	}

	fmt.Printf("attempt to reference undefined label \"%s\"", name)
	os.Exit(65)
	return nil
}

/*
ParseLine parses an MBI line into an opcode node
*/
func (p *Parser) ParseLine(line *lexer.Line, mode *mode.Mode) *node.Node {
	//check if the line defines a label
	if lexer.GetKeyword(&line.Tokens[0]) == "string" {
		p.createLabel(line)
	}

	switch mode.Name {
	case "implied":
		p.CurrentByte += 2
		return p.ParseImplied(line)
	case "immidiate":
		p.CurrentByte += 2
		return p.ParseImmidiate(line)
	case "indirect":
		p.CurrentByte += 2
		return p.ParseIndirect(line, mode.Variable)
	case "absolute":
		p.CurrentByte += 4
		return p.ParseAbsolute(line, mode.Variable)
	case "zeroPage":
		p.CurrentByte += 2
		return p.ParseZeroPage(line, mode.Variable)
	}

	return node.NewNode()
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
