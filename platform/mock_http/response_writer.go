package mock_http

import (
	"encoding/json"
	"net/http"
)

// need to learn more about syntax
type ResponseWriter struct {
	Written    []byte
	StatusCode int
}

// need to learn more about syntax
func (m *ResponseWriter) Write(w []byte) (int, error) {
	m.Written = w
	return 0, nil
}

// need to learn more about syntax
func (m *ResponseWriter) Header() http.Header {
	return http.Header{}
}

// need to learn more about syntax
func (m *ResponseWriter) WriteHeader(statusCode int) {
	m.StatusCode = statusCode
}

// need to learn more about syntax
func (m *ResponseWriter) GetBodyJSON() map[string]interface{} {
	var v map[string]interface{}
	json.Unmarshal(m.Written, &v)
	return v
}

// need to learn more about syntax
func (m *ResponseWriter) GetBodyJSONArray() []map[string]interface{} {
	var v []map[string]interface{}
	json.Unmarshal(m.Written, &v)
	return v
}

// need to learn more about syntax
func (m *ResponseWriter) GetBodyString() string {
	return string(m.Written)
}
