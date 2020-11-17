package parser

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/fabulousduck/sembler/lexer"
	"github.com/fabulousduck/sembler/parser/mode"
	"github.com/fabulousduck/sembler/parser/node"
)

func TestModeParsing(T *testing.T) {
	testCases := map[string]node.Node{
		"LDA #$44":    node.Node{&mode.Mode{"immidiate", ""}, "load_accumelator", 0xA944, false},
		"LDA $44":     node.Node{&mode.Mode{"zeroPage", ""}, "load_accumelator", 0xA544, false},
		"LDA $44,X":   node.Node{&mode.Mode{"zeroPage", "x"}, "load_accumelator", 0xB544, false},
		"LDA $4400":   node.Node{&mode.Mode{"absolute", ""}, "load_accumelator", 0xAD0044, false},
		"LDA $4400,X": node.Node{&mode.Mode{"absolute", "x"}, "load_accumelator", 0xBD0044, false},
		"LDA $4400,Y": node.Node{&mode.Mode{"absolute", "y"}, "load_accumelator", 0xB90044, false},
		"LDA ($44,X)": node.Node{&mode.Mode{"indirect", "x"}, "load_accumelator", 0xA144, false},
		"LDA ($44),Y": node.Node{&mode.Mode{"indirect", "y"}, "load_accumelator", 0xB144, false},
	}

	for key, value := range testCases {

		lexer := lexer.NewLexer("mode test", key)
		lexer.Lex()
		p := NewParser()
		parsedLine := p.ParseLine(&lexer.Lines[0], value.Mode)

		if parsedLine.Opcode != value.Opcode {
			T.Errorf(" \nline: %s\nfail: opcode\nexpect: %x\ngot:    %x\n", key, value.Opcode, parsedLine.Opcode)
			T.FailNow()

		}
	}
}

func TestLabels(T *testing.T) {
	testCase := "NOP\nNOP\nLABEL LDA $44\nJSR LABEL"
	correctOpcodes := []int{0xEA, 0xEA, 0xA544, 0x200004}
	lexer := lexer.NewLexer("label test", testCase)
	lexer.Lex()
	p := NewParser()
	p.Parse(&lexer.Lines)

	for index, parsedNode := range p.ParsedNodes {
		if correctOpcodes[index] != parsedNode.Opcode {
			T.Errorf("\nline %s\nfail: opcode\nexpect: %x\ngot:    %x\n", strconv.Itoa(index), correctOpcodes[index], parsedNode.Opcode)
			T.FailNow()
		}
	}
}

func TestExampleFile(T *testing.T) {
	var fulldir strings.Builder

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fulldir.WriteString(dir)
	fulldir.WriteString("/test_files/memory.dbg")

	content, err := ioutil.ReadFile(fulldir.String())
	if err != nil {
		log.Fatal(err)
	}

	lexer := lexer.NewLexer("file test", string(content))
	lexer.Lex()
	p := NewParser()
	p.Parse(&lexer.Lines)
}
