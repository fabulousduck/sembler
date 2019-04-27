package parser

import (
	"testing"

	"github.com/fabulousduck/sembler/lexer"
	"github.com/fabulousduck/sembler/parser/mode"
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
		mode := GetInstructionMode(&lexer.Lines[0])

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
