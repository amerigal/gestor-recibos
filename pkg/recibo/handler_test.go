package recibo

import (
	"testing"
)

const rutaEjemplo = "../../input/reciboCoviran.txt"
const numArticulosEjemplo = 13

func TestSiguienteIdReciboPrimero(t *testing.T) {
	handler := NewHandler()
	if handler.siguienteIdRecibo() != 0 {
		t.Fatalf("Primer id no asigndo correctamente.")
	}
}

func TestSiguienteIdReciboSegundo(t *testing.T) {
	handler := NewHandler()
	handler.InsertarRecibo(rutaEjemplo, "usuario")
	if handler.siguienteIdRecibo() != 1 {
		t.Fatalf("Los id no se asignan consecutivamente.")
	}
}

func TestInsertarRecibo(t *testing.T) {
	handler := NewHandler()
	id, err := handler.InsertarRecibo(rutaEjemplo, "usuario")

	if err != nil {
		t.Fatalf("Archivo válido no es procesado correctamente.")
	}

	recibo, _ := handler.GetRecibo(id)
	if recibo.usuario != "usuario" {
		t.Fatalf("Recibo inicializado incorrectamente.")

	}
}

func TestEliminarReciboNoValido(t *testing.T) {
	handler := NewHandler()
	err := handler.EliminarRecibo(0)

	if err == nil {
		t.Fatalf("Comprobación de existencia de recibo no realizada.")
	}
}

func TestEliminarReciboValido(t *testing.T) {
	handler := NewHandler()
	id, _ := handler.InsertarRecibo(rutaEjemplo, "usuario")

	handler.EliminarRecibo(id)

	_, err := handler.GetRecibo(id)

	if err == nil {
		t.Fatalf("Recibo válido no eliminado.")
	}
}

func TestGetArticulosRecibo(t *testing.T) {
	handler := NewHandler()
	id, _ := handler.InsertarRecibo(rutaEjemplo, "usuario")

	articulos, err := handler.GetArticulosRecibo(id)

	if err != nil {
		t.Fatalf("Error al obtener articulos de recibo válido.")
	}

	if len(articulos) != numArticulosEjemplo {
		t.Fatalf("Número incorrecto de articulos obtenidos de recibo válido.")
	}
}

func TestSetTipoArticuloRecibo(t *testing.T) {
	handler := NewHandler()
	id, _ := handler.InsertarRecibo(rutaEjemplo, "usuario")

	_, err := handler.SetTipoArticuloRecibo("tipoEjemplo", id, 0)

	if err != nil {
		t.Fatalf("Error al asignar tipo a artículo válido.")
	}

	articulo, _ := handler.GetArticuloRecibo(id, 0)
	if articulo.articulo.tipo != "tipoEjemplo" {
		t.Fatalf("Asignación errónea de tipo a artículo válido.")
	}
}
