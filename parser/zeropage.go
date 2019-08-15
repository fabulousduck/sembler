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

	node.Instruction = "load_accumulator"

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
		node.Opcode = getOpcodeForZeroPage(node.Instruction, "x")<<8 | byte.StringToByteSequence(integerValue)[0]
		return node
	}

	node.Opcode = getOpcodeForZeroPage(node.Instruction, "0")<<8 | byte.StringToByteSequence(integerValue)[0]
	return node
}

func getOpcodeForZeroPage(instruction string, mode string) int {
	/*the slices values are represented as follows
		[x,y]
	these are modes*/
	opcodeMap := map[string][]int{
		"load_accumulator": {0xB5, 0xA5},
	}

	if value, ok := opcodeMap[instruction]; ok {
		switch mode {
		case "x":
			return value[0]
		case "0":
			return value[1]
		}
	}
	return 0x0
}
