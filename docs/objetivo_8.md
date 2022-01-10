# Documentación Objetivo 8
## Diseño API REST

Va a desarrollarse una API REST para proporcionar acceso a los recursos del sistema y posibilitar determinadas operaciones con ellos, habilitando así las funcionalidades pedidas en las historias de usuario planteadas inicialmente. Lo primero a tener en cuenta es qué framework usar.

### Framework para API

Es necesario definir primero qué requisitos buscamos en el paquete a utilizar dado que hay varios disponibles en go para la implementación de la API. Cabe mencionar que no hay versiones de Flask ni Express en dicho lenguaje.

Buscamos un paquete que nos permita capturar peticiones http y devolver respuestas con la información que deseemos. Deberá ser compatible con el formato json para trabajar de una forma más cómoda. Asímismo, es conveniente que permita una sencilla definición de las URI a las que vamos a dar acceso, permitiendo la definición de parámetros en las mismas.

Como en otras ocasiones, es preferible que el paquete implique una sencilla configuración y que esté mantenido activamente.

He encontrado las siguientes opciones:

- [net/http](https://pkg.go.dev/net/http): Se trata del paquete de http incluido en la biblioteca estándar de go. Es bastante completo, proporcionando implementaciones para cliente y servidor de HTTP, aunque es un poco pobre en lo que se refiere al manejo de las URI. Es una opción interesante en tanto que no origina dependencias externas y, en general, el resto de paquetes externos están construidos sobre este.

- [gorilla/mux](https://pkg.go.dev/github.com/gorilla/mux): Paquete que proporciona un multiplexador de peticiones http, esto es, permite una cómoda y flexible definición de las rutas y parámetros de petición de manera que las peticiones entrantes son reconocidas y asignadas al manejador correspondiente de una manera sencilla. Se trata de un paquete estable y ampliamente extendido.

- [Revel](https://revel.github.io/): Se trata de un framework para web que proporciona funcionaliades de enrutamiento, parseo de parámetros, sesiones, templates, y tests entre otros. En particular, proporciona la funcionalidad necesaria para montar la API que buscamos. Sin embargo parece que no tiene mucha actividad el proyecto pues su [repositorio](https://github.com/revel/revel) en GitHub tiene abiertos muchos issues sin responder desde hace tiempo y no hay señales de un mantenimiento activo.

- [Beego](https://beego.vip/): Framework HTTP RESTful enfocado en el desarrollo rápido de aplicaciones en Go, incluyendo APIs, aplicaciones web y servicios de backend. Permite la definición flexible de rutas para asignarlas a los manejadores. En definitiva, proporciona toda la funcionalidad necesaria para desarrollar la API que necesitamos y el [repositorio](https://github.com/beego/beego) está mantenido activamente.

Analizando la configuración para trabajar con Beego y la amplia extensión de funcionalidades que proporciona, pienso que queda grande para la tarea que se pretende en el proyecto y que puede ser más que resuelta con el paquete net/http ya integrado en el estándar con el apoyo de los routers de gorilla/mux. En este sentido, he decidido utilizar net/http y gorilla/mux por su sencillez de uso y directa aplicación al proyecto, al tiempo que cubren sin problema las necesidades del mismo.

### Framework para tests de la API

Se requiere un paquete que permita la generación de peticiones http y que permita procesarlas a través de los manejadores implementados, testeando los distintos campos de las respuestas. He encontrado dos opciones usadas para este propósito:

- [httptest](https://pkg.go.dev/net/http/httptest): Está incluido en el paquete net/http de la biblioteca estándar de go. Permite generar peticiones http indicando la url y el método y asignándolas a los manejadores. Se devuelve la respuesta correspondiente, cuyo contenido puede ser analizado para ver si coincide con el esperado.

- [httpexpect](https://pkg.go.dev/github.com/gavv/httpexpect): Integra la generación y procesamiento de peticiones http con aserciones para testear las respuestas. Está construida sobre net/http y simplifica mucho los tests respecto a httptest. Además se trata de un paquete bastante usado y con un activo mantenimiento. Es por dichas razones que he decido utilizar este paquete para realizar los tests.

### Recursos

Lo primero importante es determinar los recursos a los que se va a dar acceso, que deben ir en la línea de lo indicado en las historias de usuario planteadas al inicio (#2, #3 y #4). Con eso en mente, he decido que el recurso básico al que acceder sea el recibo, dentro del cuál se podra acceder a los otros recursos principales que son los artículos.

Para dar acceso a la funcionalidad de recuento semanal y mensual (#3), he decidido considerar el recurso recuento, que se concretará indicando el período y el usuario.

Por último, en referencia a la funcionalidad de calcular los artículos en los que se ha realizado más gasto en una determinada zona (#4), he decidido contemplar el recurso tendencia, que se concretará con el lugar de compra.

### Rutas

A continuación describiré las rutas definidas para el acceso a los recursos antes descritos. Cabe mencionar que, aunque teóricamente se podrían realizar más operaciones con dichos recursos, he restringido las operaciones a aquellas que posibilitan y proporcionan una manera cómoda de acceder a las funcionalidades asociadas a las historias de usuario (#2, #3 y #4).
También comentar que en la definición de las URI, he seguido el estándar de considerar la forma plurar de los recursos, pese a que tras la especificación de parámetros se acabe devolviendo un único recurso.

Las rutas habilitadas son las siguientes:

`-> /status`

Habilitada con el método GET. Devuelve una respuesta con código de estado 200 y una clave "Estado" en el cuerpo con el contenido "Todo OK!".

Ejemplo de uso:

``` 
GET /status
```
---

`-> /recibos`

Habilitada con los métodos GET y POST. 

Con GET devuelve todos los recibos registrados en el sistema.

Ejemplo de uso:

``` 
GET /recibos
```

Con POST permite crear un nuevo recibo. Para ello hay que incluir en el cuerpo de la petición el valor de dos parámetros:"usuario", indicando el valor del usuario asociado al recibo, y "textoRecibo", contiendo el texto en plano con el contenido del recibo. Se insertará entonces un nuevo recibo que tendrá como usuario el indicado y generado a partir del contenido del texto plano. Se trata de la única manera de insertar recibos en la aplicación y está directamente motivada por la historia de usuario #2.

Ejemplo de uso:

```
POST /recibos
```
(Incluyendo en el cuerpo de la petición los parámetros indicados.)

---

`-> /recibos/{idR}`

Habilitada con los métodos GET y DELETE. Proporciona acceso al recibo de id *idR*, para devolverlo o eliminarlo respectivamente.

Ejemplos de uso:

```
GET /recibos/123
```
```
DELETE /recibos/321
```
---

`-> /recibos/{idR}/articulos`

Habilitada con el método GET. Devuelve los artículos del recibo con id *idR*.

Ejemplo de uso:

```
GET /recibos/123/articulos
```
---

`-> /recibos/{idR}/articulos/{idA}`

Habilitada con los métodos GET y PATCH. 

Con GET devuelve el artículo con id *idA* del recibo con id *idR*.

Ejemplo de uso:

```
GET /recibos/123/articulos/1
```

Con PATCH hay que pasarle en el cuerpo de la petición el valor del parámetro "tipo". Modifica el atributo tipo del artículo con id *idA* del recibo con id *idR*. La idea es que el cálculo de los recuentos y las tendencias de compra se realizan agrupando los artículos en función de su tipo (en su defecto de su descripción), de manera que se puede asignar el mismo tipo a aquellos artículos que queremos que se consideren de manera conjunta.

Ejemplo de uso:

```
PATCH /recibos/123/articulos/1
```
(Incluyendo en el cuerpo de la petición el parámetro indicado.)
---

`-> /recuentos/{periodo}/{usuario}`

Habilitada con el método GET. El parámetro *periodo* puede tomar los valores *semanal* o *mensual*. Devuelve los cinco artículos en los que el usuario *usuario* ha realizado más gasto en el último mes o en la última semana respectivamente.

Ejemplo de uso:

```
GET /recuentos/semanal/amerigal
```
---

`-> /tendencias/{lugar}`

Habilitada con el método GET. Devuelve los cinco artículos en que se ha realizado mayor gasto entre todos los recibos de lugar *lugar*, con independencia de sus usuarios.

Ejemplo de uso:

```
GET /tendencias/GRANADA
```