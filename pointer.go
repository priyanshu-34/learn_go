package main

import "fmt"

func nonPointerFunc(a int) {
	a = 0;
}

func pointerFunc(a *int) {
	*a = 0;
}

func pointer(){
	fmt.Println("Pointer.....")
	a := 1
	fmt.Println("Initial value of a:", a)
	nonPointerFunc(a)
	fmt.Println("Value of a after nonPointerFunc:", a)
	
	pointerFunc(&a)
	fmt.Println("Value of a after pointerFunc:", a)

	// Way to create a pointer in go:
	b := new(42)
	fmt.Println("Value of b: ", *b);
	nonPointerFunc(*b)
	fmt.Println("Value of b after nonPointerFunc:", *b)
	pointerFunc(b)
	fmt.Println("Value of b after pointerFunc:", *b)


}