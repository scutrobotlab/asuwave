//go:build !release

package main

import (
	"io/fs"
	"os"
)

func getFS() fs.FS {
	return os.DirFS("dist")
}
