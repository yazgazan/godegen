package godegen

import (
	"fmt"
	"go/build"
	"path/filepath"

	"golang.org/x/tools/go/packages"
)

type cachedPackage struct {
	Mode    packages.LoadMode
	Package *packages.Package
}

var packageCache = map[string]cachedPackage{}

func Package(path string, flags ...packages.LoadMode) (*packages.Package, error) {
	var mode packages.LoadMode

	for _, flag := range flags {
		mode |= flag
	}

	cachePath := path
	if build.IsLocalImport(path) {
		// When loading a package using a local path (not fully qualified, i.e '.', '..', etc)
		// the absolute path should be used for caching
		cachePath, _ = filepath.Abs(path)
	}

	if pkg, ok := packageCache[cachePath]; ok && pkg.Mode|mode == pkg.Mode { // if pkg.mode is missing flags, `pkg.mod | mode` will have more bits set than pkg.mode
		return pkg.Package, nil
	} else if ok {
		mode |= pkg.Mode
	}

	pkgs, err := packages.Load(&packages.Config{Mode: mode}, path)
	if err != nil {
		return nil, err
	}
	if len(pkgs) == 0 {
		return nil, fmt.Errorf("no packages found for %q", path)
	}
	if len(pkgs) > 1 {
		return nil, fmt.Errorf("too many packages for %q", path)
	}
	_ = packages.PrintErrors(pkgs)

	packageCache[cachePath] = cachedPackage{
		Mode:    mode,
		Package: pkgs[0],
	}
	return pkgs[0], nil
}
