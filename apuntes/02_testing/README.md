# Testing en Go

Go a diferencia de otros lenguajes, viene con un sistema de testing integrado en el lenguaje de forma nativa, a diferencia de otros que se apoyan en librerías externas.

Los test nos sirven para probar partes del código, para verificar que el resultado de cierto fragmento de lógica devuelva resultados esperados.

## Tipos de test

Los test tienen tres tipos o tres capas de testeo, van desde una capa mas general que se encarga de probar funcionalidades completas, una capa intermedia que prueba la integración de los componentes y una capa mas baja que se encarga de probar funciones o métodos individuales.

- **Test funcional**: Los test funcionales se encargan de probar la funcionalidad completa de la aplicación. Por ejemplo, si tenemos una aplicación que se encarga de gestionar usuarios, un test funcional se encargaría de probar que la aplicación puede crear, leer, actualizar y eliminar usuarios correctamente.
- **Test de integración**: Los test de integración se encargan de probar la integración de los componentes de la aplicación. Por ejemplo, si tenemos una aplicación que se conecta a una base de datos, un test de integración se encargaría de probar que la aplicación se conecta correctamente a la base de datos y que puede realizar operaciones de lectura y escritura.
- **Test unitario**: Los test unitarios son en términos burdos, funciones que prueban funciones. Se encargan de probar funciones o métodos individuales, asegurándose de que dado x o y parámetros de entrada se espera que la función devuelva un resultado esperado. Esto es util cuando se trabaja en un equipo grande, ya que si alguien modifica una función, los test unitarios se encargan de verificar que la función sigue devolviendo el resultado esperado para de esta forma alertar al desarrollado en caso de que algún cambio rompió una parte de la aplicación.

Estoy tres tipos de testing se suelen combinar para tener una cobertura completa de la aplicación, se suele decir que **un 80% de test deben ser unitarios, un 15% de test de integración y un 5% de test funcionales**.

### Test unitarios

Para crear tests unitarios en Go, se crea un archivo con el nombre del archivo que se quiere testear seguido de `_test.go`, por ejemplo si tenemos un archivo `main.go` y queremos testearlo, creamos un archivo `main_test.go`.

Por ejemplo dada una función:

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

En donde tenemos dos caminos posibles, uno en donde el usuario existe y otro en donde el usuario no existe, podemos crear un test unitario para probar ambos casos:

```go
package domain

import "testing"

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	if user != nil {
		t.Error("We were not expecting a user with id 0")
	}

	if err == nil {
		t.Error("We were expecting an error when user id is 0")
	}
}

func TestGetUserUserFound(t *testing.T) {
    user, err := GetUser(123)

    if user == nil {
        t.Error("We were expecting a user with id 123")
    }

    if err != nil {
        t.Error("We were not expecting an error when user id is 123")
    }
}
```

Para correr los test se pueden ejecutar varios comandos, todos con diferentes propósitos:

- `go test`: Corre todos los test en el directorio actual.
- `go test -v`: Corre todos los test en el directorio actual y muestra el output de cada test.
- `go test -run TestGetUserNoUserFound`: Corre un test en particular.
- `go test -cover`: Muestra el porcentaje de cobertura de los test.
- `go test -coverprofile cover.out`: Guarda el porcentaje de cobertura de los test en un archivo.
- `go tool cover -html=cover.out`: Muestra un reporte de cobertura en un navegador.

Algo particular de Go es que los test del core no tienen asserts, por lo que se debe hacer manualmente, por ejemplo con `if` y `t.Error`, esto es porque cuando se usan los asserts, si un assert falla, el test se detiene y no se ejecutan los asserts siguientes, por lo que no se puede saber si hay mas errores en el test. Pero existe una librería llamada `testify` que provee asserts para Go que no tiene este problema, si un assert falla, el test sigue corriendo y se ejecutan los asserts siguientes.

```go
package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "We were not expecting a user with id 0")
	assert.NotNil(t, err, "We were expecting an error when user id is 0")
	assert.EqualValues(t, http.StatusNotFound, err.Status, "We were expecting 404 when user is not found")
	assert.EqualValues(t, "not_found", err.Code, "We were expecting 'not_found' error code when user is not found")
}

func TestGerUserNoError(t *testing.T) {

	user, err := GetUser(123)

	assert.Nil(t, err, "We were not expecting an error when user id is 123")
	assert.NotNil(t, user, "We were expecting a user with id 123")
	assert.EqualValues(t, 123, user.Id, "We were expecting a user with id 123")
	assert.EqualValues(t, "Diego", user.FirstName, "We were expecting a user with first name Diego")
	assert.EqualValues(t, "Obando", user.LastName, "We were expecting a user with last name Obando")
	assert.EqualValues(t, "example@email.com", user.Email, "We were expecting a user with email")
}
```

Los test tienen 3 partes, se suelen llamar:

- **Given**: Se encarga de preparar el entorno para el test.
- **When**: Se encarga de ejecutar la función que se quiere testear.
- **Then**: Se encarga de verificar que el resultado de la función es el esperado.

También estas partes se suelen llamar como `Arrange`, `Act` y `Assert` o `Initialization`, `Execution` y `Validation`, o de varias otras forma, pero la idea es la misma, preparar el entorno, ejecutar la función y verificar el resultado.

### Benchmark

Los benchmarks son test que se encargan de medir el rendimiento de una función.

Por ejemplo dado una función:

```go
package utils

// []int {9, 7, 5, 3, 1, 2, 4, 6, 8}
// []int {1, 2, 3, 4, 5, 6, 7, 8, 9}
func BubbleSort(elements []int) []int {

	keepRuning := true

	for keepRuning {
		keepRuning = false

		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepRuning = true
			}
		}
	}

	return elements
}
```

En donde estamos ordenando un slice de enteros, podemos crear un benchmark para medir el rendimiento de la función:

```go
package utils

import (
    "fmt"
    "testing"
)

func BenchmarkBubbleSort(b *testing.B) {
    elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8}

    for i := 0; i < b.N; i++ {
        BubbleSort(elements)
    }
}

func BenchmarkBubbleSortWorstCase(b *testing.B) {
    elements := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

    for i := 0; i < b.N; i++ {
        BubbleSort(elements)
    }
}
```

En donde `b.N` es la cantidad de veces que se va a ejecutar el benchmark, Go se encarga de ejecutar el benchmark varias veces y mostrar el promedio de las ejecuciones.

Para correr los benchmarks se puede ejecutar varios comandos, todos con diferentes propósitos:

- `go test -bench .`: Corre todos los benchmarks en el directorio actual.
- `go test -bench . -benchmem`: Corre todos los benchmarks en el directorio actual y muestra el uso de memoria.
- `go test -bench . -benchmem -cpuprofile cpu.out`: Corre todos los benchmarks en el directorio actual, muestra el uso de memoria y guarda el uso de CPU en un archivo.
- `go tool pprof cpu.out`: Muestra un reporte de uso de CPU en un navegador.

### Como estructurar artefactos y mocks en Go

Para estructurar los artefactos y mocks en Go, se puede seguir la siguiente estructura:

```
.
├── domain
│   ├── user.go
│   └── user_test.go
├── interfaces
│   ├── user.go
│   └── user_test.go
├── services
│   ├── user_service.go
│   └── user_service_test.go
└── utils
    ├── utils.go
    └── utils_test.go
```

En donde `domain` se encarga de tener los modelos de la aplicación, `interfaces` se encarga de tener las interfaces de los servicios, `services` se encarga de tener los servicios de la aplicación y `utils` se encarga de tener funciones de utilidad.

En donde los test de cada artefacto se encuentran en el mismo directorio que el artefacto, por ejemplo el test de `domain/user.go` se encuentra en `domain/user_test.go`.

Para los mocks, se pueden crear en el mismo directorio que el artefacto que se quiere mockear, por ejemplo si se quiere mockear `services/user_service.go`, se puede crear un archivo `services/user_service_mock.go`. En donde se puede crear un mock de la siguiente forma:

```go
package services

import (
    "github.com/stretchr/testify/mock"
)

type MockUserService struct {
    mock.Mock
}

func (m *MockUserService) GetUser(userId int64) (*User, *utils.ApplicationError) {
    args := m.Called(userId)

    if args.Get(0) == nil {
        return nil, args.Get(1).(*utils.ApplicationError)
    }

    return args.Get(0).(*User), nil
}
```

En donde se crea un struct que implementa la interfaz que se quiere mockear, en este caso `UserService`, y se implementa el método que se quiere mockear, en este caso `GetUser`.

Este código consiste en que cuando se llame al método `GetUser` del mock, se llame al método `Called` de `mock.Mock` con el argumento `userId`, y se verifique si el primer valor retornado es `nil`, si es `nil` se retorna `nil` y el segundo valor retornado, si no es `nil` se retorna el primer valor retornado y `nil`.

Para usar el mock, se puede hacer de la siguiente forma:

```go
package services

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestGetUserNotFoundInDatabase(t *testing.T) {
    mockRepo := new(MockUserService)
    mockRepo.On("GetUser", int64(0)).Return(nil, &utils.ApplicationError{Message: "user 0 was not found", Status: http.StatusNotFound, Code: "not_found"})

    service := NewUserService(mockRepo)

    user, err := service.GetUser(0)

    assert.Nil(t, user)
    assert.NotNil(t, err)
    assert.EqualValues(t, http.StatusNotFound, err.Status)
    assert.EqualValues(t, "not_found", err.Code)
}
```

En donde se crea un mock del servicio, se le dice que cuando se llame al método `GetUser` con el argumento `0`, retorne `nil` y un error, se crea una instancia del servicio con el mock y se llama al método `GetUser` con el argumento `0`, se verifica que el usuario sea `nil`, que el error no sea `nil`, que el status del error sea `404` y que el código del error sea `not_found`.