package recibo

import "time"

// ArticuloApi representa un ArticuloRecibo a ser devuelto por la API
type ArticuloApi struct {
	Id          uint    `json:"id"`
	Descripcion string  `json:"descripcion"`
	Cantidad    uint    `json:"cantidad"`
	Tipo        string  `json:"tipo"`
	Precio      float32 `json:"precio"`
	TipoIva     byte    `json:"tipoIva"`
}

// ToArticuloApi construye un ArticuloApi a partir de un ArticuloRecibo
func ToArticuloApi(a ArticuloRecibo) ArticuloApi {
	aApi := ArticuloApi{
		Id:          a.id,
		Descripcion: a.articulo.descripcion,
		Cantidad:    a.cantidad,
		Tipo:        a.articulo.tipo,
		Precio:      a.articulo.precio,
		TipoIva:     a.articulo.tipoIVA,
	}
	return aApi
}

// ToArticulosApi construye un slice de ArticuloApi a partir de un slice de ArticuloRecibo
func ToArticulosApi(a []ArticuloRecibo) []ArticuloApi {
	var aApi []ArticuloApi

	for _, articulo := range a {
		aApi = append(aApi, ToArticuloApi(articulo))
	}
	return aApi
}

// ReciboApi representa un Recibo a ser devuelto por la API
type ReciboApi struct {
	Id              uint          `json:"id"`
	Usuario         string        `json:"usuario"`
	LugarCompra     string        `json:"lugarCompra"`
	Establecimiento string        `json:"establecimiento"`
	FechaCompra     time.Time     `json:"fechaCompra"`
	Articulos       []ArticuloApi `json:"articulos"`
}

// ToReciboApi construye un ReciboApi a partir de un Recibo
func ToReciboApi(id uint, r Recibo) ReciboApi {
	rApi := ReciboApi{
		Id:              id,
		Usuario:         r.usuario,
		LugarCompra:     r.lugarCompra,
		Establecimiento: r.establecimiento,
		FechaCompra:     r.fechaCompra,
		Articulos:       ToArticulosApi(r.articulos),
	}
	return rApi
}

// ToRecibosApi construye un slice de ReciboApi a partir de un map que asocia id a Recibo
func ToRecibosApi(r map[uint]Recibo) []ReciboApi {
	var rApi []ReciboApi

	for id, recibo := range r {
		rApi = append(rApi, ToReciboApi(id, recibo))
	}
	return rApi
}

// ArticuloRecuentoApi representa un ArticuloRecuento a ser devuelto por la API
type ArticuloRecuentoApi struct {
	Tipo  string  `json:"tipo"`
	Gasto float32 `json:"gasto"`
}

// ToArticuloRecuentoApi construye un ArticuloRecuentoApi a partir de un ArticuloRecuento
func ToArticuloRecuentoApi(a ArticuloRecuento) ArticuloRecuentoApi {
	return ArticuloRecuentoApi{a.tipo, a.gasto}
}

// ToArticulosApi construye un slice de ArticuloRecuentoApi a partir de un slice de ArticuloRecuento
func ToArticulosRecuentoApi(a []ArticuloRecuento) []ArticuloRecuentoApi {
	var aApi []ArticuloRecuentoApi

	for _, articulo := range a {
		aApi = append(aApi, ToArticuloRecuentoApi(articulo))
	}
	return aApi
}
