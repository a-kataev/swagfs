package swagfs

import (
	"io/fs"

	"github.com/a-kataev/swagfs/files"
	"github.com/a-kataev/swagfs/tmpl"
)

type swagFS struct {
	root, override fs.FS
}

var _ fs.FS = (*swagFS)(nil)

func FS(templates ...tmpl.File) fs.FS {
	f := &swagFS{
		root: files.DistFS(),
	}

	override, err := tmpl.FS(templates...)
	if err == nil {
		f.override = override
	}

	return f
}
func (f *swagFS) Open(name string) (fs.File, error) {
	if f.override != nil {
		if file, err := f.override.Open(name); err == nil {
			return file, nil
		}
	}

	return f.root.Open(name)
}
