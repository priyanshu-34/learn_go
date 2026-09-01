package main

import "fmt"

func addTwoNumbers(a int, b int) int {
	return a + b
}

func doubleNumber(num *int) {
	*num = *num * 2
}

func multiReturn()(int, int){
	return 10, 20
}

func variadicFunc(nums ...int) int {
	fmt.Println("Numbers: ", nums)
	total:= 0

	for _, num := range nums {
		total += num
	}
	return total
}

func function(){
	fmt.Println("Hello, Function!")

	// Normal function with parameters and return type
	result := addTwoNumbers(5, 10)
	fmt.Println("Result:", result)

	// Function with pointer parameter
	a:=2
	fmt.Println("Original value:", a)
	doubleNumber(&a)
	fmt.Println("Doubled value:", a)

	// Function with multiple return values
	x, y := multiReturn()
	fmt.Println("Multiple return values:", x, y)
	_, z := multiReturn()
	fmt.Println("Only second return value:", z)

	//Variadic functions
	total := variadicFunc(1, 2, 3, 4, 5)
	fmt.Println("Total:", total)
	total2 := variadicFunc(10, 20)
	fmt.Println("Total2:", total2)

	arr := []int{1, 2, 3, 4, 5}
	total3 := variadicFunc(arr...)
	fmt.Println("Total3:", total3)

	
}
