package main

import "fmt"

type Person struct {
	name string
	age  int
	pet  bool
}

func main() {
	p1 := Person{
		"Arafat",
		23,
		false,
	}

	var p2 Person

	p3 := Person{}

	p4 := Person{
		name: "Arafat",
		age:  23,
	}

	var p5 = map[string][]string{
		"Name": {"Arafat", "Talha", "Alfie"},
		"Age":  {"23", "24", "23"},
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(p4)
	fmt.Println(p5)
	fmt.Println(p5["Age"])
	fmt.Println(p2.age)

}
