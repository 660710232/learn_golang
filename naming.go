package main
import ("fmt")

func naming() {
	// Naming in golang is important
	// can use multi variable name
	// such as camel case, pascal case, snake case

	// camel case
	// first letter is lower case and each word is capitalized
	var firsName string = "John"

	// pascal case
	// first letter is upper case and each word is capitalized (like camel case but first letter is upper case)
	var FirstName string = "John"

	// snake case
	// all letter is lower case and each word connected with underscore "_"
	var first_name string = "John"
	
	fmt.Println(firsName)
	fmt.Println(FirstName)
	fmt.Println(first_name)

	// you can use it all but it should be consistent in your code
}