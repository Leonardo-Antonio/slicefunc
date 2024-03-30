package slicefunc

// English: Checks if at least one element in the slice satisfies the provided test function. It returns true if it finds an element that satisfies the test, and false if none is found. It's useful when you need to verify whether any element in a slice meets a certain condition.
// Español: Comprueba si al menos un elemento del slice cumple con la función de prueba proporcionada. Retorna true si encuentra un elemento que cumple la prueba, y false si no encuentra ninguno. Es útil cuando necesitas verificar si algún elemento en un slice cumple una cierta condición.
/*
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
*/
func Some[Type any](data []Type, callback func(Type) bool) bool {
	for i := 0; i < len(data); i++ {
		if callback(data[i]) {
			return true
		}
	}

	return false
}
