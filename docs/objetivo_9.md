# Documentación Objetivo 9
## Implementación API REST

### Conversión a json

Una de las tareas a la hora de procesar las peticiones HTTP es convertir las estructuras a devolver por la API al formato json. Para ello, go proporciona el paquete "encoding/json" que permite dicha conversión automáticamente a partir de un struct o array de structs a convertir. El problema es que solo trabaja con los atributos exportados de los struct y en este proyecto son muchos los atributos no exportados dado que no se consideró necesario. Además hay objetos como los recibos en los que el id no está almacendo en el propio recibo, sino que es el Handler el que los maneja. Se hace necesario entonces crear unos tipos nuevos con todos sus atributos exportados y con toda la información que se quiere que sea devuelta por la API. Esta problemática queda recogida en el issue #47 y está resuelta con los tipos especificados en [apitypes.go](pkg/recibo/apitypes.go).


### Códigos de estado devueltos

Previo a la implementación, ha sido necesario determinar los códigos de estado a devolver como respuesta de las distintas peticiones. Se han utilizado los siguientes.

- 200 (OK): Correcto procesamiento de la petición en los métodos GET y en el método PATCH (para cambiar el tipo de un artículo).

- 201 (Created): Correcta creación del recurso en respuesta al método POST (para los recibos).

- 204 (No Content): Correcto procesamiento de la petición DELETE (para los recibos). Indica que el cuerpo de la respuesta está vacío.

- 400 (Bad Request): Error por ausencia o invalidez de parámetros en las peticiones.

- 404 (Not Found): Error por recurso no encontrado. Empleado en las peticiones con GET, PATCH y DELETE*.

\*Hay cierto debate sobre el uso del código 404 en respuesta a las peticiones DELETE debido a su idempotencia. En el [RFC 7231](https://www.rfc-editor.org/rfc/rfc7231#section-4.2.2) se aclara está cuestión, indicando que la idempotencia se refiere a que el efecto producido sobre el servidor por múltiples peticiones es el mismo que el producido por una única petición. Sin embargo, no implica que todas las peticiones deban tener la misma respuesta, como aclara después: "It knows that repeating the request will have the same intended effect, even if the original request succeeded, though the response might differ". Es por ello que he utilizado el código 404 como respuesta a las peticiones DELETE cuando el recurso a eliminar no existe.

### Inserción de recibos

El único manejador cuya implementación tiene un mayor interés es el que procesa la inserción de recibos con el método POST. Este recibe el contenido del recibo a insertar en el cuerpo de la petición HTTP y crea un archivo con dicho contenido en el directorio /tmp y con un nombre construido a partir del nombre del usuario y de la fecha y hora en que se realiza la inserción. Es dicho archivo el que procesa la función leerRecibo implementada en el comienzo del proyecto.

### Tests

He incluido en el archivo [api_test.go](pkg/recibo/api_test.go) todos los tests que verifican el correcto funcionamiento de la API. Como ya indiqué en la documentación del objetivo anterior, he hecho uso del paquete httpexpect para realizar las peticiones y aserciones sobre las respuestas. Además, he hecho uso del paquete httptest de la biblioteca estándar, que permite crear y levantar servidores de prueba a los que mandar las peticiones, sin necesidad de levantar de forma externa un servidor en un puerto concreto. Cabe decir que para asegurar la independencia de los tests y que el servidor quede correctamente cerrado al finalizar los tests, se levanta un servidor nuevo en cada test en vez de usar uno global para todos los tests. Asimismo, he incluido una función auxiliar que inserta un recibo de prueba para tener contenido sobre el que trabajar con las peticiones GET, PATCH y DELETE.

### Integración de logging

Resulta importante el registro de la información asociada a las peticiones que recibe el servidor. La manera más limpia de realizarlo es a través de una función de middleware que actúe depués de que los manejadores específicos hayan procesado la petición correspondiente. Esta función de middleware tiene acceso a los objetos http.ResponseWriter y http.Request que utilizan los manejadores específicos y a partir del objeto http.Request se puede obtener el método y la URI de la petición. Pensé que sería también interesante incluir en el logging el código de respuesta de la petición. El problema es que el objeto http.ResponseWriter no exporta dicho atributo. Una de las maneras más estándar en las que este problema ha sido resuelto por los usuarios de go es mediante la implementación de un wrapper del http.ResponseWriter que modifique el método WriteHeader para que se guarde el código de estado en un atributo cuando se escriba en la cabecera. Es este enfoque el que he seguido implementando el struct logHttp e integrándolo en la función LogHttpMiddleware.