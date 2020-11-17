package parser

import (
	"strconv"

	"github.com/fabulousduck/sembler/lexer"
	"github.com/fabulousduck/sembler/parser/byte"
	"github.com/fabulousduck/sembler/parser/node"
)

/*
ParseZeroPage parses an instruction in zeropage form
*/
func (p *Parser) ParseZeroPage(line *lexer.Line, mode string) *node.Node {
	node := node.NewNode()
	var integerValue string
	node.Instruction = line.CurrentToken().Type

	//check if a label is used
	if line.NextToken().Type == "string" {
		label := p.getLabelByName(line.CurrentToken().Value)
		integerValue = strconv.Itoa(label.Pos)
		line.Advance()
	} else {

		line.Expect([]string{"dollar"})
		line.Advance()

		line.Expect([]string{"integer"})
		line.Advance()
		integerValue = line.CurrentToken().Value
	}

	if mode == "x" || mode == "y" {
		line.ExpectSequence([][]string{
			{"comma"},
			{"character"},
		})
		node.Opcode = getOpcodeForZeroPage(node.Instruction, mode)<<8 | byte.StringToByteSequence(integerValue)[0]
		return node
	}

	node.Opcode = getOpcodeForZeroPage(node.Instruction, "0")<<8 | byte.StringToByteSequence(integerValue)[0]
	return node
}

func getOpcodeForZeroPage(instruction string, mode string) int {

	/*the slices values are represented as follows
		[x,y,0]
	these are modes*/
	opcodeMap := map[string][]int{
		"load_accumulator":           {0xB5, 0x0, 0xA5},
		"load_x_register":            {0x0, 0xB6, 0xA6},
		"load_y_register":            {0xB4, 0x0, 0xA4},
		"logical_right_shift":        {0x56, 0x0, 0x46},
		"bitwise_or_accumulator":     {0x15, 0x0, 0x05},
		"rotate_left":                {0x36, 0x0, 0x26},
		"rotate_right":               {0x76, 0x0, 0x66},
		"subtract_with_carry":        {0xF5, 0x0, 0xE5},
		"store_accumulator":          {0x95, 0x0, 0x85},
		"store_x_register":           {0x0, 0x96, 0x86},
		"store_y_register":           {0x94, 0x0, 0x84},
		"add_mem_accumulator_carry":  {0x75, 0x0, 0x65},
		"and_memory_accumulator":     {0x35, 0x0, 0x25},
		"arithmetic_shift_left":      {0x16, 0x0, 0x06},
		"test_with_accumulator":      {0x0, 0x0, 0x24},
		"compare_memory_accumulator": {0xD5, 0x0, 0xC5},
		"compare_memory_x":           {0x0, 0x0, 0xE4},
		"compare_memory_y":           {0x0, 0x0, 0xC4},
		"decrement_memory":           {0xD6, 0x0, 0xC6},
		"exclusive_memory_or":        {0x55, 0x0, 0x45},
		"increment_memory":           {0xF6, 0x0, 0xE6},
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
