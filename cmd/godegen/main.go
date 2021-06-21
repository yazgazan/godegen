package main

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/Pimmr/rig"
	"github.com/yazgazan/godegen"
	"golang.org/x/tools/go/packages"
)

type Configuration struct {
	Target   string `flag:",require"`
	Template string `flag:"tpl,require"`

	OutputFile string `flag:"out,require"`

	Args RigMap
}

var targetSelectorRegexp = regexp.MustCompile("^(\"([^\"]+)\"\\.)?(.+)$")

func main() {
	c := Configuration{
		Args: RigMap{},
	}

	err := rig.ParseStruct(&c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(2)
	}

	outPkgPath := filepath.Dir(c.OutputFile)

	outPkg, err := godegen.Package(outPkgPath, packages.NeedTypes, packages.NeedImports, packages.NeedName, packages.NeedModule)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading destination package: %v\n", err)
		os.Exit(1)
	}

	targetPkg := outPkg

	targetParts := targetSelectorRegexp.FindAllStringSubmatch(c.Target, -1)
	if targetParts[0][2] != "" {
		// external
		targetPkg, err = godegen.Package(targetParts[0][2], packages.NeedTypes, packages.NeedImports, packages.NeedName, packages.NeedModule)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading target package: %v\n", err)
			os.Exit(1)
		}
	}
	targetName := targetParts[0][3]

	tpl := template.New("main")

	imports := godegen.NewImports(targetPkg.Types)

	var importsPlaceholder string
	tpl = tpl.Funcs(template.FuncMap{
		"Imports": func() string {
			if importsPlaceholder != "" {
				panic(errors.New("Imports() already called"))
			}

			importsPlaceholder = strconv.FormatUint(rand.Uint64(), 36)
			return importsPlaceholder
		},
		"Export": func(s string) string {
			return strings.Title(s)
		},
		"Args": func() RigMap {
			return c.Args
		},
	})

	tpl, err = tpl.ParseFiles(c.Template)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading template: %v\n", err)
		os.Exit(1)
	}
	tpl = tpl.Templates()[0]

	targetObj := targetPkg.Types.Scope().Lookup(targetName)

	buf := &bytes.Buffer{}
	switch godegen.UnderlyingKind(targetObj) {
	default:
		fmt.Fprintln(os.Stderr, "Error: unsupported kind")
		os.Exit(1)
	case godegen.KindInterface:
		err = tpl.Execute(buf, godegen.NewInterface(imports, targetName, targetObj))
	case godegen.KindStruct:
		err = tpl.Execute(buf, godegen.NewStruct(imports, targetName, targetObj))
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing template: %v\n", err)
		os.Exit(1)
	}

	b := buf.Bytes()
	b = bytes.Replace(b, []byte(importsPlaceholder), []byte(imports.Statement()), 1)

	f, err := os.Create(c.OutputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	_, err = f.Write(b)
	f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	gofmt := exec.Command("gofmt", "-s", "-w", c.OutputFile)
	gofmt.Stdout = os.Stdout
	gofmt.Stderr = os.Stderr
	err = gofmt.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running `gofmt -s -w %q`: %v\n", c.OutputFile, err)
	}

	goimports := exec.Command("goimports", "-w", "-local", targetPkg.Module.Path, c.OutputFile)
	goimports.Stdout = os.Stdout
	goimports.Stderr = os.Stderr
	err = goimports.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running `goimports -w -local %q %q`: %v\n", targetPkg.Module.Path, c.OutputFile, err)
	}

	fmt.Fprintf(os.Stdout, "wrote %s\n", c.OutputFile)
}

type RigMap map[string]string

func (m RigMap) Set(s string) error {
	ss := strings.Split(s, ";")
	for _, v := range ss {
		if strings.TrimSpace(v) == "" {
			continue
		}

		err := m.set(v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m RigMap) set(s string) error {
	ss := strings.SplitN(s, "=", 2)
	if len(ss) == 1 {
		return fmt.Errorf("malformed key=value pair %q", s)
	}

	m[strings.TrimSpace(ss[0])] = strings.TrimSpace(ss[1])

	return nil
}

func (m RigMap) String() string {
	ss := make([]string, 0, len(m))
	for k, v := range m {
		ss = append(ss, k+"="+v)
	}

	sort.Strings(ss)

	return strings.Join(ss, ";")
}

func (m RigMap) Has(key string) bool {
	_, ok := m[key]
	return ok
}

func (m RigMap) Get(key string) string {
	return m[key]
}

func (m RigMap) GetOrDefault(key, defaultValue string) string {
	v, ok := m[key]
	if !ok {
		return defaultValue
	}

	return v
}
