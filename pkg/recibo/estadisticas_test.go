package recibo

import (
	"fmt"
	"testing"
	"time"
)

func TestAgruparArticulosInexistente(t *testing.T) {
	articulo, _ := NewArticulo("descripcion", "tipo", 2.50, 'A')
	articuloRecibo := ArticuloRecibo{0, 1, articulo}
	recibo, _ := NewRecibo([]ArticuloRecibo{articuloRecibo}, time.Now(), "usuario", "l", "e")

	_, err := agruparArticulos([]Recibo{recibo}, "inexistente", time.Time{}, "inexistente")

	if err == nil {
		t.Fatalf("Inexistencia de recibos que cumplan los criterios no reportada.")
	}
}

func TestGetRecuentoSemanalCriterios(t *testing.T) {
	articulo, _ := NewArticulo("descripcion", "tipo1", 2.50, 'A')

	articuloRecibo1 := ArticuloRecibo{0, 1, articulo}
	reciboValido1, _ := NewRecibo([]ArticuloRecibo{articuloRecibo1}, time.Now(), "usuario", "l", "e")

	articuloRecibo2 := ArticuloRecibo{0, 2, articulo}
	reciboValido2, _ := NewRecibo([]ArticuloRecibo{articuloRecibo2}, time.Now().Add(-6*24*time.Hour), "usuario", "l", "e")

	reciboNoValido1, _ := NewRecibo([]ArticuloRecibo{articuloRecibo1}, time.Now().Add(-8*24*time.Hour), "usuario", "l", "e")
	reciboNoValido2, _ := NewRecibo([]ArticuloRecibo{articuloRecibo1}, time.Now(), "otro usuario", "l", "e")

	recibos := []Recibo{reciboValido1, reciboValido2, reciboNoValido1, reciboNoValido2}
	recuento, err := GetRecuentoSemanal(recibos, "usuario")

	if err != nil {
		t.Fatalf("Error al encontrar recibos que satisfagan las condiciones en recuento semanal.")
	}

	gasto := articulo.precio * (1 + iva[articulo.tipoIVA]) * 3
	if recuento[0].Tipo != "tipo1" || recuento[0].Gasto != gasto {
		t.Fatalf("Error al contemplar criterios en recuento semanal")
	}
}

func TestGetRecuentoSemanalTipos(t *testing.T) {
	var articulosRecibo [6]ArticuloRecibo

	for i := range articulosRecibo {
		tipo := fmt.Sprintf("tipo%d", i)
		precio := float32(i + 1)
		articulo, _ := NewArticulo("descripcion", tipo, precio, 'A')
		articulosRecibo[i] = ArticuloRecibo{0, 1, articulo}
	}

	recibo1, _ := NewRecibo([]ArticuloRecibo{articulosRecibo[0], articulosRecibo[1], articulosRecibo[2]}, time.Now(), "usuario", "l", "e")
	recibo2, _ := NewRecibo([]ArticuloRecibo{articulosRecibo[3], articulosRecibo[4], articulosRecibo[5]}, time.Now(), "usuario", "l", "e")

	recibos := []Recibo{recibo1, recibo2}
	recuento, _ := GetRecuentoSemanal(recibos, "usuario")

	if len(recuento) != 5 {
		t.Fatalf("Error en el número de artículos en recuento semanal.")
	}

	calculoCorrecto := true
	for i := range recuento {
		tipo := fmt.Sprintf("tipo%d", 5-i)
		gasto := float32(6-i) * (1 + iva['A'])
		if recuento[i].Tipo != tipo || recuento[i].Gasto != gasto {
			calculoCorrecto = false
		}
	}

	if !calculoCorrecto {
		t.Fatalf("Error al contemplar tipos en recuento semanal")
	}
}

func TestGetRecuentoMensual(t *testing.T) {
	var articulosRecibo [6]ArticuloRecibo

	for i := range articulosRecibo {
		tipo := fmt.Sprintf("tipo%d", i)
		precio := float32(i + 1)
		articulo, _ := NewArticulo("descripcion", tipo, precio, 'A')
		articulosRecibo[i] = ArticuloRecibo{0, 1, articulo}
	}

	reciboValido1, _ := NewRecibo([]ArticuloRecibo{articulosRecibo[0], articulosRecibo[1], articulosRecibo[2]}, time.Now(), "usuario", "l", "e")
	reciboValido2, _ := NewRecibo([]ArticuloRecibo{articulosRecibo[3], articulosRecibo[4], articulosRecibo[5]}, time.Now().Add(-29*24*time.Hour), "usuario", "l", "e")
	reciboNoValido1, _ := NewRecibo([]ArticuloRecibo{articulosRecibo[0]}, time.Now().Add(-32*24*time.Hour), "usuario", "l", "e")
	reciboNoValido2, _ := NewRecibo([]ArticuloRecibo{articulosRecibo[0]}, time.Now(), "otro usuario", "l", "e")

	recibos := []Recibo{reciboValido1, reciboValido2, reciboNoValido1, reciboNoValido2}
	recuento, _ := GetRecuentoMensual(recibos, "usuario")

	if len(recuento) != 5 {
		t.Fatalf("Error en el número de artículos en recuento mensual.")
	}

	calculoCorrecto := true
	for i := range recuento {
		tipo := fmt.Sprintf("tipo%d", 5-i)
		gasto := float32(6-i) * (1 + iva['A'])
		if recuento[i].Tipo != tipo || recuento[i].Gasto != gasto {
			calculoCorrecto = false
		}
	}

	if !calculoCorrecto {
		t.Fatalf("Error al calcular recuento mensual")
	}
}

func TestGetTendencia(t *testing.T) {
	var articulosRecibo [6]ArticuloRecibo

	for i := range articulosRecibo {
		tipo := fmt.Sprintf("tipo%d", i)
		precio := float32(i + 1)
		articulo, _ := NewArticulo("descripcion", tipo, precio, 'A')
		articulosRecibo[i] = ArticuloRecibo{0, 1, articulo}
	}

	reciboValido1, _ := NewRecibo([]ArticuloRecibo{articulosRecibo[0], articulosRecibo[1], articulosRecibo[2]}, time.Now(), "u", "lugar", "e")
	reciboValido2, _ := NewRecibo([]ArticuloRecibo{articulosRecibo[3], articulosRecibo[4], articulosRecibo[5]}, time.Now().Add(-100*24*time.Hour), "u", "lugar", "e")
	reciboNoValido1, _ := NewRecibo([]ArticuloRecibo{articulosRecibo[0]}, time.Now(), "u", "otro lugar", "e")
	reciboNoValido2, _ := NewRecibo([]ArticuloRecibo{articulosRecibo[0]}, time.Now(), "u", "otro lugar", "e")

	recibos := []Recibo{reciboValido1, reciboValido2, reciboNoValido1, reciboNoValido2}
	recuento, _ := GetTendencia(recibos, "lugar")

	if len(recuento) != 5 {
		t.Fatalf("Error en el número de artículos en cálculo de tendencia.%d", len(recuento))
	}

	calculoCorrecto := true
	for i := range recuento {
		tipo := fmt.Sprintf("tipo%d", 5-i)
		gasto := float32(6-i) * (1 + iva['A'])
		if recuento[i].Tipo != tipo || recuento[i].Gasto != gasto {
			calculoCorrecto = false
		}
	}

	if !calculoCorrecto {
		t.Fatalf("Error al calcular tendencia")
	}
}
