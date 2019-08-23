package parser

import (
	"github.com/fabulousduck/sembler/lexer"
	"github.com/fabulousduck/sembler/parser/node"
)

/*
ParseImplied parses an instruction in implied form

*/
func (p *Parser) ParseImplied(line *lexer.Line) *node.Node {
	node := node.NewNode()

	node.Opcode = getOpcodeForImplied(line.CurrentToken().Value)

	return node
}

func getOpcodeForImplied(instruction string) int {
	/*the slices values are represented as follows
		[x,y,0]
	these are modes*/
	opcodeMap := map[string]int{
		"no_operation": 0xAE,
	}

	if value, ok := opcodeMap[instruction]; ok {
		return value
	}
	return 0x0
}
