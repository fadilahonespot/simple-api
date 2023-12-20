package mocks

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type MockHeader struct {
	Key   string
	Value string
}

type DataValidator struct {
	ValidatorData *validator.Validate
}

func (cv *DataValidator) Validate(i interface{}) error {
	return cv.ValidatorData.Struct(i)
}

func MockEcho(method, path string, headers []MockHeader, body interface{}) (c echo.Context, rec *httptest.ResponseRecorder) {
	e := echo.New()

	e.Validator = &DataValidator{ValidatorData: validator.New()}

	rec = httptest.NewRecorder()
	req := httptest.NewRequest(method, path, nil)
	if body != nil {
		byteBody, _ := json.Marshal(body)
		req = httptest.NewRequest(method, path, bytes.NewReader(byteBody))
	}

	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	if headers != nil {
		for _, header := range headers {
			req.Header[header.Key] = []string{header.Value}
		}
	}

	c = e.NewContext(req, rec)
	return
}

type mockReader struct {
	StubRead func() (n int, err error)
}

func NewErrReader() *mockReader {
	return new(mockReader)
}

func (e *mockReader) Read(p []byte) (n int, err error) {
	return e.StubRead()
}
