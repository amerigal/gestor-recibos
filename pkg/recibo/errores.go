package recibo

import "fmt"

// ErrorArticulo representa un error en la creación de un Articulo
type errorArticulo struct {
	err string
}

// ErrorRecibo representa un error en un objeto Recibo
type errorRecibo struct {
	err string
}

// ErrorReciboLectura representa un error en la lectura de un Recibo
type errorReciboLectura struct {
	err string
}

// ErrorRecuento representa un error en un recuento de gasto
type errorRecuento struct {
	err string
}

// ErrorArticulo implementa la interfaz Error
func (e *errorArticulo) Error() string {
	return fmt.Sprintf("Error al crear artículo: %s", e.err)
}

// ErrorRecibo implementa la interfaz Error
func (e *errorRecibo) Error() string {
	return fmt.Sprintf("Error en objeto Recibo: %s", e.err)
}

// ErrorReciboLectura implementa la interfaz Error
func (e *errorReciboLectura) Error() string {
	return fmt.Sprintf("Error al leer Recibo: %s", e.err)
}

// ErrorRecuento implementa la interfaz Error
func (e *errorRecuento) Error() string {
	return fmt.Sprintf("Error al hacer recuento: %s", e.err)
}
