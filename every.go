package slicefunc

// English: Checks if all elements in a slice satisfy a provided test function. It returns true if all elements pass the test, and false if at least one does not. It's useful when you need to verify whether all elements in a slice meet a certain condition.
// Español: Verifica si todos los elementos de un slice cumplen con una función de prueba proporcionada. Retorna true si todos los elementos pasan la prueba, y false si al menos uno no la cumple. Es útil cuando necesitas verificar si todos los elementos en un slice cumplen una cierta condición
/*
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
*/
func Every[Type any](data []Type, callback func(Type) bool) bool {
	for i := 0; i < len(data); i++ {
		if !callback(data[i]) {
			return false
		}
	}

	return true
}
