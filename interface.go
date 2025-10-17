package godegen

import (
	"fmt"
	"go/types"
	"sort"
	"strings"
)

type Interface struct {
	Name string
	typ  *types.Interface

	imports *Imports
}

func NewInterface(imports *Imports, name string, obj types.Object) Interface {
	return Interface{
		Name: name,
		typ:  UnderlyingType(obj).(*types.Interface),

		imports: imports,
	}
}

func (iface Interface) Methods() []Method {
	mm := make([]Method, 0, iface.typ.NumMethods())

	for i := 0; i < iface.typ.NumMethods(); i++ {
		m := iface.typ.Method(i)
		mm = append(mm, Method{
			Name: m.Name(),
			fn:   m,
			typ:  m.Type().(*types.Signature),

			imports: iface.imports,
		})
	}

	sort.Slice(mm, func(i, j int) bool {
		return strings.Compare(mm[i].Name, mm[j].Name) < 0
	})

	return mm
}

type Method struct {
	Name string
	fn   *types.Func
	typ  *types.Signature

	imports *Imports
}

func (m Method) Args() []Arg {
	params := m.typ.Params()
	aa := make([]Arg, 0, params.Len())

	for i := 0; i < params.Len(); i++ {
		v := params.At(i)
		aa = append(aa, Arg{
			Name: v.Name(),
			typ:  v,

			imports: m.imports,
		})
	}
	if m.typ.Variadic() {
		aa[len(aa)-1].variadic = true
	}

	return aa
}

func (m Method) Arg(name string) (Arg, error) {
	params := m.typ.Params()

	for i := 0; i < params.Len(); i++ {
		v := params.At(i)
		if v.Name() == name {
			return Arg{
				Name: v.Name(),
				typ:  v,

				imports: m.imports,
			}, nil
		}
	}

	return Arg{}, fmt.Errorf("no argument named %q found", name)
}

func (m Method) HasArg(name string) bool {
	params := m.typ.Params()

	for i := 0; i < params.Len(); i++ {
		v := params.At(i)
		if v.Name() == name {
			return true
		}
	}

	return false
}

func (m Method) Returns() []Arg {
	rets := m.typ.Results()
	aa := make([]Arg, 0, rets.Len())

	for i := 0; i < rets.Len(); i++ {
		r := rets.At(i)
		aa = append(aa, Arg{
			Name: r.Name(),
			typ:  r,

			imports: m.imports,
		})
	}

	return aa
}

func (m Method) HasReturns() bool {
	return m.typ.Results().Len() > 0
}

func (m Method) Return(i int) (Arg, error) {
	rets := m.typ.Results()
	if rets.Len() == 0 {
		return Arg{}, fmt.Errorf("no returns for method %q", m.Name)
	}

	r := rets.At(0)
	return Arg{
		Name: r.Name(),
		typ:  r,

		imports: m.imports,
	}, nil
}

type Arg struct {
	Name     string
	typ      *types.Var
	variadic bool

	imports *Imports
}

func (a Arg) TypeString() string {
	return a.imports.Type(a.typ.Type())
}

func (a Arg) TypeCanonical() string {
	return a.imports.TypeCanonical(a.typ.Type())
}

func (a Arg) IsStruct() bool {
	typ := a.typ.Type()
	if p, ok := typ.(*types.Pointer); ok {
		typ = p.Elem()
	}

	_, ok := Underlying(typ).(*types.Struct)

	return ok
}

func (a Arg) IsInterface() bool {
	typ := a.typ.Type()

	_, ok := Underlying(typ).(*types.Interface)

	return ok
}

func (a Arg) Struct() Struct {
	typ := a.typ.Type()
	if p, ok := typ.(*types.Pointer); ok {
		typ = p.Elem()
	}

	named, _ := typ.(*types.Named)
	return Struct{
		Name:  a.Name,
		typ:   Underlying(typ).(*types.Struct),
		named: named,

		imports: a.imports,
	}
}

func (a Arg) Interface() Interface {
	typ := a.typ.Type()

	return Interface{
		Name: a.Name,
		typ:  Underlying(typ).(*types.Interface),

		imports: a.imports,
	}
}
