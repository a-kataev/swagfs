package static

import (
	"bytes"
	"io"
	"io/fs"
	"time"
)

type info struct {
	name string
	size int64
}

var _ fs.FileInfo = (*info)(nil)

func (*info) IsDir() bool {
	return false
}

func (*info) ModTime() time.Time {
	return time.Time{}
}

func (*info) Mode() fs.FileMode {
	return fs.ModePerm
}

func (i *info) Name() string {
	return i.name
}

func (i *info) Size() int64 {
	return i.size
}

func (*info) Sys() any {
	return nil
}

type file struct {
	info   *info
	reader *bytes.Reader
}

var (
	_ fs.File       = (*file)(nil)
	_ io.ReadCloser = (*file)(nil)
)

func (f *file) Close() error {
	return nil
}

func (f *file) Read(b []byte) (int, error) {
	return f.reader.Read(b)
}

func (f *file) Seek(o int64, w int) (int64, error) {
	return f.reader.Seek(o, w)
}

func (f *file) Stat() (fs.FileInfo, error) {
	return f.info, nil
}

type AddFunc = func(*Static)

func File(name string, data []byte) AddFunc {
	return func(f *Static) {
		if data == nil {
			data = []byte{}
		}

		f.files[name] = data
	}
}

type Static struct {
	files map[string][]byte
}

var _ fs.FS = (*Static)(nil)

func FS(files ...AddFunc) *Static {
	mfs := &Static{
		files: make(map[string][]byte, len(files)),
	}

	for _, fn := range files {
		fn(mfs)
	}

	return mfs
}

func (f *Static) Open(name string) (fs.File, error) {
	data, ok := f.files[name]
	if !ok {
		return nil, fs.ErrNotExist
	}

	return &file{
		info: &info{
			name: name,
			size: int64(len(data)),
		},
		reader: bytes.NewReader(data),
	}, nil
}
