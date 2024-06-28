## Dependencias:
### para usar Router:
 go get github.com/gorilla/mux
### JWT:
 go get github.com/golang-jwt/jwt/v4
### websocket:
 go get github.com/gorilla/websocket
### variables de entorno:
 go get github.com/joho/godotenv


 Estructuras server.go:

Config: tiene las caracteristicas del servidor. El puerto en el que se va ejecutar, la clave secreta para generar tokens y la conexion a base de datos.

broker: Nos ayuda a tener varias instancias de servidor corriendo. Esta estructura a su vez tiene la estructura Config y el metodo Config, para ser de tipo Server.

Interface:

server: esta interface implementa el modelo de datos o estructura de config.


Patrones:
repository.

conceptos: 
- receive functions
- si no quiero usar valores puedo usar el guion bajo '_' ejemplo: 
```
 _, err := repo.db.ExecContext(...
    ```

recorrer arreglos:
rows.next

# construir imagen DOCKER de base de datos:

```
docker build . -t platzi-ws-rest-db
```
- ejecutar magen:
```
docker run -p 54321:54322 platzi-ws-rest-db
```

- cambiar archivo .env:

DATABASE_URL=postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable
