package recibo

import (
	"testing"
)

func TestDescripcionNoVacia(t *testing.T) {
	_, err := NewArticulo("", "tipo", 2.50, 'A')
	if err == nil {
		t.Fatalf("Contenido de descripción no validado.")
	}
}

func TestPrecioPositivo(t *testing.T) {
	_, err := NewArticulo("descripción", "tipo", -1, 'A')
	if err == nil {
		t.Fatalf("Positividad del precio no validada.")
	}
}
