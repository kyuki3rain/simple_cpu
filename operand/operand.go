package operand

import (
	"strconv"

	"github.com/kyuki3rain/simple_cpu/register"
)

type Operand uint16

func Translate(s string) Operand {
	if reg := register.Translate(s); reg != 0xffff {
		return Operand(reg)
	}

	if i, err := strconv.Atoi(s); err == nil {
		return Operand(i)
	}

	return Operand(0xffff)
}

func (o Operand) String() string {
	return strconv.Itoa(int(o))
}
