package main
import ("fmt")

func array() {
	// Array is a collection of elements of the same type
	// Syntax: var array_name = [size]type
	// size is the number of elements in the array
	// size is fixed and cannot be changed
	// size is optional you can use ... for size not fixed like var arr = [...]int{1, 2, 3, 4, 5}
	// like ArrayList in Java but ArrayList can grow and shrink dynamically but Array cannot grow and shrink dynamically
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	var arr2 = [...]string{"Hello", "World", "Go"}
	fmt.Println(arr2)

	// you can use :=
	// syntax: array_name := [size]type{value1, value2, value3, ...}
	arr3 := [3]float64{3.14, 2.718, 1.618}
	fmt.Println(arr3)
	arr4 := [...]bool{true, false, true}
	fmt.Println(arr4)

	// access element in array
	fmt.Println(arr[0]) // 1
	fmt.Println(arr2[1])
	// use arr[index] for access element in array and index start from 0
	// if you try to access index that is out of range it will error
	// can change element in array
	arr[0] = 10
	fmt.Println(arr) // [10 2 3 4 5]

	// array initialize with default value
	var arr5 [5]int
	fmt.Println(arr5) // [0 0 0 0 0]
	var arr6 = [5]int{1,2}
	fmt.Println(arr6) // [1 2 0 0 0]
	// can specify index
	var arr7 = [5]int{0: 1, 2: 3, 4: 5}
	fmt.Println(arr7) // [1 0 3 0 5]

	// length of array
	// use len(arr) for find it
	fmt.Println(len(arr)) // 5
}