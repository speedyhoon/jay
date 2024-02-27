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

var (
	ErrNoSource     = errors.New("no filename or source provided")
	ErrNoneExported = errors.New("no exported struct fields found")
)

func makeFile(pkg string, s []structTyp, option Option) ([]byte, error) {
	mergeEmbeddedStructs(s)

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

	src = []byte(fmt.Sprintf(`// Code generated by Jay; DO NOT EDIT. Tool documentation available at: https://%s%s

package %s
import "%[1]s%[2]s"
%[4]s`, pkgPrefix, pkgName, pkg, src))

	return GoFormat(src)
}

func mergeEmbeddedStructs(structs []structTyp) {
	for i := range structs {
		var structNames []string
		structs[i].fixedLen.mergeFields(structs, structNames, i)
		structs[i].variableLen.mergeFields(structs, structNames, i)
		structs[i].single.mergeFields(structs, structNames, i)
		structs[i].bool.mergeFields(structs, structNames, i)
	}
}

func (f *fieldList) mergeFields(structs []structTyp, structNames []string, structIndex int) {
	for i := 0; i < len(*f); i++ {
		if (*f)[i].typ != "struct" || !isNew(&structNames, (*f)[i].name) {
			continue
		}

		embedded := findStruct(structs, (*f)[i])
		structs[structIndex].join(embedded, (*f)[i].name)

		// Remove the struct field so its contents are not referenced twice.
		*f = Remove(*f, i)
		i--
	}
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

func findStruct(s []structTyp, f field) *structTyp {
	for i := range s {
		if s[i].name == f.aliasType {
			return &s[i]
		}
	}

	// TODO change this log message to verbose only mode.
	log.Printf("can't find %s %s used as name %s", f.typ, f.aliasType, f.name)
	return nil
}

func (s *structTyp) join(embedded *structTyp, name string) {
	if embedded == nil {
		return
	}

	appendEmbed(&s.bool, name, embedded.bool)
	appendEmbed(&s.single, name, embedded.single)
	appendEmbed(&s.fixedLen, name, embedded.fixedLen)
	appendEmbed(&s.variableLen, name, embedded.variableLen)
}

func appendEmbed(fields *fieldList, embedName string, embedded fieldList) {
	for _, e := range embedded {
		// Change the embedded name, so it is correctly referenced in the code generated.
		e.name = fmt.Sprintf("%s.%s", embedName, e.name)
		*fields = append(*fields, e)
	}
}

func (s *structTyp) makeFuncs(b *bytes.Buffer, o Option) {
	if !ast.IsExported(s.name) || !s.hasExportedFields() {
		return
	}

	if !o.SkipMarshal {
		s.makeMarshal(b, o)
	}
	if !o.SkipUnmarshal {
		s.makeUnmarshal(b, o)
	}
}

// GoFormat nicely formats the generated Go code.
func GoFormat(src []byte) (code []byte, err error) {
	code, err = format.Source(src, format.Options{
		LangVersion: strings.TrimPrefix(runtime.Version(), "go"),
		ExtraRules:  true,
	})
	if err != nil {
		return src, err
	}
	return
}
