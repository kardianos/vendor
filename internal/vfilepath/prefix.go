package vfilepath

import (
	"path"
	"strings"
)

func HasPrefixDir(path string, prefix string) bool {
	return strings.HasPrefix(makeDirPath(path), makeDirPath(prefix))
}

func makeDirPath(dirpath string) string {
	if dirpath = path.Clean(dirpath); dirpath != "/" {
		dirpath += "/"
	}
	return dirpath
}
