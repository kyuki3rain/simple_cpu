package simulator

import (
	"bytes"
	"fmt"

	"github.com/kyuki3rain/simple_cpu/opcode"
	"github.com/kyuki3rain/simple_cpu/order"
	"github.com/kyuki3rain/simple_cpu/parser"
)

type Simulator struct {
	pc       uint16
	ram      [256]uint16
	rom      [256]uint16
	register [8]uint16
	flag     bool
	status   bool
}

func Init(program string) *Simulator {
	s := &Simulator{}
	s.pc = 0

	orders := parser.Parse(program)
	for i, o := range orders {
		s.rom[i] = o.Encode()
	}

	return s
}

func (s *Simulator) Boot() {
	s.status = true

	for {
		s.Step()
		if !s.status {
			break
		}
	}

	s.status = false
}

func (s *Simulator) Step() {
	if !s.status {
		return
	}

	ord := order.Decode(s.rom[s.pc])
	s.pc++
	s.Exec(ord)
}

func (s Simulator) String() string {
	var res = bytes.NewBufferString("")
	for i := 0; i < 8; i++ {
		res.WriteString(fmt.Sprintf("REG%d: %d\n", i, s.register[i]))
	}
	return res.String()
}

func (s Simulator) GetRegister() [8]uint16 {
	return s.register
}

func (s *Simulator) Exec(o order.Order) {
	switch o.Opcode {
	case opcode.MOV:
		s.register[uint16(o.Operand1)] = s.register[uint16(o.Operand2)]
	case opcode.ADD:
		s.register[uint16(o.Operand1)] += s.register[uint16(o.Operand2)]
	case opcode.SUB:
		s.register[uint16(o.Operand1)] += s.register[uint16(o.Operand2)] ^ 0xffff + 1
	case opcode.AND:
		s.register[uint16(o.Operand1)] &= s.register[uint16(o.Operand2)]
	case opcode.OR:
		s.register[uint16(o.Operand1)] |= s.register[uint16(o.Operand2)]
	case opcode.SL:
		s.register[uint16(o.Operand1)] = s.register[uint16(o.Operand1)] << 1
	case opcode.SR:
		s.register[uint16(o.Operand1)] = uint16(s.register[uint16(o.Operand1)] >> 1)
	case opcode.SRA:
		s.register[uint16(o.Operand1)] = uint16(int16(s.register[uint16(o.Operand1)]) >> 1)
	case opcode.LDL:
		s.register[uint16(o.Operand1)] = s.register[uint16(o.Operand1)]&0xff00 + uint16(o.Operand2)
	case opcode.LDH:
		s.register[uint16(o.Operand1)] = s.register[uint16(o.Operand1)]&0x00ff + uint16(o.Operand2<<8)
	case opcode.CMP:
		s.flag = s.register[uint16(o.Operand1)] == s.register[uint16(o.Operand2)]
	case opcode.JE:
		if s.flag {
			s.pc = uint16(o.Operand2)
		}
	case opcode.JMP:
		s.pc = uint16(o.Operand2)
	case opcode.LD:
		s.register[uint16(o.Operand1)] = s.ram[uint16(o.Operand2)]
	case opcode.ST:
		s.ram[uint16(o.Operand2)] = s.register[uint16(o.Operand1)]
	case opcode.HLT:
		s.status = false
	}
}
