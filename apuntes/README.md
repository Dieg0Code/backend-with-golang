# Backend con Golang

## Introducción a MVC

MVC es un patrón de diseño, es decir, un conjunto de reglas y buenas prácticas que nos ayudan a organizar nuestro código de una forma más eficiente y mantenible. Consiste en dividir nuestra aplicación en tres partes:

- **Model**: Se encarga de gestionar los datos de nuestra aplicación.
- **View**: Se comunica con el controlador para obtener los datos que necesita y los muestra al usuario.
- **Controller**: El controlador se comunica con el modelo para obtener los datos o para modificarlo o guardarlos

En el contexto de una aplicación backend, el cliente hace peticiones al servidor que son recibidas por el controlador, este se comunica con el modelo para obtener los datos necesarios y los devuelve al cliente para que sean mostrados en la vista.

### Problemas con MVC

Uno de los problemas que presenta este patrón es que el controlador accede a la capa de datos directamente, sin un intermediario, lo que puede traer problemas como la duplicación de código o la dificultad de realizar pruebas unitarias, es por eso que lo frameworks que implementan este patrón usan una capa intermedia entre el controlador y el modelo llamada **servicio**.

### Capa Service

La capa de servicio contiene toda la lógica de negocio, es decir, todas las operaciones que se pueden realizar con los datos de la aplicación. Esta capa se encarga de recibir los datos del controlador, procesarlos y enviarlos al modelo para que sean guardados o devolverlos al controlador para que sean mostrados en la vista.

### Controladores

- Son el punto de entrada de cualquier aplicación.
- Son accedidos via URL mappings. Ej `http://localhost:8080/api/v1/users`
- Solo validan las peticiones que tienen todos los parámetros requeridos.
- No contienen ninguna lógica de negocio.
- Se apoyan de los servicios para procesar cada petición.
- Devuelven la respuesta al cliente sin agregar datos adicionales.

### Servicios

- Contienen toda la lógica de negocio.
- Cada servicio es responsable de manejar una sola entidad (user service, item service, etc).
- Son ``stateless``, no guardan información entre peticiones.
- Son usualmente `singletons`, es decir, solo existe una instancia de cada servicio en la aplicación.
- Invocan otros servicios, modelos, proveedores externos y cualquier otra fuente de datos que sea necesaria para completar su tarea.
- Manejan errores, envían métricas, logs, tags y cualquier otra métrica soportada que necesite nuestra aplicación.

### Model / Domain / DAOs

- Maneja el dominio central de los servicios. Cualquier otra capa existe para soportar y servir estos objetos de dominio, en otras palabras, los modelos son el núcleo de la aplicación.
- Es la encargada de definir la estructura de los objetos de dominio.
- Es la única capa que tiene acceso a la persistencia. Es la única que sabe donde y como persistir los datos.
- Es la encargada de abstraer la lógica de persistencia creando una interfaz genérica.

## Ejemplo de MVC

Vamos a crear una aplicación que implemente el patrón MVC. La aplicación es una API básica que retorna un usuario en formato JSON. Para ello vamos a crear tres paquetes:

- **domain**: Que contiene la estructura del usuario y su DAO.
- **service**: Que contiene la lógica de negocio.
- **controller**: Que contiene el controlador de la aplicación.
- **utils**: Que contiene funciones de utilidad.

### Estructura del proyecto

```bash
.
├── domain
│   ├── user.go
│   └── user_dao.go
├── service
│   └── user_service.go
├── controller
│   └── user_controller.go
├── utils
│   └── errors_utils.go
└── main.go

```

### Creando el modelo

Primero vamos a crear el modelo del usuario en el paquete `domain/user.go`:

```go
package domain

type User struct {
	Id        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
```

La estructura `User` tiene cuatro campos: `Id`, `FirstName`, `LastName` y `Email`, los cuales son exportados para que puedan ser serializados a JSON.

Ahora vamos a crear el DAO del usuario en el paquete `domain/user_dao.go`:

```go
package domain

import (
	"fmt"
	"net/http"

	"github.com/dieg0code/go-microservices/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Fede", LastName: "Leon", Email: "example@email.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("user %v was not found", userId),
		Status:  http.StatusNotFound,
		Code:    "not_found",
	}
}
```

El DAO o **Data Access Object** es el encargado de interactuar con la base de datos o cualquier otra fuente de datos. En este caso, el DAO es un mapa que contiene un usuario de ejemplo. La función `GetUser` recibe un `userId` y retorna el usuario correspondiente o un error si no se encuentra el usuario.

### Creando el servicio

Ahora vamos a crear el servicio del usuario en el paquete `service/user_service.go`:

```go
package services

import (
	"github.com/dieg0code/go-microservices/domain"
	"github.com/dieg0code/go-microservices/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
```

El servicio es el encargado de contener la lógica de negocio de la aplicación. En este caso, el servicio solo llama a la función `GetUser` del DAO.

### Creando el controlador

Por último, vamos a crear el controlador del usuario en el paquete `controller/user_controller.go`:

```go
package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dieg0code/go-microservices/services"
	"github.com/dieg0code/go-microservices/utils"
)

func GetUser(res http.ResponseWriter, req *http.Request) {

	userIdParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "user_id must be a number",
			Status:  http.StatusBadRequest,
			Code:    "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.Status)
		res.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil {

		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.Status)
		res.Write([]byte(jsonValue))
		return
	}

	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)

}
```

El controlador es el encargado de recibir las peticiones del cliente, procesarlas y devolver una respuesta. En este caso, el controlador recibe un `userId` como parámetro, llama al servicio para obtener el usuario correspondiente y devuelve el usuario en formato JSON.

### Creando el punto de entrada

También tenemos la carpeta `app` que contiene el archivo `app.go` que es el punto de entrada de nuestra aplicación:

```go
package app

import (
	"net/http"

	"github.com/dieg0code/go-microservices/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
```

El archivo `app.go` contiene la función `StartApp` que inicia el servidor y mapea la URL `/users` al controlador `GetUser`.

Esta función se llama desde el archivo `main.go`:

```go
package main

import "github.com/dieg0code/go-microservices/app"

func main() {
	app.StartApp()
}
```

En este ejemplo simple, vemos como se implemente el patrón MVC en una aplicación de backend. El controlador recibe la petición del cliente, el servicio procesa la petición y el modelo contiene la estructura de los datos y la lógica de persistencia.

Go tiene la ventaja de que la librería estándar es muy completa y no necesitamos instalar dependencias adicionales para crear una aplicación de backend. En este ejemplo, usamos solo la librería `net/http` para crear un servidor web.