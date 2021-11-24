package main

import (
	"io"
	"os"

	"github.com/rmatsuoka/valuefs"
)

func main() {
	S := struct {
		Map map[string]string
	}{
		Map: map[string]string{"hello": "こんにちは", "world": "世界"},
	}

	fsys := valuefs.New(S)
	f, err := fsys.Open("Map/hello")
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, f) // output: "こんにちは"
	f.Close()
}
