# Documentación Objetivo 7
## Servicio de Logging

Uno de los servicios esenciales en una aplicación es el de logging, consistente en el registro de toda actividad relevante del sistema. En el caso de go hay muchos paquetes que proporcionan este servicio, de modo que para elegir uno es conveniente primero precisar lo que esperamos del paquete. 

Lo primero y más básico que necesitamos es la posibilidad de modificar la salida, para que podamos determinar si queremos que el log escriba en la consola o sobre un fichero. También, requerimos que sea un log con niveles, de manera que podamos después elegir qué niveles de mensaje queremos que se muestren, desde todos los mensajes incluyendo los de depuración hasta restringirnos a que solo se muestren los mensajes críticos que provocan la terminación de la aplicación.

Por otro lado, siempre es preferible reducir las dependencias de la aplicación. En caso de añadir una nueva dependencia, debe ser de un paquete de confianza, mantenido y que se adapte al progreso del lenguaje de programación.

Bajo dicho criterio procedo a analizar las principales opciones de servicios de logging que están disponibles para go:

- [Log](https://pkg.go.dev/log): Paquete de Log de la biblioteca estándar de go. Se trata de tipo Log muy sencillo de usar, que permite formatear la salida y que ya está integrado en go, de manera que nos ahorraría introducir dependencias. Sin embargo, se trata de un paquete muy limitado que no permite el uso de niveles. Es por ello que lo descarto.

- [Logrus](https://github.com/Sirupsen/logrus): Se trata de uno de los paquetes de Log más extendidos en go y que ha contribuido a la popularidad de los Log estructurados en el lenguaje. Se trata de un paquete muy completo, con un Log que permite el formateo de la salida, estructurado (permite mostrar la información de manera estructurada en formatos como json) y con niveles. Su problema es que, tal y como indica en su [repositorio de GitHub](https://github.com/Sirupsen/logrus), se encuentra en modo mantenimiento para evitar ocasionar problemas a aquellos que ya hacen uso del paquete, de modo que no experimentará avances a medida que el lenguaje cambie. En este sentido, aunque su funcionalidad resolvería sin problema lo que se pretende conseguir en el proyecto, creo que es conveniente optar por algún paquete que siga el avance del lenguaje.

- [Glog](https://github.com/golang/glog): Versión de código abierto del paquete de Log usado por Google y que está implementado al estilo del paquete de Log de Google de C++ ([glog](https://github.com/google/glog)). Se trata también de un Log con niveles que resuelve el problema, pero tal y como indica el repositorio no se encuentra bajo desarrollo.

- [Zerolog](https://github.com/rs/zerolog): Es uno de los paquetes de Log recomendados por los desarrolladores de Logrus como una de las opciones modernas de Log estructurados. Es también muy completo, con formateo de salida, niveles y eficiencia. Se trata de un buen candidato.

- [Zap](https://github.com/uber-go/zap): Es otro de los paquetes de Log recomendados por los desarrolladores de Logrus. Al igual que Zerolog, es estructurado, con niveles, con muchas opciones de configuración de salida y muy eficiente. Se trata de otro buen candidato.

Analizando entre Zerolog y Zap, la API de Zap me parece más cómoda, pues presenta la opción de llamar a las funciones al estilo de printf o de manera estructurada incluyendo los campos como parámetros, mientras que Zerolog realiza las llamadas a funciones de manera encadenada del tipo:
```
log.Info().Str("var",value).Msg("Esto es un mensaje de log")
```
En este sentido, he optado por realizar la implementación con Zap y dado que su configuración ha sido sencilla y ha funcionado correctamente, he decidido quedarme con este.

Cabe mencionar que Zap presenta dos versiones de Log, una que se utiliza cuando la eficiencia es crítica y que solo permite logging estructurado, y otra que aunque menos eficiente sigue siendo razonablemente rápida y permite mayor flexibilidad con mensajes al estilo de printf. Es por esto último que he decidido usar la segunda versión, llamada SugaredLogger.
