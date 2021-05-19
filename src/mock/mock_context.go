package mock

import (
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/labstack/echo"
)

type MockContext struct{}
type MockReaderCloser struct{}

func (mrc MockReaderCloser) Read(p []byte) (int, error) { return len(p), nil }
func (mrc MockReaderCloser) Close() error               { return nil }

func (mc *MockContext) Request() *http.Request {

	request := &http.Request{}

	buffRequest := []byte("{\"document_number\": \"12345678911\"}")

	request.Body = MockReaderCloser{}
	request.Body.Read(buffRequest)

	return request
}

func (mc *MockContext) SetRequest(r *http.Request) {}

func (mc *MockContext) Response() *echo.Response {

	response := &echo.Response{}
	response.Writer = &MockWriter{}

	return response
}

func (mc *MockContext) IsTLS() bool { return false }

func (mc *MockContext) IsWebSocket() bool                                       { return false }
func (mc *MockContext) Scheme() string                                          { return "" }
func (mc *MockContext) RealIP() string                                          { return "" }
func (mc *MockContext) Path() string                                            { return "" }
func (mc *MockContext) SetPath(p string)                                        {}
func (mc *MockContext) Param(name string) string                                { return "" }
func (mc *MockContext) ParamNames() []string                                    { return []string{""} }
func (mc *MockContext) SetParamNames(names ...string)                           {}
func (mc *MockContext) ParamValues() []string                                   { return []string{""} }
func (mc *MockContext) SetParamValues(values ...string)                         {}
func (mc *MockContext) QueryParam(name string) string                           { return "" }
func (mc *MockContext) QueryParams() url.Values                                 { return url.Values{} }
func (mc *MockContext) QueryString() string                                     { return "" }
func (mc *MockContext) FormValue(name string) string                            { return "" }
func (mc *MockContext) FormParams() (url.Values, error)                         { return url.Values{}, nil }
func (mc *MockContext) FormFile(name string) (*multipart.FileHeader, error)     { return nil, nil }
func (mc *MockContext) MultipartForm() (*multipart.Form, error)                 { return nil, nil }
func (mc *MockContext) Cookie(name string) (*http.Cookie, error)                { return nil, nil }
func (mc *MockContext) SetCookie(cookie *http.Cookie)                           {}
func (mc *MockContext) Cookies() []*http.Cookie                                 { return nil }
func (mc *MockContext) Get(key string) interface{}                              { return nil }
func (mc *MockContext) Set(key string, val interface{})                         {}
func (mc *MockContext) Bind(i interface{}) error                                { return nil }
func (mc *MockContext) Validate(i interface{}) error                            { return nil }
func (mc *MockContext) Render(code int, name string, data interface{}) error    { return nil }
func (mc *MockContext) HTML(code int, html string) error                        { return nil }
func (mc *MockContext) HTMLBlob(code int, b []byte) error                       { return nil }
func (mc *MockContext) String(code int, s string) error                         { return nil }
func (mc *MockContext) JSON(code int, i interface{}) error                      { return nil }
func (mc *MockContext) JSONPretty(code int, i interface{}, indent string) error { return nil }
func (mc *MockContext) JSONBlob(code int, b []byte) error                       { return nil }
func (mc *MockContext) JSONP(code int, callback string, i interface{}) error    { return nil }
func (mc *MockContext) JSONPBlob(code int, callback string, b []byte) error     { return nil }
func (mc *MockContext) XML(code int, i interface{}) error                       { return nil }
func (mc *MockContext) XMLPretty(code int, i interface{}, indent string) error  { return nil }
func (mc *MockContext) XMLBlob(code int, b []byte) error                        { return nil }
func (mc *MockContext) Blob(code int, contentType string, b []byte) error       { return nil }
func (mc *MockContext) Stream(code int, contentType string, r io.Reader) error  { return nil }
func (mc *MockContext) File(file string) error                                  { return nil }
func (mc *MockContext) Attachment(file string, name string) error               { return nil }
func (mc *MockContext) Inline(file string, name string) error                   { return nil }
func (mc *MockContext) NoContent(code int) error                                { return nil }
func (mc *MockContext) Redirect(code int, url string) error                     { return nil }
func (mc *MockContext) Error(err error)                                         {}
func (mc *MockContext) Handler() echo.HandlerFunc                               { return func(c echo.Context) error { return nil } }
func (mc *MockContext) SetHandler(h echo.HandlerFunc)                           {}
func (mc *MockContext) Logger() echo.Logger                                     { return echo.New().Logger }
func (mc *MockContext) Echo() *echo.Echo                                        { return nil }
func (mc *MockContext) Reset(r *http.Request, w http.ResponseWriter)            {}
