package main
import "fmt"

func stringAndRunes(){
	fmt.Println("Strings and Runes....")

	a := "Hello, World"
	fmt.Println("Length of the string: ", len(a))
	fmt.Println("First character of the string: ", a[0]) // This will print the ASCII value of the first character
	fmt.Println("First character of the string as string: ", string(a[0])) // This will print the first character as a string

	// runes are equivalent to characters in other languages. They are used to represent Unicode code points.
	r := 'H' // rune literal
	fmt.Println("Rune value: ", r)
	fmt.Println("Rune as string: ", string(r))
	fmt.Println("Rune as int: ", int(r)) // This will print the Unicode code point of the rune

	// You can also convert a string to a slice of runes
	runes := []rune(a)
	fmt.Println("Runes in the string: ", runes)
	fmt.Println("First rune in the string: ", runes[0])
	fmt.Println("First rune in the string as string: ", string(runes[0]))
}