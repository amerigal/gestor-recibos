package recibo

import (
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/gavv/httpexpect/v2"
)

const rutaRecibo = "../../input/reciboCoviran.txt"

var reFecha = regexp.MustCompile(`(0[1-9]|[12][0-9]|3[01])-(0[1-9]|1[0-2])-\d{4}`)

var contenidoRecibo, _ = os.ReadFile(rutaRecibo)

func insertarReciboEjemplo(e *httpexpect.Expect) {
	parametros := map[string]interface{}{
		"usuario":     "amerigal",
		"textoRecibo": string(contenidoRecibo),
	}
	e.POST("/recibos").WithForm(parametros).Expect().Status(http.StatusCreated)
}

func TestGetStatusApi(t *testing.T) {
	server := httptest.NewServer(getRouter())
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.GET("/status").Expect().Status(http.StatusOK).JSON().Object().ContainsKey("Estado").ValueEqual("Estado", "Todo OK!")
}

func TestGetRecibosApi(t *testing.T) {
	server := httptest.NewServer(getRouter())
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	insertarReciboEjemplo(e)

	e.GET("/recibos").Expect().Status(http.StatusOK).JSON().Array().Element(0).Object().ValueEqual("id", 0)
}

func TestInsertarReciboApiValido(t *testing.T) {
	server := httptest.NewServer(getRouter())
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	parametros := map[string]interface{}{
		"usuario":     "amerigal",
		"textoRecibo": string(contenidoRecibo),
	}

	e.POST("/recibos").WithForm(parametros).Expect().Status(http.StatusCreated).JSON().Object().ValueEqual("id", 0).ValueEqual("usuario", "amerigal")
}
func TestInsertarReciboApiInvalido(t *testing.T) {
	server := httptest.NewServer(getRouter())
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	parametros := map[string]interface{}{
		"usuario":     "amerigal",
		"textoRecibo": "Texto inválido Recibo",
	}

	e.POST("/recibos").WithForm(parametros).Expect().Status(http.StatusBadRequest)
}

func TestInsertarReciboApiSinParametros(t *testing.T) {
	server := httptest.NewServer(getRouter())
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.POST("/recibos").Expect().Status(http.StatusBadRequest)
}

func TestGetReciboApiValido(t *testing.T) {
	server := httptest.NewServer(getRouter())
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	insertarReciboEjemplo(e)

	e.GET("/recibos/0").Expect().Status(http.StatusOK).JSON().Object().ValueEqual("id", 0)
}

func TestGetReciboApiNotFound(t *testing.T) {
	server := httptest.NewServer(getRouter())
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.GET("/recibos/3").Expect().Status(http.StatusNotFound)
}

func TestEliminarRecibo(t *testing.T) {
	server := httptest.NewServer(getRouter())
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	insertarReciboEjemplo(e)
	e.DELETE("/recibos/0").Expect().Status(http.StatusNoContent)
	e.DELETE("/recibos/0").Expect().Status(http.StatusNotFound)
}

func TestGetArticulosReciboApiValido(t *testing.T) {
	server := httptest.NewServer(getRouter())
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	insertarReciboEjemplo(e)

	e.GET("/recibos/0/articulos").Expect().Status(http.StatusOK).JSON().Array().Element(0).Object().ValueEqual("descripcion", "PAÑUELOS CLASSIC")
}

func TestGetArticulosReciboApiNotFound(t *testing.T) {
	server := httptest.NewServer(getRouter())
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.GET("/recibos/3/articulos").Expect().Status(http.StatusNotFound)
}

func TestGetArticuloReciboApiValido(t *testing.T) {
	server := httptest.NewServer(getRouter())
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	insertarReciboEjemplo(e)

	e.GET("/recibos/0/articulos/2").Expect().Status(http.StatusOK).JSON().Object().ValueEqual("descripcion", "ATUN CLARO ACEITE")
}

func TestGetArticuloReciboApiNotFound(t *testing.T) {
	server := httptest.NewServer(getRouter())
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.GET("/recibos/3/articulos/2").Expect().Status(http.StatusNotFound)
}

func TestSetTipoArticuloReciboApiValido(t *testing.T) {
	server := httptest.NewServer(getRouter())
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	insertarReciboEjemplo(e)

	parametros := map[string]interface{}{
		"tipo": "yogur",
	}

	e.PATCH("/recibos/0/articulos/7").WithForm(parametros).Expect().Status(http.StatusOK).JSON().Object().ValueEqual("tipo", "yogur")
}

func TestSetTipoArticuloReciboApiNotFound(t *testing.T) {
	server := httptest.NewServer(getRouter())
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	parametros := map[string]interface{}{
		"tipo": "leche",
	}

	e.PATCH("/recibos/3/articulos/2").WithForm(parametros).Expect().Status(http.StatusNotFound)
}

func TestSetTipoArticuloReciboApiSinTipo(t *testing.T) {
	server := httptest.NewServer(getRouter())
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.PATCH("/recibos/3/articulos/2").Expect().Status(http.StatusBadRequest)
}

func TestGetRecuentoApiValido(t *testing.T) {
	server := httptest.NewServer(getRouter())
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	fechaReciente := time.Now().Add(-6 * 24 * time.Hour).Format("02-01-2006")
	reciboReciente := reFecha.ReplaceAllString(string(contenidoRecibo), fechaReciente)
	parametros := map[string]interface{}{
		"usuario":     "amerigal",
		"textoRecibo": reciboReciente,
	}
	e.POST("/recibos").WithForm(parametros).Expect().Status(http.StatusCreated)

	e.GET("/recuentos/semanal/amerigal").Expect().Status(http.StatusOK).JSON().Array().Element(0).Object().ValueEqual("tipo", "ESTROPAJO C/ ESPO")
}

func TestGetRecuentoApiNotFound(t *testing.T) {
	server := httptest.NewServer(getRouter())
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.GET("/recuentos/semanal/usuario-inventado").Expect().Status(http.StatusNotFound)
}

func TestGetTendenciaApiValido(t *testing.T) {
	server := httptest.NewServer(getRouter())
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	fechaReciente := time.Now().Add(-20 * 24 * time.Hour).Format("02-01-2006")
	reciboReciente := reFecha.ReplaceAllString(string(contenidoRecibo), fechaReciente)
	parametros := map[string]interface{}{
		"usuario":     "amerigal",
		"textoRecibo": reciboReciente,
	}
	e.POST("/recibos").WithForm(parametros).Expect().Status(http.StatusCreated)

	e.GET("/tendencias/GRANADA").Expect().Status(http.StatusOK).JSON().Array().Element(0).Object().ValueEqual("tipo", "ESTROPAJO C/ ESPO")
}

func TestGetTendenciaApiNotFound(t *testing.T) {
	server := httptest.NewServer(getRouter())
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.GET("/tendencias/LUGARINVENTADO").Expect().Status(http.StatusNotFound)
}
