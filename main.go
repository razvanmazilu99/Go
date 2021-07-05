package main

import (
	"fmt"
	ent "internship/golang/entity"
)

func main() {

	// name := "Atoss"
	// fmt.Println("Hello World!", name)

	// for i := 0; i < 5; i++ {
	// 	fmt.Print(i)
	// }

	// fmt.Println()

	var names = []string{"a", "b", "c", "d", "e"}
	// names2 := make([]string, 5, 5)

	// fmt.Println(names)

	person := ent.Person{Name: "Razvan", Gender: "Male"}
	fmt.Println(person)

	_, age := ReturnTypes()

	fmt.Println(age)
	fmt.Println(person.IsMale())

	for i, name := range names {
		fmt.Println(name, "At position: ", i)
	}
}

func ReturnTypes() (name string, age int) {
	return "Razvan", 21
}
