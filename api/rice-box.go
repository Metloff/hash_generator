package api

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file3 := &embedded.EmbeddedFile{
		Filename:    "html/error.html",
		FileModTime: time.Unix(1551744429, 0),

		Content: string("<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Strict//EN\"\n\"http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd\">\n<html xmlns=\"http://www.w3.org/1999/xhtml\" xml:lang=\"en\">\n\n<head>\n    <meta charset=\"utf-8\">\n    <link rel=\"icon\" href=\"/icons/fav.png\">\n    <link href=\"https://fonts.googleapis.com/css?family=Roboto+Mono\" rel=\"stylesheet\">\n    <link href=\"/a/assets/img/favicon.png\" rel=\"shortcut icon\" type=\"image/png\">\n    <style>\n        p {\n            font-family: 'Roboto Mono', monospace;\n            margin-bottom: 23px;\n            margin-left: 30px;\n            text-align: center;\n        }\n        .four {\n            font-size: 4em;\n            text-align: center;\n        }\n    </style>\n    <title>{{.ErrorCode}}</title>\n</head>\n\n<body>\n    <p class=\"four\">{{.ErrorCode}}</p>\n    <p>\n        <a href=\"/\">Главная страница</a>\n    </p>\n</body>\n\n</html>"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1551734758, 0),
		ChildFiles: []*embedded.EmbeddedFile{},
	}
	dir2 := &embedded.EmbeddedDir{
		Filename:   "html",
		DirModTime: time.Unix(1551734770, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file3, // "html/error.html"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{
		dir2, // "html"

	}
	dir2.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`../assets`, &embedded.EmbeddedBox{
		Name: `../assets`,
		Time: time.Unix(1551734758, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"":     dir1,
			"html": dir2,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"html/error.html": file3,
		},
	})
}
