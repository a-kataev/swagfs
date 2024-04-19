package tmpl

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"

	"github.com/a-kataev/swagfs/internal/static"
	"github.com/valyala/fasttemplate"
)

type File interface {
	Filename() string
	Template() string
	TagFunc(io.Writer, string) (int, error)
}

func FS(templates ...File) (fs.FS, error) {
	list := []static.AddFunc{}

	buf := bytes.NewBuffer([]byte{})
	var tmpl *fasttemplate.Template

	for _, file := range templates {
		if file == nil {
			continue
		}

		if tmpl == nil {
			var err error

			tmpl, err = fasttemplate.NewTemplate(file.Template(), "{{", "}}")
			if err != nil {
				return nil, fmt.Errorf("tmpl: template: %v: %v", file.Filename(), err)
			}
		} else {
			err := tmpl.Reset(file.Template(), "{{", "}}")
			if err != nil {
				return nil, fmt.Errorf("tmpl: reset: %v: %v", file.Filename(), err)
			}
		}

		if _, err := tmpl.ExecuteFunc(buf, file.TagFunc); err != nil {
			return nil, fmt.Errorf("tmpl: execute: %v: %v", file.Filename(), err)
		}

		list = append(list,
			static.File(file.Filename(), buf.Bytes()),
		)

		buf.Reset()
	}

	return static.FS(list...), nil
}
