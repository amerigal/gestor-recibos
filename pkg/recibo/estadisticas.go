package recibo

import (
	"fmt"
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
	Tipo string

	// Gasto es el gasto acumulado en el artículo en el recuento
	Gasto float32
}

// ErrorRecuento representa un error en un recuento de gasto

type errorRecuento struct {
	err string
}

// ErrorRecuento implementa la interfaz Error
func (e *errorRecuento) Error() string {
	return fmt.Sprintf("Error al hacer recuento: %s", e.err)
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
				tipo := articulo.Articulo.tipo
				if tipo == "" {
					tipo = articulo.Articulo.descripcion
				}

				gasto := float32(articulo.cantidad) * articulo.Articulo.precio * (1 + iva[articulo.Articulo.tipoIVA])

				articuloNuevo := true
				for i := range recuento {
					if tipo == recuento[i].Tipo { // Si ya tenemos un artículo con ese tipo en el recuento
						articuloNuevo = false
						recuento[i].Gasto += gasto
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
		return recuento[i].Gasto > recuento[j].Gasto
	})

	if len(recuento) == 0 {
		return recuento, &errorRecuento{"Ningún recibo coincide con los parámetros establecidos"}
	}

	if len(recuento) <= topSize {
		return recuento, nil
	}

	return recuento[:topSize], nil
}
