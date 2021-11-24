# jsonfs

jsonfs provides shell-like interface to manipulate JSON as a file system.

## usage
```console
$ cat example.json
{
    "glossary": {
        "title": "example glossary",
		"GlossDiv": {
            "title": "S",
			"GlossList": {
                "GlossEntry": {
                    "ID": "SGML",
					"SortAs": "SGML",
					"GlossTerm": "Standard Generalized Markup Language",
					"Acronym": "SGML",
					"Abbrev": "ISO 8879:1986",
					"GlossDef": {
                        "para": "A meta-markup language, used to create markup languages such as DocBook.",
						"GlossSeeAlso": ["GML", "XML"]
                    },
					"GlossSee": "markup"
                }
            }
        }
    }
}
$ jsonfs example.json
jsonfs % ls
widget
jsonfs % ls -l widget
-r--r--r-- 0 Nov 24 23:27 debug
dr-xr-xr-x 0 Nov 24 23:27 image
dr-xr-xr-x 0 Nov 24 23:27 text
dr-xr-xr-x 0 Nov 24 23:27 window
jsonfs % read widget/image/src
Images/Sun.pngjsonfs %
jsonfs % walk widget/window
widget/window
widget/window/height
widget/window/name
widget/window/title
widget/window/width
jsonfs % exit
$ 
```
example.json is from <www.json.org/example.html>.
