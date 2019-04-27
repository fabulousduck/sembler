package node

import "github.com/fabulousduck/sembler/parser/mode"

/*
Node contains all the info needed to generate an opcode
from a line of 6502 assembly

6502 opcodes are much like CH8 opcodes and are structured as follows

AABB

where AA tells us which instruction it is and in which mode
and   BB is the data for the instruction.

this is not true for absolute states

in the case of an absolute we have

AABBBBB

where 4400 becomes 0044

Mode: the mode the assembly was written in
Instruction: the instruction (LDA for example)
Opcode: the opcode produced from the line of assembly
*/
type Node struct {
	Mode        *mode.Mode
	Instruction string
	Opcode      uint32
}

/*
NewNode returns a new node pointer
*/
func NewNode() *Node {
	return new(Node)
}
