// El paquete recibo provee de las estructuras de datos y el funcionamiento necesario para
// reresentar un recibo de una compra de un cliente.
package recibo

import (
	"errors"
	"time"
)

// ArticuloRecibo representa un artículo concreto tal cual aparecerá en un recibo.
type ArticuloRecibo struct {
	// Cantidad es el número de unidades compradas de el artículo concreto.
	Cantidad uint

	// Articulo es un artículo tal y como podría ser vendido por cualquier
	// establecimiento.
	Articulo Articulo
}

// Recibo representa un recibo de la compra en un establecimiento, con información sobre
// precios, productos adquiridos, etcétera.
type Recibo struct {
	// Articulos es un slice de objetos de la clase ArticuloRecibo.
	articulos []ArticuloRecibo

	// FechaCompra representa la fecha en la que fue realizada la compra.
	fechaCompra time.Time

	// Usuario es una cadena que identifica al usuario que ha realizado la compra.
	usuario string

	// LugarCompra es una cadena que identifica la población en la que se
	// ha realizado la compra.
	lugarCompra string

	// Establecimiento es una cadena que corresponde al tipo de centro en el que se ha realizado la compra,
	// ya sea 'Frutería Paqui' o 'Mercadona'
	establecimiento string
}

// NewRecibo inicializa un objeto de tipo Recibo.
// Devuelve un objeto de tipo Recibo inicializado con los parámetros indicados.
func NewRecibo(articulos []ArticuloRecibo, fechaCompra time.Time, usuario string,
	lugarCompra string, establecimiento string) (Recibo, error) {
	var recibo Recibo

	for _, articulo := range articulos {
		if articulo.Cantidad == 0 {
			return recibo, errors.New("cantidad nula")
		}
	}

	if fechaCompra.After(time.Now()) {
		return recibo, errors.New("fecha futura")
	}

	recibo = Recibo{
		articulos:       articulos,
		fechaCompra:     fechaCompra,
		usuario:         usuario,
		lugarCompra:     lugarCompra,
		establecimiento: establecimiento,
	}
	return recibo, nil
}
