# Credit-Assigment-REST-API

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

Asignacion de inversion (/credit-assigment):
 - Requiere body:
    {"original_investment":3200.0}

Consulta de asignacion (/credit-assigment/:id):
 - Mandar como param el id de mongo de la asignacion a consultar

Consulta de estadisticas (/statistics):
 - Mandar un GET simple

Health check (Health):
 - Peticion para verificar estado del servicio

Health Cheack Dependencies (Health Dependencies):
 - Peticion para verificar estado del servicio y sus dependencias
````