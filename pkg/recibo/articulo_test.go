package recibo

import (
	"testing"
)

// TestDescripcionNoVacia comprueba la validación de que
// la descripción no sea vacía al crear un Articulo
func TestDescripcionNoVacia(t *testing.T) {
	_, err := NewArticulo("", "tipo", 2.50, 'A')
	if err == nil {
		t.Fatalf("Contenido de descripción no validado.")
	}
}

// TestPrecioPositivo comprueba la validación de que
// el precio sea positivo al crear un Articulo
func TestPrecioPositivo(t *testing.T) {
	_, err := NewArticulo("descripción", "tipo", -1, 'A')
	if err == nil {
		t.Fatalf("Positividad del precio no validada.")
	}
}

// TestTipoIVAABC comprueba la validación de que
// el tipo de IVA pertenezca a {A,B,C} al crear un Articulo
func TestTipoIVAABC(t *testing.T) {
	_, err := NewArticulo("descripción", "tipo", 2.50, 'D')
	if err == nil {
		t.Fatalf("Valor de tipoIVA no validado.")
	}
}
