package parser

import (
	"strconv"

	"github.com/fabulousduck/sembler/lexer"
	"github.com/fabulousduck/sembler/parser/byte"
	"github.com/fabulousduck/sembler/parser/node"
)

/*
ParseAbsolute parses an instruction in absolute form
*/
func (p *Parser) ParseAbsolute(line *lexer.Line, mode string) *node.Node {
	node := node.NewNode()
	var integerValueString string

	node.Instruction = line.CurrentToken().Type

	if line.NextToken().Type == "string" {
		label := p.getLabelByName(line.NextToken().Value)
		integerValueString = strconv.Itoa(label.Pos)
		line.Advance()
	} else {
		line.Expect([]string{"dollar"})
		line.Advance()

		line.Expect([]string{"integer"})
		line.Advance()
		integerValueString = line.CurrentToken().Value
	}

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
	if len(bytes) < 2 {
		bytes = byte.AppendBytes(bytes, 2)
	}

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
		"load_accumulator":           {0xBD, 0xB9, 0xAD},
		"jump_with_save":             {0x0, 0x0, 0x20},
		"load_x_register":            {0x0, 0xBE, 0xAE},
		"load_y_register":            {0xBC, 0x0, 0xAC},
		"logical_right_shift":        {0x5E, 0x0, 0x4E},
		"bitwise_or_accumulator":     {0x1D, 0x19, 0x0D},
		"rotate_left":                {0x3E, 0x0, 0x2E},
		"rotate_right":               {0x7E, 0x0, 0x6E},
		"subtract_with_carry":        {0xFD, 0xF9, 0xED},
		"store_accumulator":          {0x9D, 0x99, 0x8D},
		"store_x_register":           {0x0, 0x0, 0x8E},
		"store_y_register":           {0x0, 0x0, 0x8C},
		"add_mem_accumulator_carry":  {0x7D, 0x79, 0x6D},
		"and_memory_accumulator":     {0x3D, 0x39, 0x2D},
		"arithmetic_shift_left":      {0x1E, 0x0, 0x0E},
		"test_with_accumulator":      {0x0, 0x0, 0x2C},
		"compare_memory_accumulator": {0xDD, 0xD9, 0xCD},
		"compare_memory_x":           {0x0, 0x0, 0xEC},
		"compare_memory_y":           {0x0, 0x0, 0xC4},
		"decrement_memory":           {0xDE, 0x0, 0xCE},
		"exclusive_memory_or":        {0x5D, 0x59, 0x4D},
		"increment_memory":           {0xFE, 0x0, 0xEE},
		"jump":                       {0x0, 0x0, 0x4C},
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
