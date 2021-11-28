# Documentación Objetivo 6
## Elección de servicio de Integración Continua

La idea es incluir dos servicios de integración continua.

En primer lugar, he considerado el servicio Travis-CI, pues se trata de un servicio de IC con una fácil y rápida configuración que se puede sincronizar directamente con GitHub para el propósito que buscamos. El problema es que la versión de prueba gratuita que ofrece requiere la introducción de una tarjeta de crédito y los créditos gratuitos que te proporcionan son solo válidos durante 30 días. Es por ello que lo he descartado en busca de una mejor opción.

Por otro lado, he considerado Circle-CI que, al igual que Travis-CI, tiene una fácil configuración y permite la sincronización directa con GitHub. Sin embargo, en este caso la versión gratuita no tiene ningún requerimiento y los créditos que proporcionan se renuevan cada mes. Por lo tanto, he decidido trabajar con este servicio de IC.

Para la elección del segundo sistema de IC he buscado uno que permitiera una fácil sincronización con GitHub para configurar de forma secilla la Checks Api, que permita almacenar secretos para el inicio de sesión en Docker y que esté mantenido. Tanto GitHub Actions como Travis-CI y Circle-CI satisfacen las condiciones y son los contemplados por tratarse de los principales sistemas de integración continua. Travis-CI queda descartado por las razones antes indicadas y Circle-CI ya ha sido elegido como primer sistema de CI, de modo que he decidido utilizar las GitHub Actions para configurar un segundo sistema de IC.

## Tareas

Se han configurado dos tareas relativas a la Integración Continua:
- <u>Tests Unitarios</u>: Se ha configurado una GitHub Action [aquí](../.github/workflows/test.yaml) que lanza el contenedor de test desarrollado en el objetivo 5.

- <u>Test de versiones de go válidas</u>: Se ha configurado Circle-CI [aquí](../.circleci/config.yml) para lanzar un job por cada versión del lenguaje cuya validez se desea comprobar. He contemplado las versiones de go 1.16 y 1.17 dado que son las únicas dos estables actualmente. Para verificar la validez se realiza una compilación, de la que se espera que no se obtengan errores en las versiones válidas. Para configurar la ejecución de las diferentes versiones se ha definido una matriz CI (en este caso unidimensional) con una lista de las versiones a comprobar.