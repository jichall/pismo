package mock

import (
	"log"
	"net/http"
)

type MockWriter struct {
	header int
}

func (mw *MockWriter) Header() http.Header { return http.Header{} }
func (mw *MockWriter) Write(b []byte) (int, error) {
	log.Printf("%s", b)
	return len(b), nil
}

func (mw *MockWriter) WriteHeader(code int) {
	mw.header = code
}
