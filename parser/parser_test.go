package parser

import (
	"testing"

	"github.com/fabulousduck/sembler/lexer"
)

func TestModeDetect(T *testing.T) {

	testCases := map[string]Mode{
		"ADC #$44":    Mode{"Immidiate", ""},
		"ADC $44":     Mode{"zeroPage", ""},
		"ADC $44,X":   Mode{"zeroPage", "x"},
		"ADC $4400":   Mode{"absolute", ""},
		"ADC $4400,X": Mode{"absolute", "x"},
		"ADC $4400,Y": Mode{"absolute", "y"},
		"ADC ($44,X)": Mode{"indirect", "x"},
		"ADC ($44),Y": Mode{"indirect", "y"},
	}

	T.Logf("fuck")

	for key, value := range testCases {
		T.Logf("testing: %s\n", key)

		lexer := lexer.NewLexer("mode test", key)
		parser := NewParser()
		lexer.Lex()
		mode := parser.getInstructionMode(&lexer.Lines[0])

		if mode.Name != value.Name {
			T.Errorf("line: %s\n fail: name\n expect: %s\n got:    %s\n", key, value.Name, mode.Name)
			T.Fail()
		}

		if mode.Variable != value.Variable {
			T.Errorf("line: %s\n fail: variable\n expect: %s\n got:    %s\n", key, value.Variable, mode.Variable)
			T.Fail()
		}
		T.Logf("line: %s\n success: name variable\n expect: %s %s\n got:    %s %s\n", key, value.Name, value.Variable, mode.Name, mode.Variable)

	}
}
