/*
El paquete recibo provee de las estructuras de datos y el funcionamiento necesario para
representar un recibo de una compra de un cliente.
*/
package recibo

import (
	"cloud.google.com/go/civil"
)

/*
ArticuloRecibo representa un artículo concreto tal cual aparecerá en un recibo.
*/
type ArticuloRecibo struct {
	// Cantidad es el número de unidades compradas de el artículo concreto.
	Cantidad int

	// Descuento consiste en el porcentaje de descuento aplicado al precio del
	// artículo por la tienda.
	//
	// Se representa como un número real entre 0 y 1, siendo 0 ningún descuento y 1
	// un descuento del 100% sobre el precio total del producto.
	Descuento float32

	// articulo es un artículo tal y como podría ser vendido por cualquier
	// establecimiento.
	Articulo Articulo
}

/*
Recibo representa un recibo de la compra en un establecimiento, con información sobre
precios, productos adquiridos, etcétera.
*/
type Recibo struct {
	// Articulos es un array de objetos de la clase ArticuloRecibo.
	Articulos []ArticuloRecibo

	// FechaCompra representa la fecha en la que fue realizada la compra.
	FechaCompra civil.Date

	// Usuario es una cadena que identifica al usuario que ha realizado la compra.
	Usuario string

	// LugarCompra es una cadena que identifica la población en la que se
	// ha realizado la compra.
	LugarCompra string

	// MetodoPago es una cadena que identifica el método de pago usado por el cliente
	// para pagar la compra.
	MetodoPago string

	// Establecimiento es una cadena que corresponde al tipo de centro en el que se ha realizado la compra,
	// ya sea 'Frutería Paqui' o 'Mercadona'
	Establecimiento string
}
