# Proyecto Infraestructura Virtual

## Descripción del proyecto

Desarrollo de un gestor de gastos. Se partirá de recibos de compra en formato texto. La aplicación reconocerá los distintos artículos del recibo y el gasto realizado en ellos. De esta manera, se llevará un control del gasto acumulado en los distintos artículos a lo largo del tiempo y se generarán estadísticas al respecto. Además se producirán informes con los artículos más comprados entre todos los usuarios.

## Guía de uso

Lo primero es instalar este repositorio, bien descargándolo directamente en la interfaz web o clonándolo por línea de comandos.

Es también necesario tener instalado [go](https://golang.org/), cuyas instrucciones de descarga e instalación se pueden encontrar [aquí](https://golang.org/doc/install).

Asimismo, se requiere tener instalado el gestor de tareas [Task](https://taskfile.dev/#/), que se puede obtener mediante el siguiente comando:

```shell
go install github.com/go-task/task/v3/cmd/task@latest
```

Una vez hecho esto, podemos ya instalar las dependencias del proyecto con la orden:
```shell
task installdeps
```
A continuación, se puede revisar la correcta sintaxis del código mediante:
```shell
task check
```
También es posible lanzar los tests que comprueban el correcto funcionamiento del sistema:
```shell
task test
```

## Contenedor

Este proyecto dispone de un contenedor que se actualiza automáticamente en Docker Hub cada vez que se hace un push a main. Por el momento, dicho contenedor permite, una vez tenemos descargado el repositorio, el lanzamiento de los tests mediante la orden:
```
task docker
```
Además, se ha configurado el proyecto mediante Circle-CI para que se lance el contenedor de tests automáticamente al incorporar código.

## Documentación adicional

- [Configuración previa de la cuenta](docs/objetivo_0.md)
- [Descripción de los usuarios y planificación del proyecto](docs/objetivo_1.md)
- [Elección del lenguaje de desarrollo y diseño inicial](docs/objetivo_2.md)
- [Elección del gestor de tareas](docs/objetivo_3.md)
- [Tests unitarios](docs/objetivo_4.md)
- [Elección de contenedor base](docs/objetivo_5.md)
- [Elección de servicio de IC](docs/objetivo_6.md)
- [Servicio de logging](docs/objetivo_7.md)
- [Servicio de configuración](docs/objetivo_7.md#Servicio-de-configuración)
- [Diseño de API](docs/objetivo_8.md)
- [Implementación de API](docs/objetivo_9.md)