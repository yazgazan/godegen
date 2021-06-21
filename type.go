package godegen

import (
	"fmt"
	"go/types"
)

type Type struct {
	Kind Kind
	typ  types.Type

	imports *Imports
}

func (t Type) Struct() Struct {
	if t.Kind != KindStruct {
		panic(fmt.Errorf("%q is not a struct (kind %q)", t.typ.String(), t.Kind))
	}

	return Struct{
		Name: UnderlyingTypeName(t.typ),
		typ:  Underlying(t.typ).(*types.Struct),

		imports: t.imports,
	}
}

func (t Type) Interface() Interface {
	if t.Kind != KindInterface {
		panic(fmt.Errorf("%q is not an interface (kind %q)", t.typ.String(), t.Kind))
	}

	return Interface{
		Name: UnderlyingTypeName(t.typ),
		typ:  Underlying(t.typ).(*types.Interface),

		imports: t.imports,
	}
}
