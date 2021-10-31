package valuefs

import (
	"io/fs"
	"path"
	"reflect"
	"time"
)

type fileInfo struct {
	base string
	v    reflect.Value
}

func newFileInfo(name string, v reflect.Value) *fileInfo {
	return &fileInfo{base: path.Base(name), v: v}
}

func (info *fileInfo) Name() string {
	return info.base
}

func (info *fileInfo) Size() int64 {
	rv := recElem(info.v)
	if !rv.IsValid() {
		return 0
	}
	return int64(rv.Type().Size())
}

func (info *fileInfo) Type() fs.FileMode {
	if isDir(info.v) {
		return fs.ModeDir
	}
	return 0
}

func (info *fileInfo) Mode() fs.FileMode {
	if isDir(info.v) {
		return fs.ModeDir | 0555
	}
	return 0444
}

func (info *fileInfo) ModTime() time.Time {
	return time.Now()
}

func (info *fileInfo) IsDir() bool {
	return isDir(info.v)
}

func (info *fileInfo) Sys() interface{} {
	return nil
}

func (info *fileInfo) Info() (fs.FileInfo, error) {
	return fs.FileInfo(info), nil
}
