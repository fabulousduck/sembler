package mbi

import (
	"testing"

	"github.com/fabulousduck/sembler/parser/node"

	"github.com/fabulousduck/sembler/lexer"
	"github.com/fabulousduck/sembler/parser/mode"
)

func TestLDAParsing(T *testing.T) {
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
		mbiNode := ParseMBI(&lexer.Lines[0], value.Mode)

		if mbiNode.Opcode != value.Opcode {
			T.Errorf(" \nline: %s\nfail: opcode\nexpect: %x\ngot:    %x\n", key, value.Opcode, mbiNode.Opcode)
			T.FailNow()

		}

		T.Logf(" \nline: %s\nsuccess: opcode\nexpect: %x\ngot:    %x\n\n", key, value.Opcode, mbiNode.Opcode)

	}

}
