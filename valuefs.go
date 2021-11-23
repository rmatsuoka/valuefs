package valuefs

import (
	"errors"
	"fmt"
	"io/fs"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type vfs struct {
	v reflect.Value
}

func recElem(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Interface || v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}

func isDir(v reflect.Value) bool {
	switch recElem(v).Kind() {
	case reflect.Array, reflect.Slice, reflect.Struct, reflect.Map:
		return true
	}
	return false
}

func New(v interface{}) fs.FS {
	return fs.FS(&vfs{reflect.ValueOf(v)})
}

func (fsys *vfs) namev(name string) (reflect.Value, error) {
	if !fs.ValidPath(name) {
		return reflect.Value{}, fs.ErrInvalid
	}
	v := fsys.v
	if name == "." {
		return v, nil
	}

	nodes := strings.Split(name, "/")
	for ; len(nodes) > 0; nodes = nodes[1:] {
		v = recElem(v)
		switch v.Kind() {
		case reflect.Slice, reflect.Array:
			u, err := strconv.ParseUint(nodes[0], 10, 0)
			i := int(u)
			if i >= v.Len() || err != nil {
				return reflect.Value{}, fs.ErrNotExist
			}
			v = v.Index(i)
		case reflect.Struct:
			v = v.FieldByName(nodes[0])
			if !v.IsValid() {
				return reflect.Value{}, fs.ErrNotExist
			}
		case reflect.Map:
			var nv reflect.Value
			for _, key := range v.MapKeys() {
				if fmt.Sprint(key) == nodes[0] {
					nv = v.MapIndex(key)
					break
				}
			}
			if !nv.IsValid() {
				return reflect.Value{}, fs.ErrNotExist
			}
			v = nv
		default:
			return reflect.Value{}, fs.ErrNotExist
		}
	}
	return v, nil
}

func (fsys *vfs) Open(name string) (fs.File, error) {
	v, err := fsys.namev(name)
	if err != nil {
		return nil, &fs.PathError{"open", name, err}
	}
	if isDir(v) {
		return fs.File(newDir(name, v)), nil
	}
	return fs.File(newFile(name, v)), nil
}

func (fsys *vfs) ReadDir(name string) ([]fs.DirEntry, error) {
	v, err := fsys.namev(name)
	if err != nil {
		return nil, &fs.PathError{Op: "readdir", Path: name, Err: err}
	}

	vv := recElem(v)
	var ds []fs.DirEntry
	switch vv.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < vv.Len(); i++ {
			ds = append(ds, newFileInfo(strconv.Itoa(i), vv.Index(i)))
		}
	case reflect.Struct:
		for i := 0; i < vv.NumField(); i++ {
			ds = append(ds, newFileInfo(vv.Type().Field(i).Name, vv.Field(i)))
		}
	case reflect.Map:
		for _, key := range vv.MapKeys() {
			ds = append(ds, newFileInfo(fmt.Sprint(key), vv.MapIndex(key)))
		}
	default:
		return []fs.DirEntry{}, &fs.PathError{
			Op:   "readdir",
			Path: name,
			Err:  errors.New("not a directory"),
		}
	}
	sort.Slice(ds, func(i, j int) bool { return ds[i].Name() < ds[j].Name() })
	return ds, nil
}
