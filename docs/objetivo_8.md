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
