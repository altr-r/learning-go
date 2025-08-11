package learningstructs

import "fmt"

type person struct {
	name string
	age  int
	pet  bool
}

func LearningStructs() {
	var fred person

	fmt.Println(fred) // ouput {"" 0 false} zero values for each field in the struct

	bob := person{}
	fmt.Println(bob) // using literal, output: {"" 0 false}

	// comma separated values style
	julia := person{
		"Julia",
		21,
		true,
	}
	fmt.Println(julia)

	// map literal style
	beth := person{
		pet:  false,
		name: "Beth",
	}
	fmt.Println(beth)
	/*
	* You use the names of the fields in the struct to specify the values.
	* This style has some advantages.
	* It allows you to specify the fields in any order, and you don’t need to provide a value for all fields.
	* Any field not specified is set to its zero value.
	 */

	fmt.Println(beth.name) // field is accessed through dot notation
}
