package slicefunc

// English: Create a new slice with all elements that pass a test implemented by the provided function. It's useful for selecting elements from an array that meet certain conditions.
// Español: Crea un nuevo slice con todos los elementos que pasan una prueba implementada por la función proporcionada. Es útil para seleccionar elementos de un arreglo que cumplen ciertas condiciones.
/*
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
*/
func Filter[Type any](data []Type, callback func(Type) bool) []Type {
	var result []Type
	for i := 0; i < len(data); i++ {
		if callback(data[i]) {
			result = append(result, data[i])
		}
	}

	return result
}
