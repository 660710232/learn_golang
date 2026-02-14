package main
import ("fmt")

// Go program like C++
// function or method use func funcname()
// main func is entry point
func main() {

	// Decaring variables
	// Syntax: var variable_name type = value
	// can use everywhere in your code space
	// like global variable and local variable
	var number int = 10
	fmt.Println(number)

	// go have another way to decare variable without type
	var age = 21
	fmt.Println(age)

	// Short variable declaration
	// Syntax: variable_name := value
	// can only use inside funtion if use outside function it error
	name := "Ten"
	fmt.Println(name)

	// if you not assign value to variable, it will have zero value
	var score int
	var score2 float64
	var grade string
	var isPassed bool
	fmt.Println(score) // 0
	fmt.Println(score2) // 0
	fmt.Println(grade) // "" (empty)
	fmt.Println(isPassed) // false
}