package api

import (
	"log"
	"net/http"
	"text/template"
	"time"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
)

const (
	readTimeOut  = 600 // время на ответ сервера в секундах.
	writeTimeOut = 600 // время на прочтение данных сервером в секундах.
)

var (
	indexT *template.Template
	errorT *template.Template
)

// Manager — менеджер API.
type Manager interface {
	Listen(addr string) error
}

type manager struct {
	router *mux.Router
}

// route описывает поля обработчика запроса.
// Используется при добавлении путей и обработчиков запросов в мультиплексов mux.
type route struct {
	Method   string
	Path     string
	Name     string
	HndlrGen func() http.Handler
	Wrappers []func(inner http.Handler) http.Handler
}

// NewManager возвращает объект менеджера API.
func NewManager() Manager {
	m := &manager{
		router: mux.NewRouter(),
	}

	// Templates.
	templateBox, err := rice.FindBox("../assets")
	if err != nil {
		log.Fatal(err)
	}

	// get file contents as string
	var errorString, indexString string
	if errorString, err = templateBox.String("html/error.html"); err != nil {
		log.Fatal(err)
	}
	if indexString, err = templateBox.String("html/index.html"); err != nil {
		log.Fatal(err)
	}

	// parse and execute the template
	errorT = template.Must(template.New("md5").Parse(errorString))
	indexT = template.Must(template.New("md5").Parse(indexString))

	// Методы API.
	m.addHandlers([]route{
		{
			Method:   "POST",
			Path:     "/",
			Name:     "GenerateHash",
			HndlrGen: m.hGenerateHash,
			Wrappers: []func(inner http.Handler) http.Handler{m.wrapRecover},
		},
		{
			Method:   "GET",
			Path:     "/",
			Name:     "GetPage",
			HndlrGen: m.hGetPage,
			Wrappers: []func(inner http.Handler) http.Handler{m.wrapRecover},
		},
	})

	m.router.NotFoundHandler = m.hNotFound()
	return m
}

// addHandlers добавляет пути и обработчики запросов в мультиплексор mux.
// wrapRecover и wrapDuration добавляются по умолчанию.
func (m *manager) addHandlers(routes []route) {
	for _, r := range routes {
		h := r.HndlrGen()
		for _, w := range r.Wrappers {
			h = w(h)
		}

		m.router.Methods(r.Method).
			Path(r.Path).
			Name(r.Name).
			Handler(m.wrapRecover(h))
	}
}

// sendErr отправляет клиенту ответ с ошибкой.
func (m *manager) sendErr(w http.ResponseWriter, r *http.Request, code int) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(code)

	errorT.Execute(w, map[string]interface{}{
		"ErrorCode": code,
	})
}

// Listen начинает прослушивание порта и обработки запросов.
func (m *manager) Listen(addr string) error {
	log.Println("Public API started on addr", addr)

	srv := &http.Server{
		Addr:    addr,
		Handler: m.router,

		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	return srv.ListenAndServe()
}
