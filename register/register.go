package register

import "strings"

type Register uint16

const (
	REG0 Register = iota
	REG1
	REG2
	REG3
	REG4
	REG5
	REG6
	REG7
)

func Translate(s string) Register {
	switch strings.ToUpper(s) {
	case "REG0":
		return REG0
	case "REG1":
		return REG1
	case "REG2":
		return REG2
	case "REG3":
		return REG3
	case "REG4":
		return REG4
	case "REG5":
		return REG5
	case "REG6":
		return REG6
	case "REG7":
		return REG7
	default:
		return 0xffff
	}
}
