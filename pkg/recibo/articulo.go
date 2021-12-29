package recibo

import (
	"fmt"
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
	//   A -> IVA superreducido, 4%.
	//   B -> IVA reducido, 10%.
	//   C -> IVA general del 21%.
	tipoIVA byte
}

// ErrorArticulo representa un error en la creación de un Articulo

type errorArticulo struct {
	err string
}

// ErrorArticulo implementa la interfaz Error
func (e *errorArticulo) Error() string {
	return fmt.Sprintf("Error al crear artículo: %s", e.err)
}

// NewArticulo inicializa un objeto de tipo Articulo.
// Devuelve un objeto de tipo Articulo inicializado con los parámetros indicados.
func newArticulo(descripcion string, tipo string, precio float32, tipoIVA byte) (Articulo, error) {
	var articulo Articulo

	if descripcion == "" {
		return articulo, &errorArticulo{"descripción vacía"}
	}

	if precio < 0 {
		return articulo, &errorArticulo{"precio negativo"}
	}

	if tipoIVA != 'A' && tipoIVA != 'B' && tipoIVA != 'C' {
		return articulo, &errorArticulo{"tipoIVA valor incorrecto"}
	}

	articulo = Articulo{
		descripcion: descripcion,
		tipo:        tipo,
		precio:      precio,
		tipoIVA:     tipoIVA,
	}
	return articulo, nil
}

// SetTipo modifica el atributo tipo del Articulo art
func (art *Articulo) setTipo(tipo string) string {
	art.tipo = tipo
	return art.tipo
}
