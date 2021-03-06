// El paquete recibo provee de las estructuras de datos y el funcionamiento necesario para
// representar un recibo de una compra de un cliente.
package recibo

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Expresión regular para reconocer fecha formato: 02-01-2006 15:04
const regFecha = `(0[1-9]|[12][0-9]|3[01])-(0[1-9]|1[0-2])-\d{4}\s([01][0-9]|2[0-3]):([0-5][0-9])`

// Expresión regular para reconocer artículo formato: 2  PAÑUELOS CLASSIC    1.31   2.62  C
const regArticulo = `\d+\s+.*\d+\.\d*\s+[ABC]\n`

// Establecimiento correspondiente al formato inicial
const establecimiento1 = "ALIMENTACION GARO S.L"

// Lugar de compra correspondiente al formato inicial
const lugarCompra1 = "GRANADA"

// Formato fecha inicial
const layout = "02-01-2006 15:04"

// Posición de las unidades en formato inicial de artículo
const posUnidad = 0

// Posición del precio empezando por la derecha en formato inicial de artículo
const posPrecioInv = 3

// Posición del IVA empezando por la derecha en formato inicial de artículo
const posIVAInv = 1

// ArticuloRecibo representa un artículo concreto tal cual aparecerá en un recibo.
type ArticuloRecibo struct {
	// Id es el identificador de ArticuloRecibo dentro de un Recibo
	id uint

	// Cantidad es el número de unidades compradas de el artículo concreto.
	cantidad uint

	// Articulo es un artículo tal y como podría ser vendido por cualquier
	// establecimiento.
	articulo Articulo
}

// Recibo representa un recibo de la compra en un establecimiento, con información sobre
// precios, productos adquiridos, etcétera.
type Recibo struct {
	// Articulos es un slice de objetos de la clase ArticuloRecibo.
	articulos []ArticuloRecibo

	// FechaCompra representa la fecha en la que fue realizada la compra.
	fechaCompra time.Time

	// Usuario es una cadena que identifica al usuario que ha realizado la compra.
	usuario string

	// LugarCompra es una cadena que identifica la población en la que se
	// ha realizado la compra.
	lugarCompra string

	// Establecimiento es una cadena que corresponde al tipo de centro en el que se ha realizado la compra,
	// ya sea 'Frutería Paqui' o 'Mercadona'
	establecimiento string
}

// NewRecibo inicializa un objeto de tipo Recibo.
// Devuelve un objeto de tipo Recibo inicializado con los parámetros indicados.
func newRecibo(articulos []ArticuloRecibo, fechaCompra time.Time, usuario string,
	lugarCompra string, establecimiento string) (Recibo, error) {
	var recibo Recibo

	for _, articulo := range articulos {
		if articulo.cantidad == 0 {
			return recibo, &errorRecibo{"cantidad nula"}
		}
	}

	if fechaCompra.After(time.Now()) {
		return recibo, &errorRecibo{"fecha futura"}
	}

	recibo = Recibo{
		articulos:       articulos,
		fechaCompra:     fechaCompra,
		usuario:         usuario,
		lugarCompra:     lugarCompra,
		establecimiento: establecimiento,
	}
	return recibo, nil
}

// GetArticulo devuelve el ArticuloRecibo con id 'idArticulo'
func (recibo *Recibo) getArticulo(idArticulo uint) (*ArticuloRecibo, error) {
	var articulo *ArticuloRecibo
	encontrado := false

	for i := range recibo.articulos {
		if recibo.articulos[i].id == idArticulo {
			encontrado = true
			articulo = &recibo.articulos[i]
		}
	}

	if !encontrado {
		return articulo, &errorRecibo{fmt.Sprintf("no existe ningún articulo con id %d", idArticulo)}
	}
	return articulo, nil
}

// SetUsuario modifica el atributo usuario del Articulo art
func (recibo *Recibo) setUsuario(usuario string) {
	recibo.usuario = usuario
}

// SetTipo modifica el atributo tipo del articulo con idArticulo en el Recibo recibo
func (recibo *Recibo) setTipo(idArticulo uint, tipo string) error {
	articulo, err := recibo.getArticulo(idArticulo)

	if err != nil {
		return err
	}

	articulo.articulo.tipo = tipo

	return nil
}

// SiguienteId devuelve el siguiente al mayor id de un slice de artículos
func siguienteId(articulos []ArticuloRecibo) uint {
	if articulos == nil {
		return 0
	}

	if len(articulos) == 0 {
		return 0
	}

	maxId := uint(0)
	for _, articulo := range articulos {
		if articulo.id > maxId {
			maxId = articulo.id
		}
	}
	return maxId + 1
}

// LeerRecibo recibe un string referente a la ruta de un archivo
// que contiene un recibo de compra en texto plano y devuelve
// un objeto Recibo con la información proporcionada.
func leerRecibo(archivo string) (Recibo, error) {
	var recibo Recibo
	var articulosRecibo []ArticuloRecibo

	// Abrimos archivo para lectura
	data, err := ioutil.ReadFile(archivo)
	if err != nil {
		return recibo, err
	}
	contenido := string(data)

	// Identificamos formato del recibo
	formatoValido, _ := regexp.MatchString(establecimiento1, contenido)
	if !formatoValido {
		return recibo, &errorReciboLectura{"formato no reconocido"}
	}

	// Obtenemos fecha de compra
	reg := regexp.MustCompile(regFecha)
	fecha := reg.Find([]byte(contenido))
	fechaCompra, err := time.Parse(layout, string(fecha))
	if err != nil {
		return recibo, &errorReciboLectura{"fecha no válida"}
	}

	// Obtenemos líneas con artículos
	regAr := regexp.MustCompile(regArticulo)
	lineasArticulo := regAr.FindAll([]byte(contenido), -1)
	if lineasArticulo == nil {
		return recibo, &errorReciboLectura{"recibo sin artículos"}
	}

	// Creamos artículo para cada línea de artículos
	for _, art := range lineasArticulo {
		// Obtenemos los distintos campos del artículo
		art2 := strings.Fields(string(art))

		// Definimos posiciones de los atributos de acuerdo con el formato
		posUnd := posUnidad
		posPrecio := len(art2) - posPrecioInv
		posIVA := len(art2) - posIVAInv

		// Obtenemos atributos a partir de las posiciones y conversiones
		und_, _ := strconv.Atoi(art2[posUnd])
		und := uint(und_)
		precio_, _ := strconv.ParseFloat(art2[posPrecio], 32)
		precio := float32(precio_)
		tipoIVA := []byte(art2[posIVA])
		descripcion := strings.Join(art2[posUnd+1:posPrecio], " ")

		// Asignamos identificador
		id := siguienteId(articulosRecibo)

		// Creamos objeto Articulo
		articulo, err := newArticulo(descripcion, "", precio, tipoIVA[0])
		if err != nil {
			return recibo, err
		}
		articuloRecibo := ArticuloRecibo{id, und, articulo}

		// Añadimos articuloRecibo a slice de ArticuloRecibo
		articulosRecibo = append(articulosRecibo, articuloRecibo)
	}

	// Construimos objeto Recibo
	recibo, err = newRecibo(articulosRecibo, fechaCompra, "", lugarCompra1, establecimiento1)
	if err != nil {
		return recibo, err
	}

	return recibo, nil
}
