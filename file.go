package valuefs

import (
	"bytes"
	"fmt"
	"io/fs"
	"reflect"
)

type file struct {
	path string
	v    reflect.Value
	buf  *bytes.Buffer
}

func newFile(path string, v reflect.Value) *file {
	rv := recElem(v)
	buf := new(bytes.Buffer)
	if rv.IsValid() {
		buf.WriteString(fmt.Sprint(rv))
	}
	return &file{
		path: path,
		v:    v,
		buf:  buf,
	}
}

func (f *file) Stat() (fs.FileInfo, error) {
	return fs.FileInfo(newFileInfo(f.path, f.v)), nil
}

func (f *file) Read(b []byte) (int, error) {
	return f.buf.Read(b)
}

func (f *file) Close() error {
	return nil
}
