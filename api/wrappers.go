package api

import (
	"errors"
	"log"
	"net/http"
	"runtime/debug"

	bugsnag "github.com/bugsnag/bugsnag-go"
)

// wrapRecover отправляет ошибку пользователю в формате JSON, в случае возникновения паники.
func (m *manager) wrapRecover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			rec := recover()
			if rec != nil {
				switch t := rec.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("Panic with unknown error")
				}

				log.Println(err, map[string]interface{}{
					"stacktrace": string(debug.Stack()),
					"uri":        r.RequestURI,
				})
				bugsnag.Notify(err)
				m.sendErr(w, r, 500)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
