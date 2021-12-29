package recibo

import (
	"testing"
	"time"
)

func TestCantidadesNoNulas(t *testing.T) {
	articulo, _ := newArticulo("descripcion", "tipo", 2.50, 'A')
	articuloRecibo := ArticuloRecibo{0, 0, articulo}
	_, err := newRecibo([]ArticuloRecibo{articuloRecibo}, time.Now(), "u", "l", "e")
	if err == nil {
		t.Fatalf("Positividad de las cantidades no validada.")
	}
}

func TestFechaNoFutura(t *testing.T) {
	articulo, _ := newArticulo("descripcion", "tipo", 2.50, 'A')
	articuloRecibo := ArticuloRecibo{0, 1, articulo}
	fecha := time.Now().Add(time.Hour * 24)
	_, err := newRecibo([]ArticuloRecibo{articuloRecibo}, fecha, "u", "l", "e")
	if err == nil {
		t.Fatalf("Validez de la fecha no comprobada.")
	}
}

func TestLeerReciboArchivo(t *testing.T) {
	_, err := leerRecibo("ruta/invalida/de/archivo")
	if err == nil {
		t.Fatalf("Corrección de lectura de archivo no comprobada.")
	}
}

func TestSiguienteIdPrimero(t *testing.T) {
	recibo, _ := newRecibo([]ArticuloRecibo{}, time.Now(), "u", "l", "e")
	if recibo.siguienteId() != 0 {
		t.Fatalf("Primer id no asigndo correctamente.")
	}
}

func TestSiguienteIdSegundo(t *testing.T) {
	articulo, _ := newArticulo("descripcion", "tipo", 2.50, 'A')
	articuloRecibo := ArticuloRecibo{0, 1, articulo}
	recibo, _ := newRecibo([]ArticuloRecibo{articuloRecibo}, time.Now(), "u", "l", "e")
	if recibo.siguienteId() != 1 {
		t.Fatalf("Los id no se asignan consecutivamente.")
	}
}

func TestSetTipoIdNoExistente(t *testing.T) {
	recibo, _ := newRecibo([]ArticuloRecibo{}, time.Now(), "u", "l", "e")
	_, err := recibo.setTipo(0, "tipo")
	if err == nil {
		t.Fatalf("Comprobación de existencia de id no realizada.")
	}
}

func TestSetTipo(t *testing.T) {
	articulo, _ := newArticulo("descripcion", "tipo", 2.50, 'A')
	articuloRecibo := ArticuloRecibo{0, 1, articulo}
	recibo, _ := newRecibo([]ArticuloRecibo{articuloRecibo}, time.Now(), "u", "l", "e")
	recibo.setTipo(0, "tipoNuevo")
	if recibo.articulos[0].articulo.tipo != "tipoNuevo" {
		t.Fatalf("Asignación incorrecta del atributo tipo.")
	}
}
