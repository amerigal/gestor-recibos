package main

import (
	"fmt"
	"proyecto_iv/proyecto/pkg/recibo"
	"time"

	"cloud.google.com/go/civil"
)

func main() {

	var articulo = recibo.NewArticulo(
		"Leche entera 1L",
		"PULEVA",
		"LECHE",
		0.80,
		"GRANADA",
		"ESPAÑA",
		civil.Date{2021, time.April, 01},
		'C')

	var articulo_recibo = recibo.ArticuloRecibo{1, 0, articulo}
	array_articulos := []recibo.ArticuloRecibo{articulo_recibo}

	var recibo = recibo.NewRecibo(
		array_articulos,
		civil.Date{2021, time.October, 01},
		"01",
		"Granada",
		"Bizum",
		"Coviran")

	fmt.Printf("Recibo de ejemplo: %+v\n", recibo)
}
