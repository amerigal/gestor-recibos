package recibo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestGetStatusApi(t *testing.T) {
	server := httptest.NewServer(getRouter())
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.GET("/status").Expect().Status(http.StatusOK).JSON().Object().ContainsKey("Estado").ValueEqual("Estado", "Todo OK!")
}
