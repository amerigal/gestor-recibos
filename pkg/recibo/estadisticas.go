package recibo

import (
	"sort"
	"time"
)

// Número de artículos con mayor gasto a devolver
const topSize = 5

// Valores de IVA en función del tipo
var iva = map[byte]float32{'A': 0.04, 'B': 0.1, 'C': 0.21}

// ArticuloRecuento representa un artículo de cara a realizar un recuento del gasto
type ArticuloRecuento struct {
	// Tipo es el tipo del artículo o, en su ausencia, su descripción
	tipo string

	// Gasto es el gasto acumulado en el artículo en el recuento
	gasto float32
}

// AgruparArticulos realiza un recuento de los gastos realizados en los artículos de
// aquellos recibos que satisfacen los siguientes criterios:
// 		- Si usuario != "" deben coincidir usuario y recibo.usuario
//		- Si lugar != "" deben coincidir lugar y recibo.lugarCompra
//		- Si fechaInicio != fechaNula, recibo.fechaCompra debe estar entre fechaInicio y time.Now()
// Se devuelven los 5 artículos con mayor gasto acumulado en dichos recibos, o todos los que hayan
// de ser menos de 5. En caso de no haya ningún artículo al final del recuento se devolver'un error.
func agruparArticulos(recibos []Recibo, usuario string, fechaIncio time.Time, lugar string) ([]ArticuloRecuento, error) {
	var recuento []ArticuloRecuento

	for _, recibo := range recibos {
		// Evaluamos criterios
		usuarioValido := recibo.usuario == usuario || usuario == ""
		lugarValido := recibo.lugarCompra == lugar || lugar == ""
		fechaValida := (recibo.fechaCompra.After(fechaIncio) && recibo.fechaCompra.Before(time.Now())) || fechaIncio.IsZero()

		if usuarioValido && lugarValido && fechaValida { // Si el recibo satisface las condiciones
			for _, articulo := range recibo.articulos {
				tipo := articulo.articulo.tipo
				if tipo == "" {
					tipo = articulo.articulo.descripcion
				}

				gasto := float32(articulo.cantidad) * articulo.articulo.precio * (1 + iva[articulo.articulo.tipoIVA])

				articuloNuevo := true
				for i := range recuento {
					if tipo == recuento[i].tipo { // Si ya tenemos un artículo con ese tipo en el recuento
						articuloNuevo = false
						recuento[i].gasto += gasto
					}
				}
				if articuloNuevo {
					recuento = append(recuento, ArticuloRecuento{tipo, gasto})
				}
			}
		}
	}

	// Ordenamos el recuento por gasto en orden decreciente
	sort.Slice(recuento, func(i, j int) bool {
		return recuento[i].gasto > recuento[j].gasto
	})

	if len(recuento) == 0 {
		return recuento, &errorRecuento{"Ningún recibo coincide con los parámetros establecidos"}
	}

	if len(recuento) <= topSize {
		return recuento, nil
	}

	return recuento[:topSize], nil
}

// GetRecuentoSemanal devuelve los artículos en los que usuario ha realizado mayor
// gasto en la última semana.
func getRecuentoSemanal(recibos []Recibo, usuario string) ([]ArticuloRecuento, error) {
	return agruparArticulos(recibos, usuario, time.Now().Add(-7*24*time.Hour), "")
}

// GetRecuentoMensual devuelve los artículos en los que usuario ha realizado mayor
// gasto en el último mes.
func getRecuentoMensual(recibos []Recibo, usuario string) ([]ArticuloRecuento, error) {
	return agruparArticulos(recibos, usuario, time.Now().Add(-30*24*time.Hour), "")
}

// GetTendencia devuelve los artículos en los que se ha realizado mayor gasto en lugarCompra.
func getTendencia(recibos []Recibo, lugarCompra string) ([]ArticuloRecuento, error) {
	return agruparArticulos(recibos, "", time.Time{}, lugarCompra)
}
