package lda

import (
	"fmt"
	"os"
	"strconv"

	"github.com/davecgh/go-spew/spew"

	"github.com/fabulousduck/sembler/lexer"
	"github.com/fabulousduck/sembler/parser/mode"
	"github.com/fabulousduck/sembler/parser/node"
)

/*
ParseLDA parses an lda line into an opcode node
*/
func ParseLDA(line *lexer.Line, mode *mode.Mode) *node.Node {

	modeBytePrefixes := map[string]byte{
		"immidiate": 0xA9,
		"zeroPage":  0xA5,
		"zeroPageX": 0xB5,
		"absolute":  0xAD,
		"absoluteX": 0xBD,
		"absoluteY": 0xB9,
		"indirectX": 0xA1,
		"indirectY": 0xB1,
	}

	node := node.NewNode()

	node.Instruction = "load_accumelator"

	spew.Dump(line)

	intValue, err := strconv.Atoi(line.Tokens[2].Value)
	if err != nil {
		fmt.Printf("invalid value %s\n", line.Tokens[2].Value)
		os.Exit(65)
	}

	switch mode.Name {
	case "immidiate":
		node.Opcode = uint32(modeBytePrefixes["immidiate"]<<2 | byte(intValue))
		break
	case "indirect":
		if mode.Variable == "x" {
			node.Opcode = uint32(modeBytePrefixes["indirectX"]<<2 | byte(intValue))
			break
		}
		node.Opcode = uint32(modeBytePrefixes["indirectY"]<<2 | byte(intValue))
		break
	case "absolute":
		//check if this splitting is correct
		intRS, err := strconv.Atoi(line.Tokens[2].Value[:2])
		intLS, err := strconv.Atoi(line.Tokens[2].Value[2:])

		if err != nil {
			fmt.Printf("LDA absolute call with non 4 length hex")
			os.Exit(65)
		}

		if mode.Variable == "x" {
			node.Opcode = uint32(modeBytePrefixes["absoluteX"]<<4 | byte(intLS))
			node.Opcode = uint32(byte(node.Opcode<<2) | byte(intRS))
			break
		} else {

		}
		break
	case "zeroPage":
		break
	}

	return node
}
