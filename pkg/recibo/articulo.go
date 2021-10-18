package recibo

import (
	"cloud.google.com/go/civil"
)

//Articulo representa un artículo vendido por un establecimiento.

type Articulo struct {
	// descripcion consiste en una breve descripción del artículo tal cual aparece en
	// un recibo, como 'Leche COVAP entera 1L'.
	descripcion string

	// fabricante representa al fabricante del producto.
	fabricante string

	// tipo representa una cadena a partir de la cual poder agrupar artículos similares,
	// como todos los tipos de leche, pan o frutas de diversos fabricantes.
	tipo string

	// precio consiste en el del artículo en euros.
	precio float32

	// lugarFabricacion es una cadena que representa la ciudad en la que se ha fabricado
	// el producto.
	lugarFabricacion string

	// paisFabricacion es una cadena que representa el país en el que se ha fabricado el producto.
	paisFabricacion string

	// fechaFabricacion representa la fecha en la que el artículo fue fabricado.
	fechaFabricacion civil.Date

	// tipoIVA es un carácter ASCII que podrá tener los siguientes valores:
	//   A -> IVA general del 21%.
	//   B -> IVA reducido, 10%.
	//   C -> IVA superreducido, 4%.
	tipoIVA byte
}

func NewArticulo(descripcion string, fabricante string, tipo string, precio float32,
	lugarFabricacion string, paisFabricacion string, fechaFabricacion civil.Date,
	tipoIVA byte) Articulo{
		articulo := Articulo{
			descripcion: descripcion,
			fabricante: fabricante,
			tipo: tipo,
			precio: precio,
			lugarFabricacion: lugarFabricacion,
			paisFabricacion: paisFabricacion,
			fechaFabricacion: fechaFabricacion,
			tipoIVA: tipoIVA,
		}
		return articulo
	}
