package parser

import (
	"github.com/fabulousduck/sembler/lexer"
	"github.com/fabulousduck/sembler/parser/byte"
	"github.com/fabulousduck/sembler/parser/node"
)

/*
ParseImmidiate parses an instruction in immidiate form
*/
func ParseImmidiate(line *lexer.Line) *node.Node {
	node := node.NewNode()
	immidiateModeBytePrefix := 0xA9

	node.Instruction = "load_accumelator"

	line.ExpectSequence([][]string{
		{"hashtag"},
		{"dollar"},
	})

	line.Expect([]string{"integer"})
	line.Advance()
	integerValue := line.CurrentToken().Value

	node.Opcode = immidiateModeBytePrefix<<8 | byte.StringToByteSequence(integerValue)[0]

	return node
}
