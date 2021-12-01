package parser

import (
	"github.com/kyuki3rain/simple_cpu/opcode"
	"github.com/kyuki3rain/simple_cpu/registor"
)

type Order struct {
	opcode   opcode.Opcode
	operands []registor.Registor
}

func Perse(program string) [256]int {
	strings.split(program, "\n")
}
