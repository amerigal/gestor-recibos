package recibo

import (
	"encoding/json"
	"net/http"
)

type ApiHandler struct {
	handler Handler
}

func NewApiHandler() ApiHandler {
	return ApiHandler{NewHandler()}
}

func (h *ApiHandler) GetStatusApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)
	resp["Estado"] = "Todo OK!"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		GetLogger().Error("Error al convertir en json: %s", err)
	}

	w.Write(jsonResp)
	w.WriteHeader(http.StatusOK)
}

func (h *ApiHandler) GetRecibosApi(w http.ResponseWriter, r *http.Request) {
}

func (h *ApiHandler) InsertarReciboApi(w http.ResponseWriter, r *http.Request) {
}

func (h *ApiHandler) GetReciboApi(w http.ResponseWriter, r *http.Request) {
}

func (h *ApiHandler) EliminarReciboApi(w http.ResponseWriter, r *http.Request) {
}

func (h *ApiHandler) GetArticulosReciboApi(w http.ResponseWriter, r *http.Request) {
}

func (h *ApiHandler) GetArticuloReciboApi(w http.ResponseWriter, r *http.Request) {
}

func (h *ApiHandler) SetTipoArticuloReciboApi(w http.ResponseWriter, r *http.Request) {
}

func (h *ApiHandler) GetRecuentoApi(w http.ResponseWriter, r *http.Request) {
}

func (h *ApiHandler) GetTendenciaApi(w http.ResponseWriter, r *http.Request) {
}
