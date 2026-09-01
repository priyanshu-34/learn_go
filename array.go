package main

import "fmt"

func array(){
	var arr [5]int;  // Here each element of the array will be initialized with 0
	fmt.Println(arr)
	fmt.Println(len(arr))

	arr[4] = 100;
	fmt.Println(arr)

	b:= [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	b = [...]int{11, 12, 13, 14, 15}
	fmt.Println(b)

	b = [...]int{100, 3:22, 45}
	fmt.Println(b)
}