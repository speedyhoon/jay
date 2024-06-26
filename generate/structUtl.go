package generate

import (
	"bytes"
	"fmt"
	"github.com/speedyhoon/jay"
	"go/ast"
	"go/token"
	"reflect"
	"strconv"
	"strings"
)

const (
	StructTagName = "j"
	IgnoreFlag    = "-"
	tagSymbol     = '`'
)

// getTag returns the value associated with key "j" in the tag string.
func getTag(b *ast.BasicLit) string {
	if b == nil {
		return ""
	}

	if b.Kind != token.STRING {
		return ""
	}

	return strings.TrimSpace(reflect.StructTag(unwrapTagValue(b.Value)).Get(StructTagName))
}

// unwrapTagValue removes the leading and trailing grave (`) if present.
func unwrapTagValue(str string) string {
	if len(str) >= 2 && str[0] == tagSymbol && str[len(str)-1] == tagSymbol {
		return str[1 : len(str)-1]
	}
	return str
}

func isBuiltIn(typ string) bool {
	switch typ {
	case "bool", "byte", "float32", "float64", "int", "int8", "int16", "int32", "int64", "rune", "string", "uint", "uint8", "uint16", "uint32", "uint64":
		return true
	}
	return false
}

func (o Option) isSupportedType(t interface{}, dirList *dirList, pkg string) (f field, ok bool) {
	switch d := t.(type) {
	case *ast.Ident:
		if d.Obj == nil {
			if !isBuiltIn(d.Name) && dirList != nil {
				// Object might be declared in another file in the same package.
				d.Obj = findImportedType((*dirList)[pkg].files, (*dirList)[pkg].pkg, d.Name)
				if d.Obj == nil {
					return f, false
				}
				f, ok = o.isSupportedType(d.Obj, dirList, pkg)
				f.aliasType = d.Name
			}
			f.resolveBuiltinAlias(d.Name)
			f.isFixedLen = o.isLenFixed(f.typ)
			return f, true
		}

		f, ok = o.isSupportedType(d.Obj, dirList, pkg)
		f.aliasType = d.Name

	case nil:
	// Ignore.
	case *ast.SelectorExpr:
		f, ok = o.isSupportedSelector(d, dirList)

	case *ast.Object:
		if d.Kind != ast.Typ || d.Name == "" {
			lg.Println(d)
			return f, false
		}
		f, ok = o.isSupportedType(d.Decl, dirList, pkg)

	//case *ast.StructType:
	// TODO not yet implemented.
	case *ast.ArrayType:
		f, ok = o.isSupportedType(d.Elt, dirList, pkg)
		if !ok || f.isAliasDef {
			// f.isAliasDef prevents types like []float where `type float float32` which can't be easily converted to []float32 in one line.
			// However, `type float = float32`, `type floats = []float32` & `type floats []float32` can be easily converted in one line.
			return f, false
		}
		f.arrayType = f.typ
		f.typ = "[]" + f.typ
		f.arraySize, ok = calcArraySize(d.Len)
		f.isFixedLen = f.isFixedLen && f.arraySize != typeSlice

	case *ast.TypeSpec:
		f, ok = o.isSupportedType(d.Type, dirList, pkg)
		f.isAliasDef = f.isAliasDef || d.Assign == token.NoPos

	default:
		lg.Printf("type %T not expected in Option.isSupportedType()", d)
	}
	return
}

func findImportedType(files []*ast.File, pkg, typName string) *ast.Object {
	// TODO doesn't handle aliased imports
	var found []*ast.Object

	for _, file := range files {
		if file == nil || file.Name == nil || file.Name.Name != pkg || file.Scope == nil {
			continue
		}

		for key, obj := range file.Scope.Objects {
			if obj == nil || obj.Decl == nil {
				continue
			}
			if key == typName {
				found = append(found, obj)
				//return obj
			}
		}
	}

	switch len(found) {
	case 0:
		lg.Println("no imports found:", pkgSelName(pkg, typName))
	case 1:
		return found[0]
	default:
		lg.Printf("%d imports found: %s", len(found), pkgSelName(pkg, typName))
	}
	return nil
}

func (o Option) isSupportedSelector(d *ast.SelectorExpr, dirList *dirList) (f field, ok bool) {
	x, ok := d.X.(*ast.Ident)
	if !ok {
		return
	}

	switch x.Name {
	case "time":
		switch d.Sel.Name {
		case "Time", "Duration":
			f.resolveBuiltinAlias(pkgSelName(x.Name, d.Sel.Name))
			f.pkgReq = x.Name
			f.isFixedLen = true
			return f, true
		}
	}

	obj := findImportedType(dirList.allFiles(), x.Name, d.Sel.Name)
	if obj != nil {
		f, ok = o.isSupportedType(obj, nil, "")
		f.aliasType = d.Sel.Name
		return f, f.typ != ""
	}

	return
}

// resolveBuiltinAlias replaces the built-in alias with the underlining name to reduce the quantity of types to support.
func (f *field) resolveBuiltinAlias(typ string) {
	switch typ {
	case "byte":
		f.typ = "uint8"
	case "rune":
		f.typ = "int32"
	case "time.Duration":
		f.typ = "int64"
		f.aliasType = typ
		f.isAliasDef = true
	default:
		if f.typ == "" {
			f.typ = typ
		}
	}
}

func pkgSelName(pkg, selector string) string {
	return fmt.Sprintf("%s.%s", pkg, selector)
}

func packageName(f *ast.File) string {
	if f != nil && f.Name != nil {
		return f.Name.Name
	}
	return "main"
}

func (o *Option) newFieldArray(arraySize int, arrayType string) (f field) {
	f = field{
		arraySize:  arraySize,
		arrayType:  arrayType,
		typ:        genType(arraySize, arrayType),
		isFixedLen: arraySize >= 0 && o.isLenFixed(arrayType),
	}
	return
}
func genType(arraySize int, typ string) string {
	switch arraySize {
	case typeSlice:
		return fmt.Sprintf("[]%s", typ)
	case typeNotArrayOrSlice:
		return typ
	default:
		return fmt.Sprintf("[%d]%s", arraySize, typ)
	}
}

const (
	typeSlice           = -1
	typeNotArrayOrSlice = 0
	typeArray           = 1
)

// calcArraySize returns -1 for a slice, 0 as invalid & >= 1 for array size.
func calcArraySize(x interface{}) (size int, ok bool) {
	switch d := x.(type) {
	case nil:
		return typeSlice, true
	case *ast.BasicLit:
		if d.Kind != token.INT {
			lg.Println("unhandled token type", d.Kind, d.Value)
			return typeNotArrayOrSlice, false
		}

		u, err := strconv.ParseUint(d.Value, 10, 24)
		if err != nil {
			lg.Println("invalid array size", d.Value)
			return typeNotArrayOrSlice, false
		}
		return int(u), true

	case *ast.ValueSpec:
		if len(d.Values) == 1 {
			return calcArraySize(d.Values[0])
		}
	case *ast.Ident:
		if d.Obj != nil {
			return calcArraySize(d.Obj.Decl)
		}
	default:
		lg.Printf("unhandled case %t in calcArraySize", d)
	}
	return typeNotArrayOrSlice, false
}

func (s *structTyp) hasExportedFields() bool {
	return len(s.fixedLen) >= 1 || len(s.variableLen) >= 1 || len(s.bool) >= 1 || len(s.single) >= 1
}

func (s *structTyp) addExportedFields(names []*ast.Ident, f field) {
	f.structTyp = s
	for m := range names {
		f.name = names[m].Name
		switch f.typ {
		case "bool":
			s.bool = append(s.bool, f)
			continue
		case "byte", "int8", "uint8":
			s.single = append(s.single, f)
			continue
		}
		// TODO add support for adding tiny enums using <= 7 bits

		if f.isFixedLen {
			s.fixedLen = append(s.fixedLen, f)
		} else {
			s.variableLen = append(s.variableLen, f)
		}
	}
}

func isLen(typ string) uint {
	switch typ {
	case "bool", "byte", "int8", "uint8", "string", "int", "uint":
		return 1
	case "int16", "uint16":
		return 2
	case "int32", "rune", "uint32", "float32":
		return 4
	case "int64", "uint64", "float64":
		return 8
	}
	return 0
}

func (o Option) isLenFixed(typ string) bool {
	switch typ {
	case "int":
		return o.FixedIntSize
	case "string":
		return false
	case "uint":
		return o.FixedUintSize
	}
	return true
}

func bufWriteF(b *bytes.Buffer, format string, a ...any) {
	b.WriteString(fmt.Sprintf(format, a...))
}

func (s *structTyp) defineTrackingVars(buf *bytes.Buffer, byteIndex uint) (at, end string) {
	switch len(s.variableLen) {
	case 0:
		return
	case 1:
		at = Utoa(byteIndex)
	default:
		if s.variableLen[0].typ == "[]bool" {
			bufWriteF(buf, "at, end := %d, %[1]d+%s(%s)\n", byteIndex, nameOf(jay.SizeBools, nil), lenVariable(0))
		} else {
			bufWriteF(buf, "at, end := %d, %[1]d+%s\n", byteIndex, multiples(s.variableLen[0], lenVariable(0)))
		}
		at, end = "at", "end"
	}
	return
}

func (s *structTyp) tracking(buf *bytes.Buffer, i int, endVar string, byteIndex uint, varType string) (at, end string) {
	if endVar == "" {
		return Utoa(byteIndex), ""
	}

	if i == len(s.variableLen)-1 {
		return "end", ""
	}
	if i >= 1 {
		if varType == "[]bool" {
			bufWriteF(buf, "at, end = end, end+%s(%s)\n", nameOf(jay.SizeBools, nil), lenVariable(i))
		} else {
			bufWriteF(buf, "at, end = end, end+%s\n", multiples(s.variableLen[i], lenVariable(i)))
		}
	}
	return "at", "end"
}

func lenVariable(index int) string {
	const lengthVarPrefix = "l"
	return lengthVarPrefix + Utoa(uint(index))
}
