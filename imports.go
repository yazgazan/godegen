package godegen

import (
	"bytes"
	"fmt"
	"go/types"

	"golang.org/x/tools/go/packages"
)

type Imports struct {
	targetPkg *types.Package
	m         map[string]string // map[path]identifier

	index int
}

func NewImports(pkg *types.Package) *Imports {
	return &Imports{
		targetPkg: pkg,
		m:         map[string]string{},
	}
}

func (imports *Imports) Add(pkgPath string) string {
	if id, ok := imports.m[pkgPath]; ok {
		return id
	}

	pkg, err := Package(pkgPath, packages.NeedName)
	if err != nil {
		panic(err)
	}

	id := fmt.Sprintf("%sGodegen%x", pkg.Name, imports.index)
	imports.m[pkgPath] = id

	return id
}

func (imports *Imports) AddWithName(pkgPath, pkgName string) string {
	if id, ok := imports.m[pkgPath]; ok {
		return id
	}

	id := fmt.Sprintf("%sGodegen%x", pkgName, imports.index)
	imports.m[pkgPath] = id
	imports.index++

	return id
}

func (imports *Imports) Type(typ types.Type) string {
	var (
		pkgPath, pkgName string
	)

	elem, prefix := Elem(typ)

	buf := &bytes.Buffer{}
	types.WriteType(buf, elem, func(pkg *types.Package) string {
		if pkg == imports.targetPkg {
			return ""
		}

		pkgPath = pkg.Path()
		pkgName = pkg.Name()
		return ""
	})
	typName := buf.String()
	if pkgPath == "" {
		return prefix + typName
	}

	pkgId := imports.AddWithName(pkgPath, pkgName)
	return prefix + pkgId + "." + typName
}

func (imports *Imports) Statement() string {
	s := "import (\n"
	for pkgPath, pkgId := range imports.m {
		s += fmt.Sprintf("\t%s %q\n", pkgId, pkgPath)
	}
	s += ")"

	return s
}

// func Type(pkgPath string, t types.Type) (path, name string) {
// 	name = types.TypeString(t, func(pkg *types.Package) string {
// 		if pkg.Path() == pkgPath {
// 			return ""
// 		}
// 		return pkg.Name()
// 	})
// 	elemName := types.TypeString(t, func(pkg *types.Package) string {
// 		return ""
// 	})
// 	if slice, ok := t.(*types.Slice); ok {
// 		t, elemName = sliceElem(slice)
// 	}
// 	fullType := types.TypeString(t, func(pkg *types.Package) string {
// 		return pkg.Path()
// 	})
// 	if fullType == name {
// 		return "", name
// 	}
// 	if _, ok := t.(*types.Basic); ok {
// 		return "", name
// 	}

// 	path = strings.TrimSuffix(fullType, "."+elemName)

// 	return path, name
// }
