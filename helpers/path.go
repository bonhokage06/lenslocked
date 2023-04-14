package helpers

import (
	"os"
	"path/filepath"
	"strings"
)

func CurrentWorkingDirectory() {
	wd, _ := os.Getwd()
	for !strings.HasSuffix(wd, "lenslocked") {
		wd = filepath.Dir(wd)
	}
	os.Chdir(wd)
}
