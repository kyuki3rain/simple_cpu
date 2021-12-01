package opcode

import "strings"

type Opcode uint16

const (
	MOV Opcode = iota
	ADD
	SUB
	AND
	OR
	SL
	SR
	SRA
	LDL
	LDH
	CMP
	JE
	JMP
	LD
	ST
	HLT
	NIL
)

func (o Opcode) String() string {
	switch o {
	case MOV:
		return "MOV"
	case ADD:
		return "ADD"
	case SUB:
		return "SUB"
	case AND:
		return "AND"
	case OR:
		return "OR"
	case SL:
		return "SL"
	case SR:
		return "SR"
	case SRA:
		return "SRA"
	case LDL:
		return "LDL"
	case LDH:
		return "LDH"
	case CMP:
		return "CMP"
	case JE:
		return "JE"
	case JMP:
		return "JMP"
	case LD:
		return "LD"
	case ST:
		return "ST"
	case HLT:
		return "HLT"
	default:
		return "error"
	}
}

func Translate(s string) Opcode {
	switch strings.ToUpper(s) {
	case "MOV":
		return MOV
	case "ADD":
		return ADD
	case "SUB":
		return SUB
	case "AND":
		return AND
	case "OR":
		return OR
	case "SL":
		return SL
	case "SR":
		return SR
	case "SRA":
		return SRA
	case "LDL":
		return LDL
	case "LDH":
		return LDH
	case "CMP":
		return CMP
	case "JE":
		return JE
	case "JMP":
		return JMP
	case "LD":
		return LD
	case "ST":
		return ST
	case "HLT":
		return HLT
	default:
		return NIL
	}
}
