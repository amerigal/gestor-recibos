package recibo

import (
	"testing"
)

func TestDescripcionNoVacia(t *testing.T) {
	_, err := newArticulo("", "tipo", 2.50, 'A')
	if err == nil {
		t.Fatalf("Contenido de descripción no validado.")
	}
}

func TestPrecioPositivo(t *testing.T) {
	_, err := newArticulo("descripción", "tipo", -1, 'A')
	if err == nil {
		t.Fatalf("Positividad del precio no validada.")
	}
}

func TestTipoIVAABC(t *testing.T) {
	_, err := newArticulo("descripción", "tipo", 2.50, 'D')
	if err == nil {
		t.Fatalf("Valor de tipoIVA no validado.")
	}
}
