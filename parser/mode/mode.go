package mode

import (
	"github.com/fabulousduck/sembler/lexer"
)

/*
Mode is a struct holding info about the operation mode
*/
type Mode struct {
	Name, Variable string
}

/*
NewMode returns a new mode pointer
*/
func NewMode() *Mode {
	return new(Mode)
}

/*
GetInstructionMode gets the mode in which the line was written
*/
func GetInstructionMode(line *lexer.Line) *Mode {
	//check if its an implied instruction like BRK
	//implied instructions only have one token which is itself (the identifier)
	if len(line.Tokens) == 1 {
		return &Mode{"implied", ""}
	}

	//check preemptively if its a special case
	//instruction
	outlierMode := getModeForOutlierOpcodes(line)
	if outlierMode.Name != "" {
		return outlierMode
	}

	switch line.NextToken().Value {
	case "A":
		return &Mode{"accumulator", ""}
	case "$": //line is in hex mode. we dont care about this here
		line.Advance()
		break
	case "#":
		return &Mode{"immidiate", ""}
	case "(":
		return &Mode{"indirect", getIndirectVariable(line)}
	default: //labels
		return getModeForOutlierOpcodes(line)
	}

	return &Mode{}
}

func getModeForOutlierOpcodes(line *lexer.Line) *Mode {
	mode := NewMode()

	switch line.CurrentToken().Value {
	case "JSR":
		mode.Name = "absolute"
		mode.Variable = ""
		break
	case "BRK":
	case "BPL":
	case "BMI":
	case "BVC":
	case "BVS":
	case "BCC":
	case "BCS":
	case "BNE":
	case "BEQ":
		mode.Name = "relative"
		mode.Variable = ""
		break
	}

	return mode
}

func getIndirectVariable(line *lexer.Line) string {
	if line.HasSingleChar("X") {
		return "x"
	} else if line.HasSingleChar("Y") {
		return "y"
	} else {
		return "0"
	}
}
