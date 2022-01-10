package recibo

import (
	"encoding/json"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// ApiHandler es un manejador para procesar peticiones HTTP

type ApiHandler struct {
	handler Handler
}

// NewApiHandler construye un objeto ApiHandler inicializando sus atributos
func NewApiHandler() ApiHandler {
	return ApiHandler{NewHandler()}
}

// GetStatusApi es un manejador para procesar la petición de estado de la API
func (h *ApiHandler) GetStatusApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)
	resp["Estado"] = "Todo OK!"
	writeJsonIntoW(w, resp)
}

// GetRecibosApi es un manejador para procesar la petición de recibos de la API
func (h *ApiHandler) GetRecibosApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	writeJsonIntoW(w, ToRecibosApi(h.handler.recibos))
}

// InsertarReciboApi es un manejador para procesar la petición de inserción de recibo de la API
func (h *ApiHandler) InsertarReciboApi(w http.ResponseWriter, r *http.Request) {
	usuario := r.FormValue("usuario")
	textoRecibo := r.FormValue("textoRecibo")

	if usuario == "" || textoRecibo == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Petición Fallida: No se han proporcionado los parámetros requeridos."))
	} else {
		pathFichero := "/tmp/recibo-" + usuario + "-" + time.Now().Format("2006-01-02-15:04:05")
		err := os.WriteFile(pathFichero, []byte(textoRecibo), os.ModePerm)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Petición Fallida: Los parámetros proporcionados no permiten un procesamiento correcto."))
		} else {
			id, err := h.handler.InsertarRecibo(pathFichero, usuario)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Petición Fallida: Formato de recibo no válido"))
			} else {
				recibo, _ := h.handler.GetRecibo(id)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				writeJsonIntoW(w, ToReciboApi(id, recibo))
			}
		}
	}
}

// GetReciboApi es un manejador para procesar la petición de recibo de la API
func (h *ApiHandler) GetReciboApi(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idR, _ := strconv.Atoi(params["idR"])

	recibo, err := h.handler.GetRecibo(uint(idR))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(ToErrorApi(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		writeJsonIntoW(w, ToReciboApi(uint(idR), recibo))
	}
}

// EliminarReciboApi es un manejador para procesar la petición de eliminación de recibo de la API
func (h *ApiHandler) EliminarReciboApi(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idR, _ := strconv.Atoi(params["idR"])

	err := h.handler.EliminarRecibo(uint(idR))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(ToErrorApi(err.Error()))
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

// GetArticulosReciboApi es un manejador para procesar la petición de articulos de recibo de la API
func (h *ApiHandler) GetArticulosReciboApi(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idR, _ := strconv.Atoi(params["idR"])

	articulos, err := h.handler.GetArticulosRecibo(uint(idR))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(ToErrorApi(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		writeJsonIntoW(w, ToArticulosApi(articulos))
	}
}

// GetArticuloReciboApi es un manejador para procesar la petición de articulo de recibo de la API
func (h *ApiHandler) GetArticuloReciboApi(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idR, _ := strconv.Atoi(params["idR"])
	idA, _ := strconv.Atoi(params["idA"])

	articulo, err := h.handler.GetArticuloRecibo(uint(idR), uint(idA))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(ToErrorApi(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		writeJsonIntoW(w, ToArticuloApi(articulo))
	}
}

// SetTipoArticuloReciboApi es un manejador para procesar la petición de cambiar tipo de articulo de recibo de la API
func (h *ApiHandler) SetTipoArticuloReciboApi(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idR, _ := strconv.Atoi(params["idR"])
	idA, _ := strconv.Atoi(params["idA"])
	tipo := r.FormValue("tipo")

	if tipo == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Petición Fallida: No se han proporcionado los parámetros requeridos."))
	} else {
		articulo, err := h.handler.SetTipoArticuloRecibo(tipo, uint(idR), uint(idA))

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(ToErrorApi(err.Error()))
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			writeJsonIntoW(w, ToArticuloApi(articulo))
		}
	}

}

// GetRecuentoApi es un manejador para procesar la petición de recuento de la API
func (h *ApiHandler) GetRecuentoApi(w http.ResponseWriter, r *http.Request) {
	var recuento []ArticuloRecuento
	var err error
	params := mux.Vars(r)
	usuario, _ := params["usuario"]
	periodo, _ := params["periodo"]

	if periodo == "semanal" {
		recuento, err = h.handler.GetRecuentoSemanal(usuario)
	} else {
		recuento, err = h.handler.GetRecuentoMensual(usuario)
	}

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(ToErrorApi(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		writeJsonIntoW(w, ToArticulosRecuentoApi(recuento))
	}

}

// GetTendenciaApi es un manejador para procesar la petición de tendencia de la API
func (h *ApiHandler) GetTendenciaApi(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	lugar, _ := params["lugar"]

	tendencia, err := h.handler.GetTendencia(lugar)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(ToErrorApi(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		writeJsonIntoW(w, ToArticulosRecuentoApi(tendencia))
	}
}

// ToErrorApi construye un mensaje de error a ser devuelto por la API
func ToErrorApi(err string) []byte {
	re := regexp.MustCompile(`[\w\s]*:`)
	return []byte(re.ReplaceAllString(err, "Peticion fallida:"))
}

// WriteJsonIntoW escribe contenido en formato json en el http.ResponseWriter w
func writeJsonIntoW(w http.ResponseWriter, contenido interface{}) {
	jsonResp, _ := json.Marshal(contenido)
	w.Write(jsonResp)
}
