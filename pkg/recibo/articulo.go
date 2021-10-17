/*
El paquete recibo provee de las estructuras de datos y el funcionamiento necesario para
representar un recibo de una compra de un cliente.
*/
package recibo

import (
	"cloud.google.com/go/civil"
)

/*
Articulo representa un artículo vendido por un establecimiento.
*/
type Articulo struct {
	// Descripcion consiste en una breve descripción del artículo tal cual aparece en
	// un recibo, como 'Leche COVAP entera 1L'.
	Descripcion string

	// Fabricante representa al fabricante del producto.
	Fabricante string

	// Tipo representa una cadena a partir de la cual poder agrupar artículos similares,
	// como todos los tipos de leche, pan o frutas de diversos fabricantes.
	Tipo string

	// Precio consiste en el del artículo en euros.
	Precio float32

	// LugarFabricacion es una cadena que representa la ciudad en la que se ha fabricado
	// el producto.
	LugarFabricacion string

	// PaisFabricacion es una cadena que representa el país en el que se ha fabricado el producto.
	PaisFabricacion string

	// FechaFabricacion representa la fecha en la que el artículo fue fabricado.
	FechaFabricacion civil.Date

	// TipoIVA es un carácter ASCII que podrá tener los siguientes valores:
	//   A -> IVA general del 21%.
	//   B -> IVA reducido, 10%.
	//   C -> IVA superreducido, 4%.
	TipoIVA byte
}
