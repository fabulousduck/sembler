package lexer

import "strings"

//determineType determines the type of a string character
func determineType(character string) string {
	usableChar := strings.ToLower(character)
	types := map[string][]string{
		"character":     []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "_"},
		"integer":       []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
		"dollar":        []string{"$"},
		"comma":         []string{","},
		"dot":           []string{"."},
		"at":            []string{"@"},
		"comment":       []string{";"},
		"left_bracket":  []string{"["},
		"right_bracket": []string{"]"},
		"double_dot":    []string{":"},
		"left_paren":    []string{"("},
		"right_paren":   []string{")"},
		"hashtag":       []string{"#"},
		"plus":          []string{"+"},
		"dash":          []string{"-"},
		"slash":         []string{"/"},
		"star":          []string{"*"},
		"newline":       []string{"\r", "\n"},
		"ignoreable":    []string{"\t", " "},
	}

	for key, values := range types {
		if Contains(usableChar, values) {
			return key
		}
	}
	return "undefined_symbol"
}

/*
Contains is a function that checks if a given character is present in a map
*/
func Contains(name string, list []string) bool {
	for i := 0; i < len(list); i++ {
		if string(list[i]) == name {
			return true
		}
	}
	return false
}

func getKeyword(token *Token) string {
	keywords := map[string]string{
		"LDA": "load_accumulator",
		"LDX": "load_x_register",
		"LDY": "load_y_register",
		"LSR": "logical_right_shift",
		"NOP": "no_operation",
		"ORA": "bitwise_or_accumilator",
		"ROL": "rotate_left",
		"ROR": "rotate_right",
		"RTI": "return_from_interupt",
		"RTS": "return_from_subroutine",
		"SBC": "subtract_with_carry",
		"STA": "store_accumilator",
		"STX": "store_x_register",
		"STY": "store_y_register",
		"BRK": "force_interrupt",
		"CLV": "clear_overflow_flag",
		"CLI": "clear_interrupt_disable_status",
		"CLD": "clear_decimal_mode",
		"CLC": "clear_carry_flag",
		"SEI": "set_interrupt_disable_status",
		"SED": "set_decimal_mode",
		"SEC": "set_carry_flag",
		"JSR": "jump_with_save",
		"JMP": "jump",
		"PLP": "pull_processor_status_stack",
		"PHP": "push_processor_status_stack",
		"PLA": "pull_accumilator_stack",
		"PHA": "push_accumilator_stack",
		"TXS": "transfer_x_stack_pointer",
		"TSX": "transfer_stack_pointer_x",
		"TYA": "transfer_y_accumilator",
		"TAY": "transfer_accumilator_y",
		"TXA": "transfer_x_accoumilator",
		"TAX": "transfer_accumilator_x",
		"BVS": "branch_overflow_set",
		"BVC": "branch_overflow_clear",
		"BPL": "branch_result_plus",
		"BNE": "branch_not_equal",
		"BMI": "branch_result_minus",
		"BEQ": "branch_equal",
		"BCS": "branch_carry_set",
		"BCC": "branch_carry_clear",
		"BIT": "test_with_accumilator",
		"CPY": "compare_memory_y",
		"CPX": "compare_memory_x",
		"CMP": "compare_memory_accumilator",
		"EOR": "exclusive_memory_or",
		"AND": "and_memory_accumilator",
		"ADC": "add_mem_accumilator_carry",
		"INX": "increment_x_one",
		"INY": "increment_y_one",
	}

	if val, ok := keywords[token.Value]; ok {
		return val
	}

	if len(token.Value) > 1 {
		return "string"
	}
	return token.Type
}

//DetermineStringType will determine the type of a given string
func DetermineStringType(str string) string {
	return determineType(string([]rune(str)[0]))
}
