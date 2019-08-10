package parser

import (
	"github.com/fabulousduck/sembler/lexer"
	"github.com/fabulousduck/sembler/parser/byte"
	"github.com/fabulousduck/sembler/parser/node"
)

/*
ParseZeroPage parses an instruction in zeropage form
*/
func ParseZeroPage(line *lexer.Line, mode string) *node.Node {
	node := node.NewNode()
	zeroPageNoModeBytePrefix := 0xA5
	zeroPageXModeBytePrefix := 0xB5

	node.Instruction = "load_accumelator"

	line.Expect([]string{"dollar"})
	line.Advance()

	line.Expect([]string{"integer"})
	line.Advance()
	integerValue := line.CurrentToken().Value

	if mode == "x" {
		line.ExpectSequence([][]string{
			{"comma"},
			{"character"},
		})
	}

	if mode == "x" {
		node.Opcode = zeroPageXModeBytePrefix<<8 | byte.StringToByteSequence(integerValue)[0]
	} else {
		node.Opcode = zeroPageNoModeBytePrefix<<8 | byte.StringToByteSequence(integerValue)[0]
	}

	return node
}
