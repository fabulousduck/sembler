//Package mbi stands for "Mode Based Instruction"
//this refers to a group of instructions that have a set mode
package mbi

import (
	"github.com/fabulousduck/sembler/lexer"
	"github.com/fabulousduck/sembler/parser/byte"
	"github.com/fabulousduck/sembler/parser/mode"
	"github.com/fabulousduck/sembler/parser/node"
)

type instructionModeSet struct {
	immidiate int
	zeroPage  int
	zeroPageX int
	absolute  int
	absoluteX int
	absoluteY int
	indirectX int
	indirectY int
}
type instructionModeMap map[string]instructionModeSet

/*
ParseMBI parses an MBI line into an opcode node
*/
func ParseMBI(line *lexer.Line, mode *mode.Mode) *node.Node {
	switch mode.Name {
	case "immidiate":
		return parseImmidiate(line)
	case "indirect":
		return parseIndirect(line, mode.Variable)
	case "absolute":
		return parseAbsolute(line, mode.Variable)
	case "zeroPage":
		return parseZeroPage(line, mode.Variable)
	}

	return node.NewNode()
}

func parseImmidiate(line *lexer.Line) *node.Node {
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

func parseIndirect(line *lexer.Line, mode string) *node.Node {
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

func parseAbsolute(line *lexer.Line, mode string) *node.Node {
	node := node.NewNode()

	node.Instruction = "load_accumelator"

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

func parseZeroPage(line *lexer.Line, mode string) *node.Node {
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

func generateAbsoluteOpcode(node *node.Node, mode string, value string) int {
	absoluteNoModeBytePrefix := 0xAD
	absoluteXModeBytePrefix := 0xBD
	absoluteYModeBytePrefix := 0xB9
	bytes := byte.StringToByteSequence(value)
	var opcode int

	if mode == "x" {
		opcode = absoluteXModeBytePrefix<<16 | bytes[1]<<8 | bytes[0]
	} else if mode == "y" {
		opcode = absoluteYModeBytePrefix<<16 | bytes[1]<<8 | bytes[0]
	} else {
		opcode = absoluteNoModeBytePrefix<<16 | bytes[1]<<8 | bytes[0]
	}

	return opcode
}

/*
IsMBI checks if an instruction is a generic MBI structured instruction
*/
func IsMBI(name string) bool {
	if _, ok := getMBIMap()[name]; ok {
		return true
	}
	return false
}

func getMBIMap() instructionModeMap {
	return instructionModeMap{
		"load_accumelator": instructionModeSet{
			immidiate: 0xA9,
			zeroPage:  0xA5,
			zeroPageX: 0xB5,
			absolute:  0xAD,
			absoluteX: 0xBD,
			absoluteY: 0xB9,
			indirectX: 0xA1,
			indirectY: 0xB1,
		},
	}
}