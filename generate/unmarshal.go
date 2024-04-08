package generate

import (
	"bytes"
	"fmt"
	"github.com/speedyhoon/jay"
	"strings"
)

// makeUnmarshal ...
func (s *structTyp) makeUnmarshal(b *bytes.Buffer, o Option) {
	var byteIndex = uint(len(s.variableLen))
	buf := bytes.NewBuffer(nil)

	s.makeReadBools(buf, &byteIndex, s.receiver)
	s.readSingles(buf, &byteIndex, s.receiver, o)

	for i, f := range s.fixedLen {
		buf.WriteString(o.unmarshalLine(s, f, &byteIndex, "", "", i == 0, len(s.variableLen) == 0 && i == len(s.fixedLen)-1, ""))
		buf.WriteString("\n")
	}

	at, end := s.defineTrackingVars(buf, byteIndex)
	vLen := len(s.variableLen) - 1
	for i, f := range s.variableLen {
		at, end = s.tracking(buf, i, end, byteIndex, f.typ)
		buf.WriteString(o.unmarshalLine(s, f, &byteIndex, at, end, i == 0, i == vLen, lenVariable(i)))
		buf.WriteString("\n")
	}

	code := buf.Bytes()
	if len(code) == 0 {
		return
	}

	const exportedErr = "jay.ErrUnexpectedEOB"
	// Prevent panic: runtime error: index out of range
	var lengthCheck string
	if len(s.variableLen) == 0 {
		lengthCheck = fmt.Sprintf("if len(%s) < %d {\nreturn %s\n}", s.bufferName, byteIndex, exportedErr)
	} else {
		lengthCheck = fmt.Sprintf("%[1]s := len(%[2]s)\nif %[1]s < %[3]d {\nreturn %[4]s\n}", s.lengthName, s.bufferName, byteIndex, exportedErr)
	}
	variableLengthCheck := s.generateCheckSizes(exportedErr, byteIndex)

	bufWriteF(b,
		"func (%s *%s) UnmarshalJ(%s []byte) error {\n%s\n%s%sreturn nil\n}\n",
		s.receiver,
		s.name,
		s.bufferName,
		lengthCheck,
		variableLengthCheck,
		code,
	)
}

func (s *structTyp) generateCheckSizes(exportedErr string, totalSize uint) string {
	qty := len(s.variableLen)
	if qty == 0 {
		return ""
	}

	assignments, values, conditions := make([]string, qty), make([]string, qty), make([]string, qty)
	for i, f := range s.variableLen {
		assignments[i] = lenVariable(i)
		values[i] = fmt.Sprintf("int(%s[%d])", s.bufferName, i)
		if f.typ == "[]bool" {
			conditions[i] = fmt.Sprintf("%s(%s)", nameOf(jay.SizeBools, nil), assignments[i])
		} else {
			conditions[i] = assignments[i]
		}
	}

	return fmt.Sprintf(
		"%s := %s\nif %s < %d+%s {\nreturn %s\n}\n",
		strings.Join(assignments, ", "),
		strings.Join(values, ", "),
		s.lengthName,
		totalSize,
		strings.Join(conditions, "+"),
		exportedErr,
	)
}

func (o Option) unmarshalLine(s *structTyp, f field, byteIndex *uint, at, end string, isFirst, isLast bool, lenVar string) string {
	fun, _, totalSize := o.unmarshalFuncs(f, isLast)
	if fun == "" {
		// Unknown type, not supported yet.
		lg.Printf("no generateLine for type `%s` yet in unmarshalLine()", f.typ)
		return ""
	}

	start := *byteIndex
	*byteIndex += totalSize
	thisField := pkgSelName(s.receiver, f.name)

	switch f.typ {
	case "string":
		if f.isAliasDef {
			fun = f.aliasType
		}
		return fmt.Sprintf("%s = %s(%s)", thisField, fun, sliceExpr(s, f, at, end))
	case "[]byte", "[]uint8":
		if f.isAliasDef {
			fun = f.aliasType
		}
		if f.Required {
			return fmt.Sprintf("%s = %s", thisField, sliceExpr(s, f, at, end))
		}
		return fmt.Sprintf("if %s != 0 {\n%s = %s\n}", lenVar, thisField, sliceExpr(s, f, at, end))
	case "[]int8", "[]bool":
		return fmt.Sprintf("%s = %s(%s, %s)", thisField, fun, sliceExpr(s, f, at, end), lenVar)
	}

	if f.isArray() && f.arrayType == "int8" {
		values := make([]string, f.arraySize)
		for i := 0; i < f.arraySize; i++ {
			values[i] = fmt.Sprintf("%s(%s[%d])", f.arrayType, s.bufferName, i)
		}
		return fmt.Sprintf("%s = %s{%s}", thisField, f.typ, strings.Join(values, ", "))
	}

	fun = printFunc(fun, sliceExprU(s, f, start, *byteIndex))
	if f.isAliasDef {
		fun = printFunc(f.aliasType, fun)
	}
	return fmt.Sprintf("%s = %s", thisField, fun)
}

// unmarshalFuncs returns the function name to handle unmarshalling.
// `size` is the quantity of bytes required to represent the type.
func (o Option) unmarshalFuncs(f field, isLast bool) (funcName string, size, totalSize uint) {
	var c interface{}
	switch f.typ {
	case "byte", "uint8":
		if f.isAliasDef {
			return f.aliasType, 1, 1
		}
		return "", 1, 1
	case "int8":
		if f.isAliasDef {
			return f.aliasType, 1, 1
		}
		return "int8", 1, 1
	case "string":
		return "string", 0, 0
	case "int":
		if o.FixedIntSize {
			if o.Is32bit {
				c, size, totalSize = jay.ReadIntArch32, 4, 4
			}
			c, size, totalSize = jay.ReadIntArch64, 8, 8
			break
		}
		//c, size, totalSize = jay.ReadIntVariable, 0,1
		c, size, totalSize = jay.ReadInt, 0, 1
	case "int16":
		c, size, totalSize = jay.ReadInt16, 2, 2
	case "int32", "rune":
		c, size, totalSize = jay.ReadInt32, 4, 4
	case "float32":
		c, size, totalSize = jay.ReadFloat32, 4, 4
	case "float64":
		c, size, totalSize = jay.ReadFloat64, 8, 8
	case "int64":
		c, size, totalSize = jay.ReadInt64, 8, 8
	case "time.Duration":
		c, size, totalSize = jay.ReadDuration, 8, 8
	case "uint":
		if o.FixedUintSize {
			if o.Is32bit {
				c, size, totalSize = jay.ReadUintArch32, 4, 4
			}
			c, size, totalSize = jay.ReadUintArch64, 8, 8
			break
		}
		c, size, totalSize = jay.ReadUintVariable, 0, 1
	case "uint16":
		c, size, totalSize = jay.ReadUint16, 2, 2
	case "uint32":
		c, size, totalSize = jay.ReadUint32, 4, 4
	case "uint64":
		c, size, totalSize = jay.ReadUint64, 8, 8
	case "time.Time":
		if f.tagOptions.TimeNano {
			c, size, totalSize = jay.ReadTimeNano, 8, 8
		} else {
			c, size, totalSize = jay.ReadTime, 8, 8
		}
	case "[]int8":
		c, size, totalSize = jay.ReadInt8s, 0, 0
	case "[]bool":
		c, size, totalSize = jay.ReadBools, 0, 0

	case "[]byte", "[]uint8":
		c, size, totalSize = copyKeyword, 0, 0

	default:
		lg.Printf("no function set for type %s yet in unmarshalFuncs()", f.typ)
	}

	return nameOf(c, nil), size, totalSize
}
