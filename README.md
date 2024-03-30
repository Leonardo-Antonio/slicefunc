# Slice Funcs

### Description 

- Español:
"SliceFunc" es una librería de Go diseñada para proporcionar funcionalidades de programación funcional para “slices”. Ofrece métodos similares a los de JavaScript como .map, .filter, .sort, .flat, groups, .reduce, etc. Permitiendo a los desarrolladores de Go manipular “slices” de manera eficiente y legible. Ya sea que necesites transformar, filtrar o reducir datos, “SliceFunc” tiene las herramientas que necesitas para hacer el trabajo de manera rápida y sencilla. ¡Haz que tu código de Go sea más limpio y conciso con “SliceFunc”!

- English
SliceFunc" is a Go library designed to provide functional programming functionalities for slices. It offers methods similar to those in JavaScript such as .map, .filter, .sort, .flat, groups, .reduce, etc. Allowing Go developers to efficiently and readably manipulate slices. Whether you need to transform, filter, or reduce data, "SliceFunc" has the tools you need to get the job done quickly and easily. Make your Go code cleaner and more concise with "SliceFunc"!

### Use
`go get github.com/Leonardo-Antonio/slicefunc`

### Examples

#### Filter
  - **English**: Create a new slice with all elements that pass a test implemented by the provided function. It's useful for selecting elements from an array that meet certain conditions.
  - **Español**: Crea un nuevo slice con todos los elementos que pasan una prueba implementada por la función proporcionada. Es útil para seleccionar elementos de un arreglo que cumplen ciertas condiciones.
```go
package main

import (
	"fmt"

	"github.com/Leonardo-Antonio/slicefunc"
)

type Person struct {
	Name     string
	Id       uint32
	LastName string
	FullName string
}

var persons []Person = []Person{
	{Name: "leonardo", Id: 1, LastName: "Tonson"},
	{Name: "Oppo", Id: 2, LastName: "Kicju"},
	{Name: "Tans", Id: 1, LastName: "Tonson"},
	{Name: "Ruy", Id: 4, LastName: "Itachi"},
}

func main() {
	response := slicefunc.Filter(persons, func(p Person) bool {
		return p.Id == 1
	})

	fmt.Println(response) // [{leonardo 1 Tonson } {Tans 1 Tonson }]
}
```

#### Map
  - **English**: Takes an slice, applies a function to each element, and returns a new slice with the results. It’s useful for transforming data without modifying the original slice.
  - **Español**: Toma un slice, aplica una función a cada elemento, y devuelve un nuevo slice con los resultados. Es útil para transformar los datos sin modificar el slice original.
```go
package main

import (
	"fmt"

	"github.com/Leonardo-Antonio/slicefunc"
)

type Person struct {
	Name     string
	Id       uint32
	LastName string
	FullName string
}

var persons []Person = []Person{
	{Name: "leonardo", Id: 1, LastName: "Tonson"},
	{Name: "Oppo", Id: 2, LastName: "Kicju"},
	{Name: "Tans", Id: 1, LastName: "Tonson"},
	{Name: "Ruy", Id: 4, LastName: "Itachi"},
}

func main() {
	response := slicefunc.Map(persons, func(p Person) Person {
		p.FullName = p.Name + " " + p.LastName
		return p
	})

	fmt.Println(response) // [{leonardo 1 Tonson leonardo Tonson} {Oppo 2 Kicju Oppo Kicju} {Tans 1 Tonson Tans Tonson} {Ruy 4 Itachi Ruy Itachi}]
}
```

#### Every
  - **English**: Reverses the order of elements in a slice in place and returns a reference to the same slice. The first element becomes the last and the last becomes the first. It's useful when you need to reverse the order of elements in a slice.
  - **Español**: Invierte el orden de los elementos de un slice en su lugar y devuelve una referencia al mismo slice. El primer elemento pasa a ser el último y el último pasa a ser el primero. Es útil cuando necesitas invertir el orden de los elementos de un slice.
```go
package main

import (
	"fmt"

	"github.com/Leonardo-Antonio/slicefunc"
)

type Person struct {
	Name     string
	Id       uint32
	LastName string
	FullName string
	Nick     string
}

var persons []Person = []Person{
	{Nick: "example", Name: "leonardo", Id: 1, LastName: "Tonson"},
	{Nick: "example", Name: "Oppo", Id: 2, LastName: "Kicju"},
	{Nick: "example", Name: "Tans", Id: 1, LastName: "Tonson"},
	{Nick: "example", Name: "Ruy", Id: 4, LastName: "Itachi"},
}

func main() {
	response := slicefunc.Every(persons, func(p Person) bool {
		return p.Nick == "example"
	})

	fmt.Println(response) // true
}
```

#### Some
  - **English**: Checks if at least one element in the slice satisfies the provided test function. It returns true if it finds an element that satisfies the test, and false if none is found. It's useful when you need to verify whether any element in a slice meets a certain condition.
  - **Español**: Comprueba si al menos un elemento del slice cumple con la función de prueba proporcionada. Retorna true si encuentra un elemento que cumple la prueba, y false si no encuentra ninguno. Es útil cuando necesitas verificar si algún elemento en un slice cumple una cierta condición.

```go
package main

import (
	"fmt"

	"github.com/Leonardo-Antonio/slicefunc"
)

type Person struct {
	Name     string
	Id       uint32
	LastName string
	FullName string
}

var persons []Person = []Person{
	{Name: "leonardo", Id: 1, LastName: "Tonson"},
	{Name: "Oppo", Id: 2, LastName: "Kicju"},
	{Name: "Tans", Id: 1, LastName: "Tonson"},
	{Name: "Ruy", Id: 4, LastName: "Itachi"},
}

func main() {
	response := slicefunc.Some(persons, func(p Person) bool {
		return p.Id == 2
	})

	fmt.Println(response) // true
}
```

#### Find
  - **English**: Takes an slice, applies a function to each element, and returns a new slice with the results. It’s useful for transforming data without modifying the original slice.
  - **Español**: Devuelve el valor del primer elemento del slice que cumple con la función de prueba proporcionada. Si ningún elemento cumple la prueba, el método devuelve nil. Es útil cuando necesitas encontrar un elemento específico en un slice
```go
package main

import (
	"fmt"

	"github.com/Leonardo-Antonio/slicefunc"
)

type Person struct {
	Name     string
	Id       uint32
	LastName string
	FullName string
}

var persons []Person = []Person{
	{Name: "leonardo", Id: 1, LastName: "Tonson"},
	{Name: "Oppo", Id: 2, LastName: "Kicju"},
	{Name: "Tans", Id: 1, LastName: "Tonson"},
	{Name: "Ruy", Id: 4, LastName: "Itachi"},
}

func main() {
	response, found := slicefunc.Find(persons, func(p Person) bool {
		return p.Id == 4
	})

	fmt.Println(response, found) // {Ruy 4 Itachi }, true
}
```

#### FindIndex
  - **English**: Returns the index of the first element in a slice that satisfies the provided test function. If no element satisfies the test, the method returns -1. It's useful when you need to find the index of a specific element in a slice.
  - **Español**: Devuelve el índice del primer elemento en un slice que cumple con la función de prueba proporcionada. Si ningún elemento cumple la prueba, el método devuelve -1. Es útil cuando necesitas encontrar el índice de un elemento específico en un slice
```go
package main

import (
	"fmt"

	"github.com/Leonardo-Antonio/slicefunc"
)

type Person struct {
	Name     string
	Id       uint32
	LastName string
	FullName string
}

var persons []Person = []Person{
	{Name: "leonardo", Id: 1, LastName: "Tonson"},
	{Name: "Oppo", Id: 2, LastName: "Kicju"},
	{Name: "Tans", Id: 1, LastName: "Tonson"},
	{Name: "Ruy", Id: 4, LastName: "Itachi"},
}

func main() {
	response := slicefunc.FindIndex(persons, func(p Person) bool {
		return p.Id == 2
	})

	fmt.Println(response) // 2
}
```