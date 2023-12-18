package generate

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"log"
	"mvdan.cc/gofumpt/format"
	"runtime"
	"strings"
)

const (
	pkgName               = "jay"
	pkgPrefix             = "github.com/speedyhoon/"
	goExt                 = ".go"
	DefaultOutputFileName = pkgName + goExt
	IntSize               = 32 << (^uint(0) >> 63) // 32 or 64
)

var ErrNoneExported = errors.New("no exported struct fields found")

func makeFile(pkg string, s []structTyp, option Option) ([]byte, error) {
	for i := range s {
		s[i].mergeEmbeddedStructs(s)
	}

	buf := bytes.NewBuffer(nil)
	for i := range s {
		s[i].makeFuncs(buf, option)
	}

	src := buf.Bytes()
	if len(src) == 0 {
		return nil, ErrNoneExported
	}

	if pkg == "" {
		pkg = "main"
	}

	src = []byte(fmt.Sprintf(`// Code generated by Jay; DO NOT EDIT. Tool documentation available at: https://%s%s.

package %s
import "%[1]s%[2]s"
%[4]s`, pkgPrefix, pkgName, pkg, src))

	// Nicely format the generated Go code.
	var bb []byte
	bb, err := format.Source(src, format.Options{
		LangVersion: strings.TrimPrefix(runtime.Version(), "go"),
		ExtraRules:  true,
	})
	if err != nil {
		//log.Printf("errornous output:\n%s\n", src)
		return src, err
	}
	return bb, nil
}

func (s *structTyp) mergeEmbeddedStructs(list []structTyp) {
	var mergeStructs []string
	for i := 0; i < len(s.fixedLen); i++ {
		if s.fixedLen[i].typ != "struct" {
			continue
		}

		if !isNew(&mergeStructs, s.fixedLen[i].name) {
			continue
		}

		s.mergeFields(&s.fixedLen, &i, list)
	}

	for i := 0; i < len(s.variableLen); i++ {
		if s.variableLen[i].typ != "struct" {
			continue
		}

		if !isNew(&mergeStructs, s.variableLen[i].aliasType) {
			continue
		}

		s.mergeFields(&s.variableLen, &i, list)
	}

	return
}

func (s *structTyp) mergeFields(fields *[]field, index *int, structs []structTyp) {
	aliasType := (*fields)[*index].aliasType
	embedded := findStruct(structs, aliasType)
	if embedded == nil {
		log.Fatalf("can't find %s %s used as name %s", (*fields)[*index].typ, aliasType, (*fields)[*index].name)
	}

	// Remove the struct so its contents are not referenced twice.
	*fields = Remove(*fields, *index)
	*index--

	// Deep copy embedded structType.
	st := structTyp{name: embedded.name}
	st.fixedLen = append(st.fixedLen, embedded.fixedLen...)
	st.variableLen = append(st.variableLen, embedded.variableLen...)
	st.bool = append(st.bool, embedded.bool...)

	s.join(st, aliasType)
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

func findStruct(s []structTyp, name string) *structTyp {
	for i := range s {
		if s[i].name == name {
			return &s[i]
		}
	}
	return nil
}

func (s *structTyp) join(embedded structTyp, name string) {
	s.bool = appendEmbed(s.bool, name, embedded.bool)
	s.fixedLen = appendEmbed(s.fixedLen, name, embedded.fixedLen)
	s.variableLen = appendEmbed(s.variableLen, name, embedded.variableLen)
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

func (s *structTyp) makeFuncs(b *bytes.Buffer, o Option) {
	if !ast.IsExported(s.name) || len(s.bool) == 0 && len(s.variableLen) == 0 && len(s.fixedLen) == 0 {
		return
	}

	s.makeMarshal(b, o)
	s.makeUnmarshal(b, o)
}
