package jay

import (
	"bytes"
	"fmt"
	"log"
	"strings"
)

// MakeSize ...
func (s *Struct) MakeSize(o Option, b *bytes.Buffer) {
	//b.WriteString(fmt.Sprintf(
	//	"func (%s *%s) SizeJ() int {\n// fields: %d\n}\n",
	//	s.ReceiverName(),
	//	s.name,
	//	len(s.fixedLen),
	//))

	var qty uint
	// TODO add bool sizes
	var variableStructs []string
	for _, x := range s.fixedLen {
		qty += isLen(x.typ)
		if x.typ == "struct" {
			variableStructs = append(variableStructs, x.name)
		}
	}
	var variableFields []string
	for _, v := range s.variableLen {
		qty += isLen(v.typ)
		if v.typ == "string" { //} isLenVariable(v.typ) {
			variableFields = append(variableFields, v.name)
		} else if v.typ == "struct" {
			variableStructs = append(variableStructs, v.name)
		}
	}
	b.WriteString(fmt.Sprintf(
		"func (%s *%s) SizeJ() int {\nreturn %d%s%s\n}\n",
		s.ReceiverName(),
		s.name,
		qty,
		lengths(variableFields, s.ReceiverName()),
		structs(variableStructs, s.ReceiverName()),
	))
}
func (s *Struct) calcSize(o Option) (qty uint) {
	// TODO add bool sizes
	for _, x := range s.fixedLen {
		qty += o.typeFuncSize(x.typ)
	}
	for _, v := range s.variableLen {
		qty += o.typeFuncSize(v.typ)
	}
	return qty
}
func lengths(names []string, receiver string) string {
	if len(names) == 0 {
		return ""
	}
	receiver = fmt.Sprintf("+len(%s.", receiver)
	return receiver + strings.Join(names, ")"+receiver) + ")"
}

func structs(names []string, receiver string) string {
	if len(names) == 0 {
		return ""
	}

	//receiver = fmt.Sprintf(".SizeJ()+%s.", receiver)
	return "+" + receiver + "." + strings.Join(names, ".SizeJ()+"+receiver+".") + ".SizeJ()"
	//+c.Wheels.SizeJ()+c.Gearbox.SizeJ()+c.Roof.SizeJ()
}

func (o Option) typeFuncSize(typ string) (size uint) {
	switch typ {
	case "string", "bool", "byte", "uint8", "int8":
		return 1
	case "int16", "uint16":
		return 2
	case "int32", "rune", "float32", "uint32":
		return 4
	case "float64", "int64", "uint64":
		return 8
	case "int":
		if o.FixedIntSize {
			if o.Is32bit {
				return 4
			}
			return 8
		}
		return 1
	case "uint":
		if o.FixedUintSize {
			if o.Is32bit {
				return 4
			}
			return 8
		}
		return 1

	default:
		log.Printf("not function set for type %s yet", typ)
		return 0
	}
}
