package simulator_test

import (
	"testing"

	"github.com/kyuki3rain/simple_cpu/simulator"
)

func Test(t *testing.T) {
	program := `ldh Reg0 0
		ldl Reg0 0
		ldh Reg1 0
		ldl Reg1 1
		ldh Reg2 0
		ldl Reg2 0
		ldh Reg3 0
		ldl Reg3 10
		add Reg2 Reg1
		add Reg0 Reg2
		st  Reg0 64
		cmp Reg2 Reg3
		je  14
		jmp 8
		hlt`

	s := simulator.Init(program)
	s.Boot()
	reg := s.GetRegister()
	if reg[0] != 55 {
		t.Errorf("error!!!!!!!!!!!!!!!!!!")
	}
}
