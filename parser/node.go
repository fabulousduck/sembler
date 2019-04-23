package parser

/*
Node contains all the info needed to generate an opcode
from a line of 6502 assembly

6502 opcodes are much like CH8 opcodes and are structured as follows

AABB

where AA tells us which instruction it is and in which mode
and   BB is the data for the instruction.


*/
type Node struct {
	Mode, ModeModifier string
	Instruction        string
}

/*
NewNode returns a new node pointer
*/
func NewNode() *Node {
	return new(Node)
}
