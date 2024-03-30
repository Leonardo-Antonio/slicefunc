package slicefunc

// English: Reverses the order of elements in a slice in place and returns a reference to the same slice. The first element becomes the last and the last becomes the first. It's useful when you need to reverse the order of elements in a slice.
// Español: Invierte el orden de los elementos de un slice en su lugar y devuelve una referencia al mismo slice. El primer elemento pasa a ser el último y el último pasa a ser el primero. Es útil cuando necesitas invertir el orden de los elementos de un slice.
/*
var persons []Person = []Person{
	{Nick: "example", Name: "leonardo", Id: 1, LastName: "Tonson"},
	{Nick: "example", Name: "Oppo", Id: 2, LastName: "Kicju"},
	{Nick: "example", Name: "Tans", Id: 1, LastName: "Tonson"},
	{Nick: "example", Name: "Ruy", Id: 4, LastName: "Itachi"},
}

func main() {
	slicefunc.Reverse(persons)
	fmt.Println(persons) // [{Ruy 4 Itachi  av} {Tans 1 Tonson  av} {Oppo 2 Kicju  av} {leonardo 1 Tonson  av}]
}
*/
func Reverse[Type any](data []Type) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}

}
