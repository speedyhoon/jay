package generate

import (
	"bytes"
	"fmt"
	"github.com/speedyhoon/jay"
	"log"
	"reflect"
	"runtime"
	"strings"
)

// makeUnmarshal ...
func (s *structTyp) makeUnmarshal(b *bytes.Buffer, o Option) {
	var byteIndex uint
	buf := bytes.NewBuffer(nil)

	s.makeReadBools(buf, &byteIndex, s.receiver)
	s.readSingles(buf, &byteIndex, s.receiver, o)

	var returnInlined bool
	for i, f := range s.fixedLen {
		buf.WriteString(o.unmarshalLine(f, &byteIndex, s.receiver, "", len(s.variableLen) == 0 && i == len(s.fixedLen)-1, &returnInlined, s.bufferName))
		buf.WriteString("\n")
	}

	var at string
	if len(s.variableLen) >= 2 {
		buf.WriteString("var ok bool\n")
	}

	if len(s.variableLen) == 1 {
		at = Utoa(byteIndex)
	} else if len(s.variableLen) >= 2 {
		bufWriteF(buf, "at:=%d\n", byteIndex)
		at = "at"
	}

	vLen := len(s.variableLen) - 1
	for i, f := range s.variableLen {
		isLast := i == vLen
		buf.WriteString(o.unmarshalLine(f, &byteIndex, s.receiver, at, isLast, &returnInlined, s.bufferName))
		if !isLast {
			bufWriteF(buf, "\nif !ok {\nreturn %s.ErrUnexpectedEOB\n}\n", pkgName)
		} else {
			buf.WriteString("\n")
		}
	}

	code := buf.Bytes()
	if len(code) == 0 || byteIndex == 0 {
		return
	}

	// Prevent panic: runtime error: index out of range
	lengthCheck := fmt.Sprintf("if len(%s) < %d {\nreturn jay.ErrUnexpectedEOB\n}", s.bufferName, byteIndex)

	var returnCode string
	if !returnInlined {
		returnCode = "return nil\n"
	}

	bufWriteF(b,
		"func (%s *%s) UnmarshalJ(%s []byte) error {\n%s\n%s%s}\n",
		s.receiver,
		s.name,
		s.bufferName,
		lengthCheck,
		code,
		returnCode,
	)
}

func (o Option) unmarshalLine(f field, byteIndex *uint, receiver, at string, isLast bool, returnInlined *bool, bufferName string) string {
	fun, size, totalSize := o.unmarshalFuncs(f, isLast)
	if fun == "" && size == 0 {
		// Unknown type, not supported yet.
		log.Printf("no generateLine for type `%s` yet in unmarshalLine()", f.typ)
		return ""
	}

	start := *byteIndex
	*byteIndex += totalSize
	thisField := fmt.Sprintf("%s.%s", receiver, f.name)

	switch size {
	case 1:
		//TODO  remove -- singles no longer needed!
		return fmt.Sprintf("%s=%s", thisField, printFunc(fun, fmt.Sprintf("%s[%d]", bufferName, start)))
	default:
		if start == 0 {
			return fmt.Sprintf("%s = %s", thisField, printFunc(fun, fmt.Sprintf("%s[:%d]", bufferName, *byteIndex)))
		} else {
			return fmt.Sprintf("%s = %s", thisField, printFunc(fun, fmt.Sprintf("%s[%d:%d]", bufferName, start, *byteIndex)))
		}
	case 0:
		// Variable length size.
		slice := bufferName
		idx := at
		if start >= 1 {
			if at == "" {
				slice = fmt.Sprintf("%s[%d:]", bufferName, start)
				idx = "0"
			} else if at != "0" {
				slice = fmt.Sprintf("%s[%s:]", bufferName, at)
			}
		}

		switch f.typ {
		case "struct":
			return fmt.Sprintf("%s.%s", thisField, printFunc(fun, slice))
		}

		if isLast {
			if f.typ == "string" || f.typ == "[]byte" {
				*returnInlined = true
				return fmt.Sprintf("return %s", printFunc(fun, slice, "&"+thisField))
			}
			return fmt.Sprintf("%s, at, ok = %s", thisField, printFunc(fun, slice))
		} else {
			return fmt.Sprintf("%s, at, ok = %s", thisField, printFunc(fun, slice, idx))
		}
	}
}

// unmarshalFuncs returns the function name to handle unmarshalling.
// `size` is the quantity of bytes required to represent the type.
func (o Option) unmarshalFuncs(f field, isLast bool) (funcName string, size, totalSize uint) {
	var c interface{}
	switch f.typ {
	case "byte", "uint8":
		return "", 1, 1
	case "int8":
		return "int8", 1, 1
	case "string":
		if isLast {
			c, size, totalSize = jay.ReadStringPtrErr, 0, 1
		} else {
			c, size, totalSize = jay.ReadStringAt, 0, 1
		}
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
	case "[]byte":
		if isLast {
			c, size, totalSize = jay.ReadBytesPtrErr, 0, 1
		} else {
			c, size, totalSize = jay.ReadBytesAt, 0, 1
		}

	default:
		log.Printf("no function set for type %s yet in unmarshalFuncs()", f.typ)
		return "", 0, 0
	}

	return strings.TrimPrefix(
		runtime.FuncForPC(reflect.ValueOf(c).Pointer()).Name(),
		pkgPrefix,
	), size, totalSize
}
