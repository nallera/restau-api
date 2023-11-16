# Restaurant API

### Comentarios generales sobre la solucion
Inicialmente, para evitar consultar constantemente a la fuente web (ya que al cambiar cada 6 horas seria muy ineficiente), implemente un cache  externo (particularmente usando Redis). Esta solucion no cumplia con el requerimiento de tiempo, pero deje esa implementacion en el codigo por si es de interes (internal/restaurant/adapter/repository/cache.go).

La solucion final esta comprendida en una unica aplicacion de Golang. incluye un cache en memoria (implementacion de la interfaz **RestaurantCache**) que consulta contra el servicio web cada 5 minutos (establecido en el archivo de configuracion yaml como _cache_ttl_seconds_). La mayor parte de esas llamadas no ejecutan ni siquiera un request, ya que al traer por primera vez la lista de restaurants, se obtiene tambien la metadata del archivo (que se almacena en una implementacion de la interfaz **WebSourceMetadata**). Esto permite, usando los atributos  eTag y LastModified, solo consultar al servicio web por cambios a partir de las 6 horas (nuevamente, establecido en el archivo de configuracion yaml como _refresh_period_seconds_).   

### Ejecucion
Para ejecutar el programa, al estar contenido totalmente en la aplicacion Golang, existen dos opciones. Correr en consola (desde el root del proyecto):
-  _go run cmd/app/main.go_
- _docker-compose up_ (incluir el flag _--build_ si se hacen cambios en el codigo)

El curl para el endpoint es:
```
curl --location --request GET 'http://localhost:8080/restaurants/available' \
--header 'Content-Type: application/json' \
--data '{
"latitude": XXX,
"longitude": XXX
}'
```

Para correr los tests y el benchmark, se puede utilizar:
_go test -v ./... -bench=. -benchtime=20s_ (se puede elegir la duracion del benchmarking)
