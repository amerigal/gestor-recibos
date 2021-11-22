# Documentación Objetivo 4
## Desarrollo de tests unitarios  

Cabe mencionar en primer lugar que la biblioteca estándar de go no dispone de ningún paquete de aserciones y para no introducir dependencias externas en una etapa tan temprana, he decidido no utilizarlas para desarrollar los tests unitarios.

Por otro lado, se ha seguido el estándar de go para la realización de test reflejado [aquí](https://golang.org/doc/tutorial/add-a-test). En ese sentido, se ha hecho uso de la biblioteca testing incorporada en go y se han creado archivos articulo_test.go y recibo_test.go que testean correspondientemente a articulo.go y a recibo.go. Para la validación de los distintos atributos se ha hecho uso de la biblioteca errors también incorporada en go.

En definitiva, no se han tenido que tomar importantes decisiones técnicas respecto a los tests unitarios debido a la existencia de un estándar marcado.