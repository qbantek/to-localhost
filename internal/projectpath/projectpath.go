package projectpath

import (
	"path/filepath"
	"runtime"
)

var (
	_, file, _, _ = runtime.Caller(0)

	// RootPath is the root path of the project.
	RootPath = filepath.Join(filepath.Dir(file), "../..")
)
