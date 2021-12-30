package recibo

import "fmt"

// Handler es un manejador a través del cual se ejecutan las funcionalidades del sistema

type Handler struct {
	// Recibos es un diccionario que almacena objetos Recibo con su id como llave
	recibos map[uint]Recibo
}

func NewHandler() Handler {
	var handler Handler
	handler.recibos = make(map[uint]Recibo)
	return handler
}

// SiguienteIdRecibo devuelve el siguiente id válido para un nuevo Recibo
func (h *Handler) siguienteIdRecibo() uint {
	if h.recibos == nil {
		return 0
	}

	if len(h.recibos) == 0 {
		return 0
	}

	maxId := uint(0)
	for id := range h.recibos {
		if id > maxId {
			maxId = id
		}
	}
	return maxId + 1
}

// InsertarRecibo añade un recibo a recibos creándolo a partir del archivo
// ubicado en 'rutaArchivo' y asignándolo a 'usuario'. Devuelve el id del Recibo
// insertado.
func (h *Handler) InsertarRecibo(rutaArchivo string, usuario string) (uint, error) {
	recibo, err := leerRecibo(rutaArchivo)
	if err != nil {
		return 0, err
	}
	recibo.setUsuario(usuario)
	id := h.siguienteIdRecibo()
	h.recibos[id] = recibo
	return id, nil
}

// EliminarRecibo elimina el recibo con id 'idRecibo'
func (h *Handler) EliminarRecibo(idRecibo uint) error {
	_, existe := h.recibos[idRecibo]
	if !existe {
		return &errorHandler{fmt.Sprintf("no se encontró ningún recibo con id %d", idRecibo)}
	}
	delete(h.recibos, idRecibo)
	return nil
}

// GetRecibo devuelve el recibo con id 'idRecibo'
func (h *Handler) GetRecibo(idRecibo uint) (Recibo, error) {
	recibo, existe := h.recibos[idRecibo]
	if !existe {
		return recibo, &errorHandler{fmt.Sprintf("no se encontró ningún recibo con id %d", idRecibo)}
	}
	return recibo, nil
}

// GetArticulosRecibo devuelve los artículos del recibo con id 'idRecibo'
func (h *Handler) GetArticulosRecibo(idRecibo uint) ([]ArticuloRecibo, error) {
	recibo, existe := h.recibos[idRecibo]
	if !existe {
		return []ArticuloRecibo{}, &errorHandler{fmt.Sprintf("no se encontró ningún recibo con id %d", idRecibo)}
	}
	return recibo.articulos, nil
}

// GetArticuloRecibo devuelve el artículo con id 'idArticulo' del recibo con id 'idRecibo'
func (h *Handler) GetArticuloRecibo(idRecibo uint, idArticulo uint) (ArticuloRecibo, error) {
	recibo, existeRecibo := h.recibos[idRecibo]
	if !existeRecibo {
		return ArticuloRecibo{}, &errorHandler{fmt.Sprintf("no se encontró ningún recibo con id %d", idRecibo)}
	}

	articulo, err := recibo.getArticulo(idArticulo)

	if err != nil {
		return ArticuloRecibo{}, &errorHandler{fmt.Sprintf("no se encontró ningún artículo con id %d en el recibo con id %d", idArticulo, idRecibo)}
	}

	return *articulo, nil
}

// SetTipoArticuloRecibo asigna el tipo 'tipo' al artículo con id 'idArticulo' del recibo con id 'idRecibo'
func (h *Handler) SetTipoArticuloRecibo(tipo string, idRecibo uint, idArticulo uint) (ArticuloRecibo, error) {
	recibo, existeRecibo := h.recibos[idRecibo]
	if !existeRecibo {
		return ArticuloRecibo{}, &errorHandler{fmt.Sprintf("no se encontró ningún recibo con id %d", idRecibo)}
	}

	err := recibo.setTipo(idArticulo, tipo)

	if err != nil {
		return ArticuloRecibo{}, &errorHandler{fmt.Sprintf("no se encontró ningún artículo con id %d en el recibo con id %d", idArticulo, idRecibo)}
	}

	articulo, _ := recibo.getArticulo(idArticulo)
	return *articulo, nil
}

// GetSliceRecibos devuelve un slice con los recibos almacenados en el map h.recibos
func (h *Handler) getSliceRecibos() []Recibo {
	var values []Recibo
	for _, value := range h.recibos {
		values = append(values, value)
	}
	return values
}

// GetRecuentoSemanal devuelve el recuento semanal de gastos del usuario 'usuario'
func (h *Handler) GetRecuentoSemanal(usuario string) ([]ArticuloRecuento, error) {
	recibos := h.getSliceRecibos()
	return getRecuentoSemanal(recibos, usuario)
}

// GetRecuentoMensual devuelve el recuento mensual de gastos del usuario 'usuario'
func (h *Handler) GetRecuentoMensual(usuario string) ([]ArticuloRecuento, error) {
	recibos := h.getSliceRecibos()
	return getRecuentoMensual(recibos, usuario)
}

// GetTendencia devuelve la tendencia en el lugar 'lugarCompra'
func (h *Handler) GetTendencia(lugarCompra string) ([]ArticuloRecuento, error) {
	recibos := h.getSliceRecibos()
	return getTendencia(recibos, lugarCompra)
}
