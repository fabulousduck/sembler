package parser

import (
	"github.com/fabulousduck/sembler/lexer"
)

func (p *Parser) parseLDA(line *lexer.Line) *Node {
	node := NewNode()

	//check for immidiate mode
	if line.Tokens[1].Value == "#" {
		node.Mode = "immidiate"
	}

	return node
}
