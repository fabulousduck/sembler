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
	node.Instruction = line.Tokens[0].Type

	line.ExpectSequence([][]string{
		{"hashtag"},
		{"dollar"},
	})

	line.Expect([]string{"integer"})
	line.Advance()
	integerValue := line.CurrentToken().Value

	node.Opcode = getOpcodeForImmidiate(node.Instruction)<<8 | byte.StringToByteSequence(integerValue)[0]

	return node
}

func getOpcodeForImmidiate(instruction string) int {
	opcodeMap := map[string]int{
		"load_accumulator": 0xA9,
	}

	if value, ok := opcodeMap[instruction]; ok {
		return value
	}
	return 0x0
}
