package mode

import (
	"testing"

	"github.com/fabulousduck/sembler/lexer"
)

type outlierOpcodeTest struct {
	line         *lexer.Line
	expectedMode *Mode
}

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

		if mode.Name != value.Name {
			T.Errorf(" \nline: %s\nfail: name\nexpect: %s\ngot:    %s\n", key, value.Name, mode.Name)
			T.Fail()
		}

		if mode.Variable != value.Variable {
			T.Errorf(" \nline: %s\nfail: variable\nexpect: %s\ngot:    %s\n", key, value.Variable, mode.Variable)
			T.Fail()
		}
	}
}

func TestGetModeForOutlierOpcodes(T *testing.T) {

	testCases := map[string]outlierOpcodeTest{
		"JSR": outlierOpcodeTest{
			line: &lexer.Line{
				Tokens: []lexer.Token{
					lexer.Token{Value: "JSR", Type: "string", Line: 1, Col: 1},
					lexer.Token{Value: "$", Type: "dollar", Line: 1, Col: 2},
					lexer.Token{Value: "4400", Type: "integer", Line: 1, Col: 3},
				},
				Raw:          "JSR $4400",
				CurrentIndex: 0,
				LineIndex:    0,
			},
			expectedMode: &Mode{Name: "absolute", Variable: ""},
		},
		"BRK": outlierOpcodeTest{
			line: &lexer.Line{
				Tokens: []lexer.Token{
					lexer.Token{Value: "BRK", Type: "string", Line: 1, Col: 1},
				},
				Raw:          "BRK",
				CurrentIndex: 0,
				LineIndex:    0,
			},
			expectedMode: &Mode{Name: "relative", Variable: ""},
		},
	}

	for _, testCase := range testCases {
		mode := getModeForOutlierOpcodes(testCase.line)
		if testCase.expectedMode.Name != mode.Name {
			T.Errorf("result mode name did not match expected name\ngot: %s\nexpected: %s\n", mode.Name, testCase.expectedMode.Name)
			T.Fail()
		}
		if testCase.expectedMode.Variable != mode.Variable {
			T.Errorf("result mode variable did not match expected variable\ngot: %s\nexpected: %s\n", mode.Variable, testCase.expectedMode.Variable)
			T.Fail()
		}
	}

}
