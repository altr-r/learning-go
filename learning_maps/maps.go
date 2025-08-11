package learningmaps

import "fmt"

func LearningMaps() {
	/*
	* maps
	* maps can be declared with var map[type]type
	* the type inside the square brackets is the key and the type outside the square brackets is the value
	 */

	var firstMap map[int]string
	fmt.Println(firstMap) // ouput: map[]

	secondMap := map[int]string{} // output: map[], can also be declared in this notation using map literal
	fmt.Println(secondMap)

	/*
	* Even though the output of firstMap and secondMap seem similar, they are not.
	* It will cause go to panick if it is attempted to write values from the firstMap as that is a nil map.
	* On the other hand, it is possible to read and write from a map assigned with empty map literal
	 */

	/* a map assigned with non empty map literal
	* below is a non empty map literal which has string as the key and a slice of strings as the value
	 */

	nonEmptyMap := map[string][]string{
		"Orcas":   {"Fred", "Ralph", "Bijou"},
		"Lions":   {"Sarah", "Peter", "Billie"},
		"Kittens": {"Waldo", "Raul", "Ze"},
	}

	/*
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	*/
	// in the book the above commented way is how the map was declared, but it says redundant in the latest compiler
	fmt.Println(nonEmptyMap) // output: map[Kittens:[Waldo Raul Ze] Lions:[Sarah Peter Billie] Orcas:[Fred Ralph Bijou]]
}
