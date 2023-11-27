package jay

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
)

var ErrNoneExported = errors.New("no exported struct fields found")

func (s *Struct) GenerateFuncs(b *bytes.Buffer, o Option) {
	s.MakeMarshalJ(b)
	s.MakeMarshalJX(b, o)
	s.MakeMarshalJTo(o, b)
	s.MakeSize(o, b)
	s.MakeUnmarshal(o, b)
}

func GenerateFile(pkg string, s []Struct, option Option) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	for i := range s {
		s[i].GenerateFuncs(buf, option)
	}

	src := buf.Bytes()
	if len(src) == 0 {
		return nil, ErrNoneExported
	}

	if pkg == "" {
		pkg = "main"
	}

	src = []byte(fmt.Sprintf(`// Code generated by Jay; DO NOT EDIT.
package %s
import "%s"
%s`, pkg, reflect.TypeOf(option).PkgPath(), src))

	return src, nil
}
