package main

import "fmt"

func main() {

	//variables: form1
	var name, lastname string
	name = "Christian"
	lastname = "Bacilio De La Cruz"
	fmt.Println(name)
	fmt.Println(lastname)

	//variables: form2
	var (
		age int
	)
	age = 36

	fmt.Println(age)

	//variables: form3
	var (
		from string = "Huancayo"
	)

	fmt.Println(from)

	//variables: form4
	var (
		occupation = "Ing Software"
	)

	fmt.Println(occupation)

	//variables: form4
	var study, year = "Ricardo Palma University", 5

	fmt.Println(study)
	fmt.Println(year)

	//variables: form5:= dentro de funciones

	mom, dad := "Silvia", "Dario"

	fmt.Println(mom)
	fmt.Println(dad)
}
