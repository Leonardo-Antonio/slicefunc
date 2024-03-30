package slicefunc

// English: Takes an slice, applies a function to each element, and returns a new slice with the results. It’s useful for transforming data without modifying the original slice.
// Español: Toma un slice, aplica una función a cada elemento, y devuelve un nuevo slice con los resultados. Es útil para transformar los datos sin modificar el slice original.
/*
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
*/
func Map[Type any](data []Type, callback func(Type) Type) []Type {
	result := make([]Type, len(data))
	for i := 0; i < len(data); i++ {
		result[i] = callback(data[i])
	}

	return result
}
