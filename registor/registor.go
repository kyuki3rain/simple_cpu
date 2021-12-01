package registor

type Registor int

const (
	REG0 Registor = iota
	REG1
	REG2
	REG3
	REG4
	REG5
	REG6
	REG7
	NIL
)

func Translate(s string) Registor {
	switch s {
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
		return NIL
	}
}
