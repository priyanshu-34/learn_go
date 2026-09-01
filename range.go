package main
import "fmt"

func _range(){
	fmt.Println("Range....")

	arr := []int{1, 2, 3, 4, 5}
	for index, val := range arr {
		fmt.Println("Index:", index, "Value:", val)
	}

	fmt.Println("................................")

	brr := []string{"a", "b", "c", "d"}
	for index, val := range brr {
		fmt.Println("Index:", index, "Value:", val)
	}
	fmt.Println("................................")


	str := "Hello"
	for index, val := range str {
		fmt.Println("Index:", index, "Value:", string(val))
	}
	fmt.Println("................................")


	mp := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, val := range mp {
		fmt.Println("Key:", key, "Value:", val)
	}

}