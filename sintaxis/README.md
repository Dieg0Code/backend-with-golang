# Introducción

Go o Golang es un lenguaje de programación creado por google y publicado como código abierto en el 2012. Go fue diseñado con la idea de tener un lenguaje con rendimiento similar al de C y C++ pero con una sintaxis mas simple como la de python.

A efectos practicos go es un lenguaje compilado a binario directamente, tiene herrameias para menejo de concurrencia y es muy utilizado para el desarrollo de microservicios.

Algunos de los frameworks mas utilizados en go son:

- Gin
- Echo
- Fiber
- Beego
- Buffalo

## Sintaxis

Para crear un hola mundo en go tenemos que crear un archivo con la extensiòn `.go` y dentro de el escribir el siguiente codigo:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hola mundo")
}
```

Todos los archivos estan contenidos en un `package`, tambien debemosm importar `fmt` para poder utilizar la funcion `Println` la cual muestra por consola el mensaje que le pasemos como argumento.

Las funciones se crean con la palabra reservada `func` seguida del nombre de la funcion y los argumentos que recibe entre parentesis, en este caso la funcion `main` no recibe ningun argumento.

Para ejecutar el programa tenemos que ejecutar el siguiente comando:

```bash
go run main.go
```

## Modulos

Go tiene un sistema de modulos para manejar las dependencias de nuestros proyectos, para inicializar un modulo tenemos que ejecutar el siguiente comando:

```bash
go mod init <nombre del modulo>
```

De esta forma podemos importar paquetes de terceros en nuestro proyecto.

## Variables

Las variables se declaran usando la palabra reservada `var` seguido del nombre de esta y su tipo de dato, por ejemplo:

```go
var nombre string = "Pedro"
```

Go es un lenguaje tipado, los tipos de datos basicos son:

- string
- int
- float
- bool

Tambien podemos declarar variables sin especificar el tipo de dato, en este caso go infiere el tipo de dato de la variable:

```go
var nombre = "Pedro"
```

En Go tambien tenemos un operador especial `:=` con el que podemos declarar una variable e inicializarla de inmediato:

```go
nombre := "Pedro"
```

Esto es una forma abreviada de declarar los mismo que `var` `nombre` `string` `=` `"Pedro"`, este operador solo se puede utilizar dentro de funciones.

## Constantes

Las constantes se declaran con la palabra reservada `const` seguido del nombre de la constante y su valor, por ejemplo:

```go
const nombre string = "Pedro"
```

Como toda constante su valor no puede ser modificado.

## Control de flujo

Con el control de flujo podemos definir lógica para que nuestro programa se comporte de una forma u otra dependiendo de ciertas condiciones.

### If

```go
var edad int = 18

if edad >= 18 {
    fmt.Println("Eres mayor de edad")
} else if edad == 17 {
    fmt.Println("Casi eres mayor de edad")
}else {
    fmt.Println("Eres menor de edad")
}
```

Tambien tenemos los operadores `&&` y `||` para evaluar condiciones multiples:

```go
var edad int = 18
var nombre string = "Pedro"

if edad >= 18 && nombre == "Pedro" {
    fmt.Println("Eres mayor de edad y te llamas Pedro")
}
```

## Estructuras de datos

### Arrays

Para declarar un array tenemos que especificar el tipo de dato que va a contener y la cantidad de elementos que va a tener, por ejemplo:

```go
var numeros [3]int
```

Luego para añadir elementos a nuestro array tenemos que hacerlo de la siguiente forma:

```go
numeros[0] = 1
numeros[1] = 2
numeros[2] = 3
```

Si no añadimos elementos Go por defecto llenara nuestro array con el valor 0 en este casi ya que es un array de enteros.

Tambien podemos declarar un array con elementos de inmediato:

```go
var numeros = [3]int{1, 2, 3}
```

## Maps

Un map es una estructura de datos que nos permite almacenar datos en pares clave-valor, para declarar un map tenemos que especificar el tipo de dato de la clave y el valor, por ejemplo:

```go
myMap := make(map[string]int)
myMap["llave"] = 1
myMap["llave2"] = 2
```

Tambien podemos declarar un map con elementos de inmediato:

```go
myMap := map[string]int{"llave": 1, "llave2": 2}
```

## Bucles

### For

Para crear un bucle for tenemos que especificar una condicion, por ejemplo:

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

## Funciones

Las funciones se declaran con la palabra reservada `func` seguido del nombre de la funcion y los argumentos que recibe entre parentesis, por ejemplo:

```go
func suma(a int, b int) int {
    return a + b
}
```

Tambien podemos declarar funciones que retornen multiples valores:

```go
func suma(a int, b int) (int, int) {
    return a + b, a - b
}
```

## Structs

En Go no existen las clases como tal, en vez de eso tenemos las structs las cuales nos permiten definir multiples campos de datos, por ejemplo:

```go
type Persona struct {
    nombre string
    edad int
}
```

Para crear una instancia de una struct tenemos que hacerlo de la siguiente forma:

```go
persona := Persona{nombre: "Pedro", edad: 18}
```

## POO en Go

Go no es un lenguaje orientado a objetos, pero podemos simular la POO utilizando structs y metodos, por ejemplo:

```go
type Persona struct {
    nombre string
    edad int
}

func (p Persona) saludar() {
    fmt.Println("Hola soy", p.nombre)
}

func main() {
    persona := Persona{nombre: "Pedro", edad: 18}
    persona.saludar()
}
```