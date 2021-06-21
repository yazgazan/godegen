package godegen

import (
	"bytes"
	"fmt"
	"go/types"
)

type Kind string

const (
	KindInterface Kind = "interface"
	KindStruct    Kind = "struct"
)

func UnderlyingKind(obj types.Object) Kind {
	return TypeKind(UnderlyingType(obj))
}

func TypeKind(typ types.Type) Kind {
	switch t := Underlying(typ).(type) {
	default:
		return Kind(fmt.Sprintf("<unknown type %T>", t))
	case *types.Interface:
		return KindInterface
	case *types.Struct:
		return KindStruct
	}
}

func UnderlyingType(obj types.Object) types.Type {
	return Underlying(obj.Type())
}

func Underlying(typ types.Type) types.Type {
	switch t := typ.(type) {
	default:
		return typ
	case *types.Named:
		return Underlying(t.Underlying())
	}
}

func Elem(typ types.Type) (ret types.Type, prefix string) {
	switch t := typ.(type) {
	default:
		return typ, ""
	case *types.Slice:
		ret, prefix = Elem(t.Elem())
		return ret, "[]" + prefix
	case *types.Pointer:
		ret, prefix = Elem(t.Elem())
		return ret, "*" + prefix
	}
}

func UnderlyingTypeName(typ types.Type) string {
	elem, _ := Elem(typ)

	buf := &bytes.Buffer{}
	types.WriteType(buf, elem, func(pkg *types.Package) string {
		return ""
	})
	typName := buf.String()

	return typName
}
