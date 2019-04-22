package parser

/*
Node contains all the info needed to generate an opcode
from a line of 6502 assembly
*/
type Node struct {
	Mode        string
	Instruction string
}
