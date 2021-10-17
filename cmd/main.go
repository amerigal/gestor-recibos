package main

import (
	"fmt"
	"proyecto_iv/proyecto/pkg/recibo"
	"time"

	"cloud.google.com/go/civil"
)

func main() {

	var articulo = recibo.Articulo{"Leche entera 1L", "PULEVA", "LECHE", 0.80, "GRANADA", "ESPAÑA", civil.Date{2021, time.April, 01}, 'C'}
	var articulo_recibo = recibo.ArticuloRecibo{1, 0, articulo}
	array_articulos := []recibo.ArticuloRecibo{articulo_recibo}

	var recibo = recibo.Recibo{array_articulos, civil.Date{2021, time.October, 01}, "01", "Granada", "Bizum", "Coviran"}

	fmt.Print("Descripción del artículo: ", articulo.Descripcion, "\n")
	fmt.Print("Lugar de compra: "+recibo.LugarCompra, "\n")
}
