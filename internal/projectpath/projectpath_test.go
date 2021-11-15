package projectpath_test

import (
	"os"
	"testing"

	"github.com/qbantek/to-localhost/internal/projectpath"
)

func TestRootPath(t *testing.T) {
	// If the root path is not properly set, this test will fail
	// because it won't be able to find the index.tmpl.html file.
	// caveat: make sure index.tmpl.html is in the /templates folder!
	fileName := projectpath.RootPath + "/templates/index.tmpl.html"
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		t.Errorf("RootPath is not correct: %s", projectpath.RootPath)
	}
}
