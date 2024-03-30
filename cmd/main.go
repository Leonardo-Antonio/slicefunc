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
