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

	node.Instruction = line.Tokens[0].Type

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
	bytes := byte.StringToByteSequence(value)
	if mode != "x" && mode != "y" {
		mode = "0"
	}
	return getOpcodeForAbsolute(node.Instruction, mode)<<16 | bytes[1]<<8 | bytes[0]
}

func getOpcodeForAbsolute(instruction string, mode string) int {
	/*the slices values are represented as follows
		[x,y,0]
	these are modes*/
	opcodeMap := map[string][]int{
		"load_accumulator": {0xBD, 0xB9, 0xAD},
	}

	if value, ok := opcodeMap[instruction]; ok {
		switch mode {
		case "x":
			return value[0]
		case "y":
			return value[1]
		case "0":
			return value[2]
		}
	}
	return 0x0
}
