package fs

import (
	"strings"
)

func toSuperPath(dir string, count int) string {
	dir = strings.Replace(dir, "\\", "/", -1)
	dirList := strings.Split(dir, "/")
	return strings.Join(dirList[:len(dirList)-count], "/")
}