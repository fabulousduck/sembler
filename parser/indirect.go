package parser

import (
	"github.com/fabulousduck/sembler/lexer"
	"github.com/fabulousduck/sembler/parser/byte"
	"github.com/fabulousduck/sembler/parser/node"
)

/*
ParseIndirect parses an instruction in indirect form

*/
func ParseIndirect(line *lexer.Line, mode string) *node.Node {
	node := node.NewNode()
	indirectXModeBytePrefix := 0xA1
	indirectYModeBytePrefix := 0xB1

	node.Instruction = "load_accumelator"

	line.ExpectSequence([][]string{
		{"left_paren"},
		{"dollar"},
		{"integer"},
	})

	integerValue := line.CurrentToken().Value

	if mode == "x" {
		node.Opcode = indirectXModeBytePrefix<<8 | byte.StringToByteSequence(integerValue)[0]

		line.Expect([]string{"comma"})
		line.Advance()

		line.ExpectSequence([][]string{
			{"character"},
			{"right_paren"},
		})

	} else {
		node.Opcode = indirectYModeBytePrefix<<8 | byte.StringToByteSequence(integerValue)[0]

		line.ExpectSequence([][]string{
			{"right_paren"},
			{"comma"},
			{"character"},
		})
	}

	return node
}
