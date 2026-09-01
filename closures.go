package main
import "fmt"

func testClosure() func() int {
	i:=0

	return func() int {
		i++
		return i
	}
}

func closure(){
	fmt.Println("Closures....")

	closureFunc := testClosure()
	fmt.Println(closureFunc())
	fmt.Println(closureFunc())
	fmt.Println(closureFunc())

}