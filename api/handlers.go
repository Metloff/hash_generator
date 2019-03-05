package api

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	bugsnag "github.com/bugsnag/bugsnag-go"
	"github.com/pasztorpisti/qs"
)

// HndlrNotFound — обработчик, который вызывается для отсутствующего метода.
func (m *manager) hNotFound() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.sendErr(w, r, 404)
	})
}

func (m *manager) hGenerateHash() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Получаем параметры.
		prms := &hashPrms{}
		if ok := m.getPrms(r, prms); !ok {
			m.sendErr(w, r, 401)
			return
		}

		// Готовим фразу.
		s := fmt.Sprintf("%v+%v", prms.Word, prms.Salt)

		// Берем хэш.
		hash := fmt.Sprintf("%x", md5.Sum([]byte(s)))

		// Записываем в структуру.
		resp := hashResponse{
			Hash:    hash,
			Present: true,
			Word:    prms.Word,
			Salt:    prms.Salt,
		}

		if err := indexT.Execute(w, resp); err != nil {
			bugsnag.Notify(err)
			return
		}
	})
}

func (m *manager) hGetPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// resp := hashResponse
		if err := indexT.Execute(w, hashResponse{}); err != nil {
			bugsnag.Notify(err)
			return
		}
	})
}

// getPrms парсит входящие в методы API параметры и валидирует их.
func (m *manager) getPrms(r *http.Request, prms interface{}) bool {
	// Размер Body не может быть больше 1МБ (1048576 байт).
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Println(err, map[string]interface{}{"uri": r.RequestURI})
		return false
	}

	if err := r.Body.Close(); err != nil {
		log.Println(err, map[string]interface{}{"uri": r.RequestURI})
		return false
	}

	switch r.Header.Get("Content-Type") {
	case "application/x-www-form-urlencoded",
		"application/x-www-form-urlencoded; charset=utf-8":
		return (qs.Unmarshal(prms, string(body)) == nil)

	case "application/json", "application/json; charset=utf-8":
		if err := json.Unmarshal(body, prms); err != nil {
			// Примеры возможных ошибок:
			// - тип параметра не соответствует запрошенному типу (json: cannot unmarshal
			// number into Go value of type string);
			// - размер JSON'а превышает оговоренный размер (unexpected end of JSON input).
			bugsnag.Notify(err, bugsnag.MetaData{
				"Error": {
					"Description": "Не удалось распарсить пришедшие параметры.",
					"uri":         r.RequestURI,
					"body":        string(body),
				}})
			return false
		}
		return true

	default:
		return true
	}
}
