package recibo

import (
	"errors"
)

// Articulo representa un artículo vendido por un establecimiento.

type Articulo struct {
	// Descripcion consiste en una breve descripción del artículo tal cual aparece en
	// un recibo, como 'Leche COVAP entera 1L'.
	descripcion string

	// Tipo representa una cadena a partir de la cual poder agrupar artículos similares,
	// como todos los tipos de leche, pan o frutas de diversos fabricantes.
	tipo string

	// Precio consiste en el del artículo en euros.
	precio float32

	// TipoIVA es un carácter ASCII que podrá tener los siguientes valores:
	//   A -> IVA general del 21%.
	//   B -> IVA reducido, 10%.
	//   C -> IVA superreducido, 4%.
	tipoIVA byte
}

// NewArticulo inicializa un objeto de tipo Articulo.
// Devuelve un objeto de tipo Articulo inicializado con los parámetros indicados.
func NewArticulo(descripcion string, tipo string, precio float32, tipoIVA byte) (Articulo, error) {
	var articulo Articulo

	if descripcion == "" {
		return articulo, errors.New("descripción vacía")
	}

	if precio < 0 {
		return articulo, errors.New("precio negativo")
	}

	articulo = Articulo{
		descripcion: descripcion,
		tipo:        tipo,
		precio:      precio,
		tipoIVA:     tipoIVA,
	}
	return articulo, nil
}
