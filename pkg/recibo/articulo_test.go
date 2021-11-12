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
