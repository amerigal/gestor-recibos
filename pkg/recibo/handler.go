package recibo

import "fmt"

// Handler es un manejador a través del cual se ejecutan las funcionalidades del sistema

type Handler struct {
	// Recibos es un diccionario que almacena objetos Recibo con su id como llave
	recibos map[uint]Recibo

	// MyLogger es un Logger para el registro de la actividad de Handler
	myLogger Logger
}

// NewHandler construye un objeto Handler inicializando sus atributos
func NewHandler() Handler {
	var handler Handler
	handler.recibos = make(map[uint]Recibo)
	handler.myLogger = NewLogger()
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
		h.myLogger.Error("No se ha podido crear recibo a partir de %s.", rutaArchivo)
		return 0, err
	}
	recibo.setUsuario(usuario)
	id := h.siguienteIdRecibo()
	h.recibos[id] = recibo
	h.myLogger.Info("Creado un recibo con ID=%d y usuario=%s a partir del archivo %s", id, usuario, rutaArchivo)
	return id, nil
}

// EliminarRecibo elimina el recibo con id 'idRecibo'
func (h *Handler) EliminarRecibo(idRecibo uint) error {
	_, existe := h.recibos[idRecibo]
	if !existe {
		h.myLogger.Error("No se ha podido eliminar recibo con ID=%d. Recibo no encontrado.", idRecibo)
		return &errorHandler{fmt.Sprintf("no se encontró ningún recibo con id %d", idRecibo)}
	}
	delete(h.recibos, idRecibo)
	h.myLogger.Info("Eliminado recibo con ID=%d.", idRecibo)
	return nil
}

// GetRecibo devuelve el recibo con id 'idRecibo'
func (h *Handler) GetRecibo(idRecibo uint) (Recibo, error) {
	recibo, existe := h.recibos[idRecibo]
	if !existe {
		h.myLogger.Error("No se ha podido obtener recibo con ID=%d. Recibo no encontrado.", idRecibo)
		return recibo, &errorHandler{fmt.Sprintf("no se encontró ningún recibo con id %d", idRecibo)}
	}
	h.myLogger.Info("Obtenido recibo con ID=%d.", idRecibo)
	return recibo, nil
}

// GetArticulosRecibo devuelve los artículos del recibo con id 'idRecibo'
func (h *Handler) GetArticulosRecibo(idRecibo uint) ([]ArticuloRecibo, error) {
	recibo, existe := h.recibos[idRecibo]
	if !existe {
		h.myLogger.Error("No se han podido obtener los artículos del recibo con ID=%d. Recibo no encontrado.", idRecibo)
		return []ArticuloRecibo{}, &errorHandler{fmt.Sprintf("no se encontró ningún recibo con id %d", idRecibo)}
	}
	h.myLogger.Info("Obtenidos los artículos del recibo con ID=%d.", idRecibo)
	return recibo.articulos, nil
}

// GetArticuloRecibo devuelve el artículo con id 'idArticulo' del recibo con id 'idRecibo'
func (h *Handler) GetArticuloRecibo(idRecibo uint, idArticulo uint) (ArticuloRecibo, error) {
	recibo, existeRecibo := h.recibos[idRecibo]
	if !existeRecibo {
		h.myLogger.Error("No se ha podido obtener el artículo con ID=%d del recibo con ID=%d. Recibo no encontrado.", idArticulo, idRecibo)
		return ArticuloRecibo{}, &errorHandler{fmt.Sprintf("no se encontró ningún recibo con id %d", idRecibo)}
	}

	articulo, err := recibo.getArticulo(idArticulo)

	if err != nil {
		h.myLogger.Error("No se ha podido obtener el artículo con ID=%d del recibo con ID=%d. Artículo no encontrado.", idArticulo, idRecibo)
		return ArticuloRecibo{}, &errorHandler{fmt.Sprintf("no se encontró ningún artículo con id %d en el recibo con id %d", idArticulo, idRecibo)}
	}

	h.myLogger.Info("Obtenido el artículo con ID=%d del recibo con ID=%d.", idArticulo, idRecibo)

	return *articulo, nil
}

// SetTipoArticuloRecibo asigna el tipo 'tipo' al artículo con id 'idArticulo' del recibo con id 'idRecibo'
func (h *Handler) SetTipoArticuloRecibo(tipo string, idRecibo uint, idArticulo uint) (ArticuloRecibo, error) {
	recibo, existeRecibo := h.recibos[idRecibo]
	if !existeRecibo {
		h.myLogger.Error("No se ha podido modificar el tipo del artículo con ID=%d del recibo con ID=%d. Recibo no encontrado.", idArticulo, idRecibo)
		return ArticuloRecibo{}, &errorHandler{fmt.Sprintf("no se encontró ningún recibo con id %d", idRecibo)}
	}

	err := recibo.setTipo(idArticulo, tipo)

	if err != nil {
		h.myLogger.Error("No se ha podido modificar el tipo del artículo con ID=%d del recibo con ID=%d. Artículo no encontrado.", idArticulo, idRecibo)
		return ArticuloRecibo{}, &errorHandler{fmt.Sprintf("no se encontró ningún artículo con id %d en el recibo con id %d", idArticulo, idRecibo)}
	}

	articulo, _ := recibo.getArticulo(idArticulo)
	h.myLogger.Info("Modificado el tipo del artículo con ID=%d del recibo con ID=%d a %s.", idArticulo, idRecibo, tipo)
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
	recuento, err := getRecuentoSemanal(recibos, usuario)
	if err != nil {
		h.myLogger.Warn("No hay artículos en el recuento semanal del usuario %s.", usuario)
		return recuento, err
	}
	h.myLogger.Info("Calculado el recuento semanal del usuario %s.", usuario)
	return recuento, nil
}

// GetRecuentoMensual devuelve el recuento mensual de gastos del usuario 'usuario'
func (h *Handler) GetRecuentoMensual(usuario string) ([]ArticuloRecuento, error) {
	recibos := h.getSliceRecibos()
	recuento, err := getRecuentoMensual(recibos, usuario)
	if err != nil {
		h.myLogger.Warn("No hay artículos en el recuento mensual del usuario %s.", usuario)
		return recuento, err
	}
	h.myLogger.Info("Calculado el recuento mensual del usuario %s.", usuario)
	return recuento, nil
}

// GetTendencia devuelve la tendencia en el lugar 'lugarCompra'
func (h *Handler) GetTendencia(lugarCompra string) ([]ArticuloRecuento, error) {
	recibos := h.getSliceRecibos()
	tendencia, err := getTendencia(recibos, lugarCompra)
	if err != nil {
		h.myLogger.Warn("No hay artículos registrados en el lugar solicitado (%s).", lugarCompra)
		return tendencia, err
	}
	h.myLogger.Info("Calculada la tendencia de gasto en %s.", lugarCompra)
	return tendencia, nil
}
