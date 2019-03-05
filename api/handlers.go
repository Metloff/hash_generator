package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// HndlrNotFound — обработчик, который вызывается для отсутствующего метода.
func (m *manager) hNotFound() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.sendErr(w, r, 404)
	})
}

func (m *manager) hGenerateMD5() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prms := &hashPrms{}
		if ok := m.getPrms(r, prms); !ok {
			m.sendErr(w, r, 401)
			return
		}

		log.Println(prms)
		// hash := fmt.Sprintf("%x", md5.Sum([]byte(*s)))
		// if err := mainPageT.Execute(w, hash); err != nil {
		// 	bugsnag.Notify(err)
		// 	return
		// }
	})
}

func (m *manager) hGetPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prms := &hashPrms{}
		if ok := m.getPrms(r, prms); !ok {
			m.sendErr(w, r, 401)
			return
		}

		log.Println(prms)
		// hash := fmt.Sprintf("%x", md5.Sum([]byte(*s)))
		// if err := mainPageT.Execute(w, hash); err != nil {
		// 	bugsnag.Notify(err)
		// 	return
		// }
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

	if err := json.Unmarshal(body, prms); err != nil {
		// Примеры возможных ошибок:
		// - тип параметра не соответствует запрошенному типу (json: cannot unmarshal
		// number into Go value of type string);
		// - размер JSON'а превышает оговоренный размер (unexpected end of JSON input).
		log.Println(err, map[string]interface{}{
			"uri":  r.RequestURI,
			"body": string(body),
		})
		return false
	}

	return true
}
