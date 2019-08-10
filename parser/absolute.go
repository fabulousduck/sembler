package parser

import (
	"github.com/fabulousduck/sembler/lexer"
	"github.com/fabulousduck/sembler/parser/byte"
	"github.com/fabulousduck/sembler/parser/node"
)

/*
ParseAbsolute parses an instruction in absolute form

*/
func ParseAbsolute(line *lexer.Line, mode string) *node.Node {
	node := node.NewNode()

	node.Instruction = "load_accumelator"

	line.Expect([]string{"dollar"})
	line.Advance()

	line.Expect([]string{"integer"})
	line.Advance()
	integerValueString := line.CurrentToken().Value

	if line.Eol() {
		node.Opcode = generateAbsoluteOpcode(node, mode, integerValueString)
		return node
	}

	line.ExpectSequence([][]string{
		{"comma"},
		{"character"},
	})

	node.Opcode = generateAbsoluteOpcode(node, mode, integerValueString)

	return node
}

func generateAbsoluteOpcode(node *node.Node, mode string, value string) int {
	absoluteNoModeBytePrefix := 0xAD
	absoluteXModeBytePrefix := 0xBD
	absoluteYModeBytePrefix := 0xB9
	bytes := byte.StringToByteSequence(value)
	var opcode int

	if mode == "x" {
		opcode = absoluteXModeBytePrefix<<16 | bytes[1]<<8 | bytes[0]
	} else if mode == "y" {
		opcode = absoluteYModeBytePrefix<<16 | bytes[1]<<8 | bytes[0]
	} else {
		opcode = absoluteNoModeBytePrefix<<16 | bytes[1]<<8 | bytes[0]
	}

	return opcode
}
