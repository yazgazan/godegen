package godegen

import (
	"go/types"
	"sort"
	"strings"
)

type Struct struct {
	Name string
	typ  *types.Struct

	imports *Imports
}

func NewStruct(imports *Imports, name string, obj types.Object) Struct {
	return Struct{
		Name: name,
		typ:  UnderlyingType(obj).(*types.Struct),

		imports: imports,
	}
}

func (st Struct) Fields() []Field {
	ff := make([]Field, 0, st.typ.NumFields())

	for i := 0; i < st.typ.NumFields(); i++ {
		f := st.typ.Field(i)
		tag := Tag(st.typ.Tag(i))
		if tag.HasFlag("godegen", "skip") {
			continue
		}

		ff = append(ff, Field{
			Name: f.Name(),
			Tag:  tag,
			typ:  f.Type(),

			imports: st.imports,
		})
	}

	sort.Slice(ff, func(i, j int) bool {
		return strings.Compare(ff[i].Name, ff[j].Name) < 0
	})

	return ff
}

type Field struct {
	Name string
	Tag  Tag
	typ  types.Type

	imports *Imports
}

func (f Field) Type() Type {
	return Type{
		Kind: TypeKind(f.typ),
		typ:  f.typ,

		imports: f.imports,
	}
}

func (f Field) TypeString() string {
	return f.imports.Type(f.typ)
}

func (f Field) TypeName() string {
	return UnderlyingTypeName(f.typ)
}
