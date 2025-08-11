package learningmaps

import "fmt"

//MAPS
func LearningMaps() {
	/*
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

	fmt.Println(nonEmptyMap["Orcas"])                            // reading from map
	nonEmptyMap["Lions"] = []string{"Messi", "Neymar", "Suarez"} // writing in a map

	// The comma ok Idiom
	v, ok := nonEmptyMap["Orcas"]
	fmt.Println(v, ok)
	/*
	* v = values of the key Orcas, if the value exists ok will be true, else it will be false.
	* output: [Fred, Ralph, Bijou] true
	 */

	// deleting from map
	delete(nonEmptyMap, "Orcas") // removes a specified key from the map
	fmt.Println(nonEmptyMap)     // output: map[Kittens:[Waldo Raul Ze] Lions:[Messi Neymar Suarez]], everything under key "Orcas" deleted
	clear(nonEmptyMap)           //clears every key and value from the map
	fmt.Println(nonEmptyMap)     // output: map[]

}
