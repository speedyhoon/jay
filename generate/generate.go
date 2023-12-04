package generate

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"log"
)

const (
	pkgName   = "jay"
	pkgPrefix = "github.com/speedyhoon/"
)

var ErrNoneExported = errors.New("no exported struct fields found")

// ErrUnexpectedEOB indicates the end of the byte slice was unexpectedly encountered
// while deserialising a fixed-size block, resulting in an incomplete result.
// var ErrUnexpectedEOB = errors.New("unexpected EOB")

func generateFile(pkg string, s []Struct, option Option) ([]byte, error) {
	for i := range s {
		s[i].shouldMergeEmbeddedStructs(s)
	}

	buf := bytes.NewBuffer(nil)
	for i := range s {
		s[i].generateFuncs(buf, option)
	}

	src := buf.Bytes()
	if len(src) == 0 {
		return nil, ErrNoneExported
	}

	if pkg == "" {
		pkg = "main"
	}

	src = []byte(fmt.Sprintf(`// Code generated by Jay; DO NOT EDIT. https://%s%s
package %s
import "%[1]s%[2]s"
%[4]s`, pkgPrefix, pkgName, pkg, src))

	return src, nil
}

func (s *Struct) shouldMergeEmbeddedStructs(list []Struct) {
	var mergeStructs []string
	for i := 0; i < len(s.fixedLen); i++ {
		if s.fixedLen[i].typ == "struct" {
			var ok bool
			ok = isNew(&mergeStructs, s.fixedLen[i].name)
			if !ok {
				continue
			}

			embedded := findStruct(list, s.fixedLen[i].name)
			if embedded == nil {
				log.Fatalf("can't find %s with name %s", s.fixedLen[i].typ, s.fixedLen[i].name)
			}

			fieldName := s.fixedLen[i].name

			// Remove the struct so its contents are not referenced twice.
			s.fixedLen = Remove(s.fixedLen, i)
			i--

			s.join(embedded, fieldName)
		}
	}

	for i := 0; i < len(s.variableLen); i++ {
		if s.variableLen[i].typ == "struct" {
			var ok bool
			ok = isNew(&mergeStructs, s.variableLen[i].aliasType)
			if !ok {
				continue
			}

			embedded := findStruct(list, s.variableLen[i].aliasType)
			if embedded == nil {
				log.Fatalf("can't find %s %s used as name %s", s.variableLen[i].typ, s.variableLen[i].aliasType, s.variableLen[i].name)
			}

			fieldName := s.variableLen[i].name

			// Remove the struct so its contents are not referenced twice.
			s.variableLen = Remove(s.variableLen, i)
			i--

			s.join(embedded, fieldName)
		}
	}

	return
}

func isNew(a *[]string, s string) bool {
	for i := range *a {
		if (*a)[i] == s {
			return false
		}
	}
	*a = append(*a, s)
	return true
}

func findStruct(s []Struct, name string) *Struct {
	for i := range s {
		if s[i].name == name {
			return &s[i]
		}
	}
	return nil
}

func (s *Struct) join(embedded *Struct, name string) {
	s.bool = appendEmbed(s.bool, name, embedded.bool)
	s.fixedLen = appendEmbed(s.fixedLen, name, embedded.fixedLen)
	s.variableLen = appendEmbed(s.variableLen, name, embedded.variableLen)
	//
	//if len(embedded.bool) >= 1 {
	//	for i := range embedded.bool {
	//		embedded.bool[i].name = fmt.Sprintf("%s.%s", embedded.name, embedded.bool[i].name)
	//	}
	//	s.bool = append(s.bool, embedded.bool...)
	//}
}

func appendEmbed(fields []field, embedName string, embedded []field) []field {
	if len(embedded) >= 1 {
		for i := range embedded {
			// Change the embedded name, so it is correctly referenced in the code generated.
			embedded[i].name = fmt.Sprintf("%s.%s", embedName, embedded[i].name)
		}
		fields = append(fields, embedded...)
	}
	return fields
}

func (s *Struct) generateFuncs(b *bytes.Buffer, o Option) {
	if !ast.IsExported(s.name) || len(s.bool) == 0 && len(s.variableLen) == 0 && len(s.fixedLen) == 0 {
		return
	}

	//s.MakeMarshalJ(b)
	s.MakeMarshalJ(b, o)
	//s.MakeMarshalJTo(o, b)
	//s.MakeSize(b)
	s.MakeUnmarshal(b)
}
