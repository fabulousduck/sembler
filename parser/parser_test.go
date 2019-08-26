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

func TestModeDetect(T *testing.T) {

	testCases := map[string]mode.Mode{
		"ADC #$44":    mode.Mode{"immidiate", ""},
		"ADC $44":     mode.Mode{"zeroPage", ""},
		"ADC $44,X":   mode.Mode{"zeroPage", "x"},
		"ADC $4400":   mode.Mode{"absolute", ""},
		"ADC $4400,X": mode.Mode{"absolute", "x"},
		"ADC $4400,Y": mode.Mode{"absolute", "y"},
		"ADC ($44,X)": mode.Mode{"indirect", "x"},
		"ADC ($44),Y": mode.Mode{"indirect", "y"},
	}

	for key, value := range testCases {

		lexer := lexer.NewLexer("mode test", key)
		lexer.Lex()
		mode := mode.GetInstructionMode(&lexer.Lines[0])

		if mode.Name != value.Name {
			T.Errorf(" \nline: %s\nfail: name\nexpect: %s\ngot:    %s\n", key, value.Name, mode.Name)
			T.Fail()
		}

		if mode.Variable != value.Variable {
			T.Errorf(" \nline: %s\nfail: variable\nexpect: %s\ngot:    %s\n", key, value.Variable, mode.Variable)
			T.Fail()
		}
		T.Logf(" \nline: %s\nsuccess: name variable\nexpect: %s %s\ngot:    %s %s\n\n", key, value.Name, value.Variable, mode.Name, mode.Variable)

	}
}

func TestModeParsing(T *testing.T) {
	testCases := map[string]node.Node{
		"LDA #$44":    node.Node{&mode.Mode{"immidiate", ""}, "load_accumelator", 0xA944},
		"LDA $44":     node.Node{&mode.Mode{"zeroPage", ""}, "load_accumelator", 0xA544},
		"LDA $44,X":   node.Node{&mode.Mode{"zeroPage", "x"}, "load_accumelator", 0xB544},
		"LDA $4400":   node.Node{&mode.Mode{"absolute", ""}, "load_accumelator", 0xAD0044},
		"LDA $4400,X": node.Node{&mode.Mode{"absolute", "x"}, "load_accumelator", 0xBD0044},
		"LDA $4400,Y": node.Node{&mode.Mode{"absolute", "y"}, "load_accumelator", 0xB90044},
		"LDA ($44,X)": node.Node{&mode.Mode{"indirect", "x"}, "load_accumelator", 0xA144},
		"LDA ($44),Y": node.Node{&mode.Mode{"indirect", "y"}, "load_accumelator", 0xB144},
	}

	for key, value := range testCases {

		lexer := lexer.NewLexer("mode test", key)
		lexer.Lex()
		p := NewParser()
		mbiNode := p.ParseLine(&lexer.Lines[0], value.Mode)

		if mbiNode.Opcode != value.Opcode {
			T.Errorf(" \nline: %s\nfail: opcode\nexpect: %x\ngot:    %x\n", key, value.Opcode, mbiNode.Opcode)
			T.FailNow()

		}

		T.Logf(" \nline: %s\nsuccess: opcode\nexpect: %x\ngot:    %x\n\n", key, value.Opcode, mbiNode.Opcode)

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
	fulldir.WriteString("/test_files/memory.asm")

	content, err := ioutil.ReadFile(fulldir.String())
	if err != nil {
		log.Fatal(err)
	}

	lexer := lexer.NewLexer("file test", string(content))
	lexer.Lex()
	p := NewParser()
	p.Parse(&lexer.Lines)
}
