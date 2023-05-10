package api_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jpoikela/miniurl/api"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPI_AddUrl(t *testing.T) {
	const (
		payload = `{"url": "https://github.com/gourses/miniurl/blob/main/LICENSE"}`
		expectedStatusCode = http.StatusOK
		expectedBody = `{"url": "https://github.com/gourses/miniurl/blob/main/LICENSE", "hash": "testvalue"}`
	)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/url", strings.NewReader(payload))
	responseRecorder := httptest.NewRecorder()

	// TODO: implement
	r := httprouter.New()
	h := &strHandler{str: "testvalue"}
	api.Bind(r, h)
	r.ServeHTTP(responseRecorder, req)

	assert.Equal(t, expectedStatusCode, responseRecorder.Result().StatusCode)
	body, err := io.ReadAll(responseRecorder.Result().Body)
	require.NoError(t, err)
	assert.JSONEq(t, expectedBody, string(body))
}

type strHandler struct {
	str string
}

func (h *strHandler) AddUrl(url string) (hash  string, err error) {
	return h.str, nil
}
