# mvp_api

## Guia de ejecision.

<br>

Requisitos previos:
````
- Golang 1.18.0 o superior
- Mongo DB 4.0.0 o superiro
- Visual Studio Code (Preferentemente por el launch file)
````

<br>

Ejecucion VS Code:
````
1. Ir al archivo main.go dentro de VS Code y correr (Tecla F5)
2. Comenzar a mandar peticiones a 'localhost:5050'
````

<br>

Ejecucion Simple:
````
1. Ejecutar el comando 'go run main.go' dentro de este directorio
2. Comenzar a mandar peticiones a 'localhost:5050'
````

<br>

Pruebas:
````
Usar la coleccion de postman dentro de este directorio

Creacion de Usuario (Create User):
 - Crear usuario con "role:Admin" para tener acceso a todos los endpoints

Login (Login):
 - Hacer Login con email i/o id de mongo del usuario y password en foma de query params
 - Usar el jwt creado para el resto de peticiones como header

Consulta de Usuario (Get User):
 - Usar jwt como header
 - Mandar como param el id de mongo del usuario a consultar

Creacion de Orden (Create Order):
 - Usando mandar la peticion y validar las reglas de negocio (Peso y clasificacion de la orden)
 - Requiere body

Consulta de Orden (Get Order):
 - Usar jwt como header
 - Mandar como param el id de mongo de la orden a consultar

Update/Cancel (Update Status):
 - Usar jwt como header
 - Mandar como query params el id de la orden y el nuevo estatus

Health check (Health):
 - Peticion para verificar estado del servicio

Health Cheack Dependencies (Health Dependencies):
 - Peticion para verificar estado del servicio y sus dependencias
````

<br>

Consideraciones:
`````
Estatus validos: 
 - "creado"
 - "recolectado"
 - "en_estacion"
 - "en_ruta"
 - "entregada"
 - "cancelada"
`````