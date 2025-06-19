// Package request предоставляет удобный обёртку для работы с HTTP-запросами.
package request

import "net/http"

// Request оборачивает стандартный http.Request и добавляет дополнительные методы для удобства.
type Request struct {
	Req        *http.Request
	parsedForm bool
}
