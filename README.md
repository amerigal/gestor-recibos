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


## Documentación adicional

- [Configuración previa de la cuenta](docs/objetivo_0.md)
- [Descripción de los usuarios y planificación del proyecto](docs/objetivo_1.md)
- [Elección del lenguaje de desarrollo y diseño inicial](docs/objetivo_2.md)
- [Elección del gestor de tareas](docs/objetivo_3.md)
