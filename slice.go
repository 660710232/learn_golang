package main
import ("fmt")

func slice() {
	// Slice is a dynamic array
	// Syntax: slice_name := []datatype{value1, value2, ...}
	// common way use myslice := []datatype{}

	slice := []int{}
	// you can use len(slice) for find length of slice
	// and cap(slice) for find capacity of slice
	fmt.Println(len(slice)) // 0
	fmt.Println(cap(slice)) // 0

	// you can append element in slice
	slice = append(slice, 1, 2, 3, 4)
	fmt.Println(slice) // [1 2 3 4]
	fmt.Println(len(slice))
	fmt.Println(cap(slice)) // 4

	// create slice with array
	arr1 := [6]int{1, 2, 3, 4, 5, 6}
	myslice := arr1[1:4] // slice from index 2 to index 3 (4 is exclusive)
	fmt.Println(myslice) // [2 3 4]
	fmt.Println(len(myslice)) // 3 (length is from index 1 to index 3)
	fmt.Println(cap(myslice)) // 5 (capacity is from index 1 to end of array)

	// create slice with make function
	// Syntax: slice := make([]datatype, length, capacity)
	slice3 := make([]int, 5, 10)
	fmt.Println(slice3) // [0 0 0 0 0]
	fmt.Println(len(slice3)) // 5
	fmt.Println(cap(slice3)) // 10
	// if capacity is not defined it will same lenght

	// slice and accessed value with index like array and can change value
	slice3[0] = 1
	fmt.Println(slice3) // [1 0 0 0 0]
	// slice can append() like put value in slice
	slice3 = append(slice3, 2, 3, 4, 5)
	fmt.Println(slice3) // [1 0 0 0 0 2 3 4 5]

	// can append slice to another slice
	slice4 := []int{6, 7, 8, 9, 10}
	slice3 = append(slice3, slice4...)
	fmt.Println(slice3) // [1 0 0 0 0 2 3 4 5 6 7 8 9 10]

	// capacity of slice is automatically increased when you append element in slice
	// like lenght is 2, capacity is 4, when you append element in slice it will double the capacity

	// function copy() is used to copy slice
	// Syntax: copy(dest, src)
	slice5 := make([]int, len(slice3))
	copy(slice5, slice3)
	fmt.Println(slice5) // [1 0 0 0 0 2 3 4 5 6 7 8 9 10]
}