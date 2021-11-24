package valuefs

import (
	"errors"
	"io"
	"io/fs"
	"testing"
)

func TestOpen(t *testing.T) {
	str := "This is a string"
	pstr := &str
	S := struct {
		Bool         bool
		String       string
		Int          int
		Complex      complex64
		Slice        []string
		NilSlice     []string
		Struct       struct{ Str string }
		PtrStruct    *struct{ Str string }
		NilPtrStruct *struct{ Str string }
		Map          map[string]string
		NilMap       map[string]string
		Func         func() int
		Interface    interface{}
		NilInterface interface{}
		Ptr          *string
		NilPtr       *string
		PtrPtrString **string
	}{
		Bool:         true,
		String:       str,
		Int:          41,
		Complex:      -1+2i,
		Slice:        []string{"こんにちは世界", "hello world"},
		NilSlice:     nil,
		Struct:       struct{ Str string }{Str: "File System"},
		PtrStruct:    &struct{ Str string }{Str: "File System"},
		NilPtrStruct: nil,
		Map:          map[string]string{"hello": "こんにちは", "world": "世界"},
		NilMap:       nil,
		Func:         func() int { return 0 },
		Interface:    struct{ Str string }{Str: "Have you ever use Plan 9?"},
		NilInterface: nil,
		Ptr:          &str,
		NilPtr:       nil,
		PtrPtrString: &pstr,
	}

	fsys := New(S)
	files := []struct {
		name    string
		openErr error
		readErr error
	}{
		{"not-existing", fs.ErrNotExist, nil},
		{"Bool", nil, nil},
		{"String", nil, nil},
		{"Int", nil, nil},
		{"Complex", nil, nil},
		{"Slice", nil, errIsDir},
		{"Slice/0", nil, nil},
		{"Slice/1", nil, nil},
		{"Slice/2", fs.ErrNotExist, nil},
		{"Slice/not-existing", fs.ErrNotExist, nil},
		{"NilSlice", nil, errIsDir},
		{"Struct", nil, errIsDir},
		{"Struct/Str", nil, nil},
		{"Struct/not-existing", fs.ErrNotExist, nil},
		{"PtrStruct", nil, errIsDir},
		{"PtrStruct/Str", nil, nil},
		{"NilPtrStruct", nil, nil},
		{"NilPtrStruct/not-existing", fs.ErrNotExist, nil},
		{"Map", nil, errIsDir},
		{"Map/hello", nil, nil},
		{"Map/world", nil, nil},
		{"Map/not-existing", fs.ErrNotExist, nil},
		{"NilMap", nil, errIsDir},
		{"NilMap/not-existing", fs.ErrNotExist, nil},
		{"Func", nil, nil},
		{"Interface", nil, errIsDir},
		{"Interface/Str", nil, nil},
		{"NilInterface", nil, nil},
		{"NilInterface/not-existing", fs.ErrNotExist, nil},
		{"Ptr", nil, nil},
		{"NilPtr", nil, nil},
		{"PtrPtrString", nil, nil},
	}
	for _, file := range files {
		f, err := fsys.Open(file.name)
		if !errors.Is(err, file.openErr) {
			t.Fatal(err)
		}
		if err != nil {
			t.Logf("expected error %s: %v", file.name, err)
			continue
		}

		b, err := io.ReadAll(f)
		if !errors.Is(err, file.readErr) {
			t.Fatal(err)
		}
		if err != nil {
			t.Logf("expected error %s: %v", file.name, err)
			continue
		}
		t.Logf("%s: %q", file.name, b)
	}
}
