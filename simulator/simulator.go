package simulator

import "github.com/kyuki3rain/simple_cpu/registor"

type Simulator struct {
	pc       int
	ram      [256]int
	rom      [256]int
	registor map[registor.Registor]int
}

func (s Simulator) Init(program string) {

}
