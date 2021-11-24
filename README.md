# valuefs

Package valuefs provides a function to manipulate values as a file system.

## usage
This package provides one function `New`.
```Go
func New(v interface{}) fs.FS
```
`New` returns fs.FS which provides an interface to value v.

## example
```Go
        S := struct{
                Map map[string]string
        } {
                Map: map[string]string {"hello":"こんにちは", "world":"世界"},
        }

        fsys := valuefs.New(S)
        f, err := fsys.Open("Map/hello")
        if err != nil {
                panic(err)
        }
        io.Copy(os.Stdout, f) // output: "こんにちは"
        f.Close()
```
