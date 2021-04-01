// +build release

package main

import (
	"embed"
	"io/fs"
)

// content is our web server content.
//go:embed dist
var content embed.FS

func getFS() fs.FS {
	fsys, err := fs.Sub(content, "dist")
	if err != nil {
		panic(err)
	}
	return fsys
}
