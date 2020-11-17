package mode

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/fabulousduck/sembler/lexer"
)

func TestModeDetect(T *testing.T) {

	testCases := map[string]Mode{
		"ADC #$44":    Mode{"immidiate", ""},
		"ADC $44":     Mode{"zeroPage", ""},
		"ADC $44,X":   Mode{"zeroPage", "x"},
		"ADC $4400":   Mode{"absolute", ""},
		"ADC $4400,X": Mode{"absolute", "x"},
		"ADC $4400,Y": Mode{"absolute", "y"},
		"ADC ($44,X)": Mode{"indirect", "x"},
		"ADC ($44),Y": Mode{"indirect", "y"},
	}

	for key, value := range testCases {

		lexer := lexer.NewLexer("mode test", key)
		lexer.Lex()
		mode := GetInstructionMode(&lexer.Lines[0])

		spew.Dump(mode)

		if mode.Name != value.Name {
			T.Errorf(" \nline: %s\nfail: name\nexpect: %s\ngot:    %s\n", key, value.Name, mode.Name)
			T.Fail()
		}

		if mode.Variable != value.Variable {
			T.Errorf(" \nline: %s\nfail: variable\nexpect: %s\ngot:    %s\n", key, value.Variable, mode.Variable)
			T.Fail()
		}

		// T.Logf(" \nline: %s\nsuccess: name + variable\nexpect: %s %s\ngot:    %s %s\n\n", key, value.Name, value.Variable, mode.Name, mode.Variable)

	}
}
