package order

import (
	"github.com/kyuki3rain/simple_cpu/opcode"
	"github.com/kyuki3rain/simple_cpu/operand"
	"github.com/kyuki3rain/simple_cpu/register"
)

type Order struct {
	Opcode   opcode.Opcode
	Operand1 operand.Operand
	Operand2 operand.Operand
}

func Init(order_string []string) Order {
	order := &Order{
		Opcode:   opcode.Translate(order_string[0]),
		Operand1: 0,
		Operand2: 0,
	}

	if len(order_string) == 1 {
		return *order
	}

	if register.Translate(order_string[1]) == 0xffff {
		order.Operand2 = operand.Translate(order_string[1])

		return *order
	}

	order.Operand1 = operand.Translate(order_string[1])
	if len(order_string) == 3 {
		order.Operand2 = operand.Translate(order_string[2])
	}
	return *order
}

func (o Order) String() string {
	return "opcode: " + o.Opcode.String() + ", operand1: " + o.Operand1.String() + ", operand2: " + o.Operand2.String()
}

func (o Order) Encode() uint16 {
	return uint16(o.Opcode)<<11 + uint16(o.Operand1)<<8 + uint16(o.Operand2)
}

func Decode(i uint16) Order {
	return Order{
		Opcode:   opcode.Opcode(i >> 11),
		Operand1: operand.Operand((i >> 8) % 8),
		Operand2: operand.Operand(i % 256),
	}
}
