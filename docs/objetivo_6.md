# Documentación Objetivo 6
## Elección de servicio de Integración Continua

La idea es incluir un servicio de integración continua independiente de las GitHub Actions de las que ya hace uso el proyecto.

En primer lugar, he considerado el servicio Travis-CI, pues se trata de un servicio de IC con una fácil y rápida configuración que se puede sincronizar directamente con GitHub para el propósito que buscamos. El problema es que la versión de prueba gratuita que ofrece requiere la introducción de una tarjeta de crédito y los créditos gratuitos que te proporcionan son solo válidos durante 30 días. Es por ello que he lo he descartado en busca de una mejor opción.

Por otro lado, he considerado Circle-CI que, al igual que Travis-CI, tiene una fácil configuración y permite la sincronización directa con GitHub. Sin embargo, en este caso la versión gratuita no tiene ningún requerimiento y los créditos que proporcionan se renuevan cada mes. Por lo tanto, he decidido trabajar con este servicio de IC.