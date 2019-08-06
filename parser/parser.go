package parser

import (
	"strings"

	"github.com/fabulousduck/sembler/lexer"
	"github.com/fabulousduck/sembler/parser/lda"
	"github.com/fabulousduck/sembler/parser/mode"
	"github.com/fabulousduck/sembler/parser/node"
)

/*
Parser struct

structure on which all paring functions can be called
*/
type Parser struct {
	ParsedNodes []*node.Node
}

/*
NewParser returns a new parser structure pointer
*/
func NewParser() *Parser {
	return new(Parser)
}

/*
Parse takes a set of lexed lines and turns them into nodes

these nodes can then be made into opcodes
*/
func (p *Parser) Parse(lines []lexer.Line) {
	for _, line := range lines {
		switch line.Tokens[0].Type {
		case "load_accumelator":
			mode := GetInstructionMode(&line)
			p.ParsedNodes = append(p.ParsedNodes, lda.ParseLDA(&line, mode))
			break
		case "load_x_register":
			break
		case "load_y_register":
			break
		case "logical_right_shift":
			break
		case "no_operation":
			break
		case "bitwise_or_accumilator":
			break
		case "rotate_left":
			break
		case "rotate_right":
			break
		case "return_from_interupt":
			break
		case "return_from_subroutine":
			break
		case "subtract_with_carry":
			break
		case "store_accumilator":
			break
		case "store_x_register":
			break
		case "store_y_register":
			break
		case "force_interrupt":
			break
		case "clear_overflow_flag":
			break
		case "clear_interrupt_disable_status":
			break
		case "clear_decimal_mode":
			break
		case "clear_carry_flag":
			break
		case "set_interrupt_disable_status":
			break
		case "set_decimal_mode":
			break
		case "set_carry_flag":
			break
		case "jump_with_save":
			break
		case "jump":
			break
		case "pull_processor_status_stack":
			break
		case "push_processor_status_stack":
			break
		case "pull_accumilator_stack":
			break
		case "push_accumilator_stack":
			break
		case "transfer_x_stack_pointer":
			break
		case "transfer_stack_pointer_x":
			break
		case "transfer_y_accumilator":
			break
		case "transfer_accumilator_y":
			break
		case "transfer_x_accoumilator":
			break
		case "transfer_accumilator_x":
			break
		case "branch_overflow_set":
			break
		case "branch_overflow_clear":
			break
		case "branch_result_plus":
			break
		case "branch_not_equal":
			break
		case "branch_result_minus":
			break
		case "branch_equal":
			break
		case "branch_carry_set":
			break
		case "branch_carry_clear":
			break
		case "test_with_accumilator":
			break
		case "compare_memory_y":
			break
		case "compare_memory_x":
			break
		case "compare_memory_accumilator":
			break
		case "exclusive_memory_or":
			break
		case "and_memory_accumilator":
			break
		case "add_mem_accumilator_carry":
			break
		case "increment_x_one":
			break
		case "increment_y_one":
			break

		}
	}

}

/*
GetInstructionMode gets the mode in which the line was written
*/
func GetInstructionMode(line *lexer.Line) *mode.Mode {
	mode := mode.NewMode()

	//1 is the index at which mode is defined

	modeIndentifierChar := line.Tokens[1].Type
	operationValue := line.Tokens[2].Value

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

	if len(operationValue) == 4 {
		mode.Name = "absolute"
		return mode
	}

	mode.Name = "zeroPage"
	return mode

}

/*
FindInt looks for an integer value in a parsed line

2nd return value indicatest whether it has been found or not

1 means found
0 means not found
*/
func FindInt(l *lexer.Line) (lexer.Token, int) {
	for _, value := range l.Tokens {
		if value.Type == "integer" {
			return value, 1
		}
	}

	return l.Tokens[0], 0
}


func (p *Parser) validateSyntax() {
	
}