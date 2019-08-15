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
	node.Instruction = line.Tokens[0].Type

	line.ExpectSequence([][]string{
		{"left_paren"},
		{"dollar"},
		{"integer"},
	})

	integerValue := line.CurrentToken().Value

	if mode == "x" {
		node.Opcode = getOpcodeForIndirect(node.Instruction, "x")<<8 | byte.StringToByteSequence(integerValue)[0]

		line.Expect([]string{"comma"})
		line.Advance()

		line.ExpectSequence([][]string{
			{"character"},
			{"right_paren"},
		})

	} else {
		node.Opcode = getOpcodeForIndirect(node.Instruction, "y")<<8 | byte.StringToByteSequence(integerValue)[0]

		line.ExpectSequence([][]string{
			{"right_paren"},
			{"comma"},
			{"character"},
		})
	}

	return node
}

func getOpcodeForIndirect(instruction string, mode string) int {
	/*the slices values are represented as follows
		[x,y]
	these are modes*/
	opcodeMap := map[string][]int{
		"load_accumulator": {0xA1, 0xB1},
	}

	if value, ok := opcodeMap[instruction]; ok {
		switch mode {
		case "x":
			return value[0]
		case "y":
			return value[1]
		}
	}
	return 0x0
}
