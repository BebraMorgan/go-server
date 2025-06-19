// Package request предоставляет удобный обёртку для работы с HTTP-запросами.
package request

import (
	"encoding/json"
	"maps"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

// InitRequest создаёт новый экземпляр Request на основе стандартного http.Request.
func InitRequest(r *http.Request) *Request {
	return &Request{Req: r}
}

// All возвращает все параметры запроса из URL и тела формы в виде url.Values.
func (r *Request) All() url.Values {
	r.parseForm()
	values := url.Values{}

	maps.Copy(values, r.Req.URL.Query())
	maps.Copy(values, r.Req.Form)

	return values
}

// Query возвращает значение параметра query с ключом key,
// или defaultValue, если параметр отсутствует.
func (r *Request) Query(key string, defaultValue string) string {
	val := r.Req.URL.Query().Get(key)
	if val == "" {
		return defaultValue
	}
	return val
}

// Url возвращает URI запроса (путь + query string).
func (r *Request) Url() string {
	return r.Req.URL.RequestURI()
}

// FullUrl возвращает полный URL запроса с протоколом и хостом.
func (r *Request) FullUrl() string {
	scheme := "http"
	if r.Req.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + r.Req.Host + r.Req.URL.RequestURI()
}

// Method возвращает HTTP-метод запроса.
func (r *Request) Method() string {
	return r.Req.Method
}

// Path возвращает путь запроса.
func (r *Request) Path() string {
	return r.Req.URL.Path
}

// Json декодирует JSON-тело запроса в переданную структуру v.
func (r *Request) Json(v any) error {
	defer r.Req.Body.Close()
	decoder := json.NewDecoder(r.Req.Body)
	return decoder.Decode(v)
}

// HasFile проверяет, был ли загружен файл с указанным именем.
func (r *Request) HasFile(name string) bool {
	r.parseMultipartForm()
	if r.Req.MultipartForm == nil {
		return false
	}
	_, ok := r.Req.MultipartForm.File[name]
	return ok
}

// File возвращает первый файл с указанным именем из multipart-формы.
func (r *Request) File(name string) (*multipart.FileHeader, error) {
	r.parseMultipartForm()
	if r.Req.MultipartForm == nil {
		return nil, http.ErrMissingFile
	}
	files := r.Req.MultipartForm.File[name]
	if len(files) == 0 {
		return nil, http.ErrMissingFile
	}
	return files[0], nil
}

// Header возвращает значение заголовка с ключом key,
// или defaultValue, если заголовок отсутствует.
func (r *Request) Header(key string, defaultValue string) string {
	val := r.Req.Header.Get(key)
	if val == "" {
		return defaultValue
	}
	return val
}

// HasHeader проверяет наличие заголовка с указанным ключом.
func (r *Request) HasHeader(key string) bool {
	_, ok := r.Req.Header[key]
	return ok
}

// BearerToken извлекает токен Bearer из заголовка Authorization.
func (r *Request) BearerToken() string {
	auth := r.Req.Header.Get("Authorization")
	if auth == "" {
		return ""
	}
	parts := strings.SplitN(auth, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return ""
	}
	return parts[1]
}

// parseForm парсит форму запроса, если она ещё не была распарсена.
func (r *Request) parseForm() {
	if r.parsedForm {
		return
	}
	r.Req.ParseForm()
	r.parsedForm = true
}

// parseMultipartForm парсит multipart-форму с ограничением в 32 МБ.
func (r *Request) parseMultipartForm() {
	if r.Req.MultipartForm != nil {
		return
	}
	r.Req.ParseMultipartForm(32 << 20) // 32 MB
}
