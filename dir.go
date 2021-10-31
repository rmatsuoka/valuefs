package valuefs

import (
	"errors"
	"io/fs"
	"reflect"
)

type dir struct {
	path string
	v    reflect.Value
}

func newDir(path string, v reflect.Value) *dir {
	return &dir{path: path, v: v}
}

func (d *dir) Stat() (fs.FileInfo, error) {
	return fs.FileInfo(newFileInfo(d.path, d.v)), nil
}

func (d *dir) Read(b []byte) (int, error) {
	return 0, &fs.PathError{
		Op:   "read",
		Path: d.path,
		Err:  errors.New("is a directory"),
	}
}

func (d *dir) Close() error {
	return nil
}
