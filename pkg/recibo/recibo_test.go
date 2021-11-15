package recibo

import (
	"testing"
	"time"
)

// TestCantidadesNoNulas comprueba la validación de la positividad de las
// cantidades de los objetos ArticuloRecibo al crear un Recibo.
func TestCantidadesNoNulas(t *testing.T) {
	articulo, _ := NewArticulo("descripcion", "tipo", 2.50, 'A')
	articuloRecibo := ArticuloRecibo{0, articulo}
	_, err := NewRecibo([]ArticuloRecibo{articuloRecibo}, time.Now(), "u", "l", "e")
	if err == nil {
		t.Fatalf("Positividad de las cantidades no validada.")
	}
}

// TestFechaNoFutura comprueba la validación de que la fecha de compra no sea
// posterior al momento actual al crear un Recibo.
func TestFechaNoFutura(t *testing.T) {
	articulo, _ := NewArticulo("descripcion", "tipo", 2.50, 'A')
	articuloRecibo := ArticuloRecibo{1, articulo}
	fecha := time.Now().Add(time.Hour * 24)
	_, err := NewRecibo([]ArticuloRecibo{articuloRecibo}, fecha, "u", "l", "e")
	if err == nil {
		t.Fatalf("Validez de la fecha no comprobada.")
	}
}

// TestLeerReciboArchivo comprueba que en la función LeerRecibo se
// valide la correcta lectura del archivo proporcionado.
func TestLeerReciboArchivo(t *testing.T) {
	_, err := LeerRecibo("ruta/invalida/de/archivo")
	if err == nil {
		t.Fatalf("Corrección de lectura de archivo no comprobada.")
	}
}
