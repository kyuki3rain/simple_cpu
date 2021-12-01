package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kyuki3rain/simple_cpu/simulator"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("error1")
		return
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("error2")
		return
	}

	sim := simulator.Init(string(b))
	sim.Boot()
	fmt.Println(sim)
}
