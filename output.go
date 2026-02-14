package main
import ("fmt")

func output() {
	// Output in golang is done using fmt package
	// Print is used to print without new line
	var hello string = "Hello"
	var world string = "World"
	fmt.Print(hello) // Hello
	fmt.Print(world) // World
	// it should be in one line because we use Print

	// Println is used to print with new line
	fmt.Println(hello) // Hello
	fmt.Println(world) // World
	// it should be in two line because we use Println

	// Printf is used to print with format
	// %v is used to print value
	// %T is used to print type
	fmt.Printf("Value: %v, Type: %T\n", hello, hello) // Value: Hello, Type: string

	// Format
	// %d is used to print integer
	// %f is used to print float
	// %s is used to print string
	// %t is used to print boolean
	// it like printf in C++ or Java
}