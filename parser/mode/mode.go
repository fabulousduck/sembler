package mode

import (
	"strings"

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
	var instructionNamePos int

	if lexer.GetKeyword(&line.Tokens[0]) == "string" {
		instructionNamePos = 1
	} else {
		instructionNamePos = 0
	}

	//check if the line is for a instruction
	//that we cant parse genericly
	if lexer.IsNonGenericInstruction(line.Tokens[instructionNamePos].Value) {
		return getModeForNonGenericInstruction(line, instructionNamePos)
	}
	return getModeForGenericInstruction(line)
}

func getModeForNonGenericInstruction(line *lexer.Line, instructionNamePos int) *Mode {
	mode := NewMode()

	switch line.Tokens[instructionNamePos].Value {
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
		mode.Name = "implied"
		mode.Variable = ""
		break
	}

	return mode
}

func getModeForGenericInstruction(line *lexer.Line) *Mode {
	var modeIndentifierChar string
	var operationValue string

	mode := NewMode()
	//check if its an implied instruction like BRK
	if len(line.Tokens) == 1 {
		mode.Name = "implied"
		mode.Variable = ""
		return mode
	}

	//check if there is a label in the line which messes with the offsets
	if lexer.GetKeyword(&line.Tokens[0]) == "string" {
		modeIndentifierChar = line.Tokens[2].Type
		operationValue = line.Tokens[3].Value
	} else {
		modeIndentifierChar = line.Tokens[1].Type
		operationValue = line.Tokens[2].Value
	}

	//final character for non direct operations is the last one
	XYNonDirectLocation := strings.ToLower(line.Tokens[len(line.Tokens)-1].Value)

	//check for x or y variables
	if XYNonDirectLocation == "x" || XYNonDirectLocation == ")" {
		mode.Variable = "x"
	}

	if XYNonDirectLocation == "y" {
		mode.Variable = "y"
	}

	//only immidiate mode starts with a #
	if modeIndentifierChar == "hashtag" {
		mode.Name = "immidiate"
		mode.Variable = ""
		return mode
	}

	//if it is encapsulated, we can assume it is an indirect operation
	if modeIndentifierChar == "left_paren" {
		mode.Name = "indirect"
		return mode
	}

	//if the value of the operation is 4 characters long,
	//it is assumed it is an absolute operation
	if len(operationValue) == 4 {
		mode.Name = "absolute"
		return mode
	}

	mode.Name = "zeroPage"
	return mode
}
