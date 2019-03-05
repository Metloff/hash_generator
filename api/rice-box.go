package api

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file3 := &embedded.EmbeddedFile{
		Filename:    "html/error.html",
		FileModTime: time.Unix(1551811665, 0),

		Content: string("<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Strict//EN\"\n\"http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd\">\n<html xmlns=\"http://www.w3.org/1999/xhtml\" xml:lang=\"en\">\n\n<head>\n    <meta charset=\"utf-8\">\n    <link rel=\"icon\" href=\"/icons/fav.png\">\n    <link href=\"https://fonts.googleapis.com/css?family=Roboto+Mono\" rel=\"stylesheet\">\n    <link href=\"https://fonts.googleapis.com/css?family=Montserrat:400,700\" rel=\"stylesheet\"> \n    <link href=\"/a/assets/img/favicon.png\" rel=\"shortcut icon\" type=\"image/png\">\n    <style>\n        p {\n            font-family: 'Roboto Mono', monospace;\n            margin-bottom: 23px;\n            margin-left: 30px;\n            text-align: center;\n        }\n        .four {\n            font-size: 4em;\n            text-align: center;\n        }\n    </style>\n    <title>{{.ErrorCode}}</title>\n</head>\n\n<body>\n    <p class=\"four\">{{.ErrorCode}}</p>\n    <p>\n        <a href=\"/\">Главная страница</a>\n    </p>\n</body>\n\n</html>"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    "html/index.html",
		FileModTime: time.Unix(1551811640, 0),

		Content: string("<!DOCTYPE html>\n<html>\n  <head>\n    <meta charset=\"utf-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    <title>Hash Generator</title>\n    <link rel=\"stylesheet\" href=\"https://cdnjs.cloudflare.com/ajax/libs/bulma/0.7.4/css/bulma.min.css\">\n    <link href=\"https://fonts.googleapis.com/css?family=Montserrat:400,700\" rel=\"stylesheet\"> \n    <script defer src=\"https://use.fontawesome.com/releases/v5.3.1/js/all.js\"></script>\n  </head>\n  <body style=\"font-family: 'Montserrat'\">\n  <section class=\"section\">\n    <div class=\"container\">\n      <h1 class=\"title\">\n        MD5 Hash Generator\n      </h1>\n      <form action=\"/\" method=\"POST\">\n        <div class=\"field\">\n            <label class=\"label\">Word</label>\n            <div class=\"control has-icons-left has-icons-right\">\n              <input class=\"input is-success is-medium\" type=\"text\" name=\"word\" placeholder=\"Your text\" value={{ .Word }}>\n              <span class=\"icon is-small is-left\">\n                <i class=\"fas fa-pen-square\"></i>\n              </span>\n            </div>\n        </div>\n\n        <div class=\"field\">\n            <label class=\"label\">Salt</label>\n            <div class=\"control has-icons-left has-icons-right\">\n              <input class=\"input is-success is-medium\" type=\"text\" name=\"salt\" placeholder=\"Your salt\" value={{ .Salt }}>\n              <span class=\"icon is-small is-left\">\n                <i class=\"fas fa-fire\"></i>\n              </span>\n            </div>\n        </div>\n\n        <div class=\"field is-grouped\">\n            <div class=\"control\">\n                <button class=\"button is-link\" type=\"submit\" value=\"Submit\">Submit</button>\n            </div>\n        </div>\n          \n      </form>\n\n        {{ if .Present }}\n            <div class=\"box\" style=\"text-align:center; margin-top: 35px;\">\n                <h2>{{.Hash}}</h3>\n            </div>\n        {{ end }}\n    </div>\n  </section>\n\n\n  </body>\n</html>"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1551734758, 0),
		ChildFiles: []*embedded.EmbeddedFile{},
	}
	dir2 := &embedded.EmbeddedDir{
		Filename:   "html",
		DirModTime: time.Unix(1551773251, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file3, // "html/error.html"
			file4, // "html/index.html"

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
			"html/index.html": file4,
		},
	})
}
