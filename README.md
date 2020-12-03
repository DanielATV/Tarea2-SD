# Tarea2-SD

repo: https://github.com/DanielATV/Tarea2-SD

### Integrantes

Gonzalo Larrain 201673516-K
Daniel Toro 201673595-K

### Instrucciones de uso

Para ejecutar la tarea es necesario ejecutar make build en cada una de 
las maquinas vituales. Luego, se debe utilizar make nodo en las respectivas maquinas y se deben ingresar los siguiente valores para la Id:

dist01 VM1 DataNode1 Id = 1
dist02 VM2 DataNode2 Id = 2
dist03 VM3 DataNode3 Id = 3
dist04 VM4 NameNode  Id = 0

Finalmente en cualquiera de las máquinas se debe ejecutar make cliente para correr el cliente.

#### DataNodes

El DataNode pregunta en el modo que desea ejetuarse y la Id que deben ser la que se indico anteriormente.

#### NameNode

El NameNode pregunta el modo en que debe ejecutarse y la Id que debe ser la que se indico anteriormente.

#### Cliente

El cliente tiene dos modos de funcionamiento, Uploader y Downloader. El Uploader solo puede subir una lista predetermianda de libros. Hace una conexion aleatorio a un DataNode al quere subir un libro, en el caso de fallar (si es que esta caido) se debe ejecutar de nuevo el comando make hasta que se conecte con uno que este activo.

#### Consideraciones
- El firewall debe estar desactivado en todas las maquinas virtuales
- No se considero el caso que un nodo que mande una propuesta deje de funcionar o que ningun nodo funcione.
- La estructura del log es nombre identificador. El identificador se mapea a la maquina correspondiente.
- Solo se pueden cargar los libros que se muestran al cliente y estos deben de estar en la carpeta libros.
- El NameNode no mantiene la lista de libros si cierra y se vuelve a ejecutar.
- El el archivo log del NameNode se mantiene, pero por implentacion la ubicacion esta en memoria, por lo que se pierde al cerrar el nodo.
-Todos los nodos deben estar en el mismo modo de ejecucion(algoritmo)


