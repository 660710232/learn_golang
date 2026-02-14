package main
import ("fmt")

func constant() {
	// Constant is a value that cannot be changed
	// Syntax: const constant_name type = value
	// constant name should be in uppercase
	const PI float64 = 3.14
	fmt.Println(PI)
	// if you try to change the value of constant it will error

	// you can also declare constant without type
	const A = "Hello"
	fmt.Println(A)

	// you can declare multiple constant in block
	const (
		B = "World"
		C = 10
		D = 3.14
	)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
}