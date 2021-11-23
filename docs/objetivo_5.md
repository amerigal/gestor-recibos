# Documentación Objetivo 5
## Elección de contenedor base para el contenedor de test

Para lanzar los tests en el contendor necesitamos disponer del lenguaje go instalado, así como del task runner Task que podemos instalar mediante go.

Es por ello que mi primera opción para elegir una imagen base ha sido dirigirme a las imágenes oficiales de Docker para go, que podemos encontrar [aquí](https://hub.docker.com/_/golang/), y tomar la versión básica basada en Debian, golang:1.17. He tomado la versión 1.17 de go por ser la última versión estable a día de hoy. Dicha imagen no presenta ningún problema en la instalación del task runner ni en la ejecución de los tests. Sin embargo, la imagen construida resultante ocupa 976MB.

En pos de construir una imagen más ligera que contenga solo lo imprescindible para la ejecución de los tests, he decidido tomar la versión basada en alpine, golang:1.17-alpine. De nuevo, esta imagen base no ha presentado ningún problema. En este caso, la imagen construida resultante se queda en 350MB. Se trata entonces de una mejora sustancial respecto de la imagen probada inicialmente. Aún así, el espacio que ocupa es relativamente grande, de modo que he tratado de encontrar el origen de esa sobrecarga.

Para ello he realizado un `du -hs /usr/local/go` en el contenedor mediante la orden

```
docker run -t -v `pwd`:/app/test --entrypoint du amerigal/gestor-recibos -hs /usr/local/go
```

lo que ha dado como resultado

```
326.0M
```
En definitiva, el espacio de la imagen está, en su mayor parte, ocupado por la instalación básica de go realizada en la construcción de golang:1.17-alpine. 

Por un lado, esto nos hace pensar que no merece la pena tratar de construir una imagen partiendo de un sistema operativo más ligero e instalando go después, pues no podríamos rebajar de ese modo los 326MB de la instalación de go.

Por otro lado, surge el planteamiento de considerar si hay archivos superfluos en la instalación básica de go. Dicho tema se trata en https://groups.google.com/g/golang-nuts/c/HwAyVo9mBHs/ y una solución está en https://gist.github.com/adg/fe4ef483fc1f74bbaa83, donde se especifican los archivos que deben mantenerse y los que son irrelevantes para una instalación "lite" de go. Sin embargo, debido a la falta de argumentación en la selección de dichos archivos y la ausencia de fuentes oficiales que proporcionen dicha versión ligera, he decidido no seguir ese camino para evitar los posibles problemas que pueda ocasionar.

En conclusión, he decido partir de la versión golang:1.17-alpine.