package slicefunc

// English: Takes an slice, applies a function to each element, and returns a new slice with the results. It’s useful for transforming data without modifying the original slice.
// Español: Devuelve el valor del primer elemento del slice que cumple con la función de prueba proporcionada. Si ningún elemento cumple la prueba, el método devuelve nil. Es útil cuando necesitas encontrar un elemento específico en un slice
/*
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
*/
func Find[Type any](data []Type, callback func(Type) bool) (result Type, found bool) {
	for i := 0; i < len(data); i++ {
		if callback(data[i]) {
			result = data[i]
			return result, true
		}
	}

	return result, false
}

// English: Returns the index of the first element in a slice that satisfies the provided test function. If no element satisfies the test, the method returns -1. It's useful when you need to find the index of a specific element in a slice.
// Español: Devuelve el índice del primer elemento en un slice que cumple con la función de prueba proporcionada. Si ningún elemento cumple la prueba, el método devuelve -1. Es útil cuando necesitas encontrar el índice de un elemento específico en un slice
/*
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
*/
func FindIndex[Type any](data []Type, callback func(Type) bool) int {
	for i := 0; i < len(data); i++ {
		if callback(data[i]) {
			return i
		}
	}

	return -1
}
