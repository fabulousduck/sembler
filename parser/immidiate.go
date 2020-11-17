package parser

import (
	"github.com/fabulousduck/sembler/lexer"
	"github.com/fabulousduck/sembler/parser/byte"
	"github.com/fabulousduck/sembler/parser/node"
)

/*
ParseImmidiate parses an instruction in immidiate form
*/
func (p *Parser) ParseImmidiate(line *lexer.Line) *node.Node {
	node := node.NewNode()
	node.Instruction = line.CurrentToken().Type

	line.Expect([]string{
		"hashtag",
	})
	line.Advance()

	if line.NextToken().Type == "dollar" {
		node.ValueIsHex = true
		line.Advance()
	}

	line.Expect([]string{"integer"})
	line.Advance()

	integerValue := line.CurrentToken().Value

	node.Opcode = getOpcodeForImmidiate(node.Instruction)<<8 | byte.StringToByteSequence(integerValue)[0]

	return node
}

func getOpcodeForImmidiate(instruction string) int {
	opcodeMap := map[string]int{
		"load_accumulator":           0xA9,
		"load_x_register":            0xA2,
		"load_y_register":            0xA0,
		"bitwise_or_accumulator":     0x09,
		"subtract_with_carry":        0xE9,
		"compary_memory_x":           0xE0,
		"compare_memory_y":           0xC0,
		"compare_memory_accumulator": 0xC9,
		"and_memory_accumulator":     0x29,
		"add_mem_accumulator_carry":  0x69,
	}

	if value, ok := opcodeMap[instruction]; ok {
		return value
	}
	//TODO invalid mode for instruction error
	return 0x0
}
