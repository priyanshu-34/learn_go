package main

import (
	"fmt"
)

func slices(){
	var s []string // Use this to create slices when you don't have the fix size
	fmt.Println(s)
	fmt.Println(s == nil, len(s) == 0)

	s = append(s, "4")
	fmt.Println(s)

	t := make([]string, 4) // Use this to create slices when you have to create fix size of slice
	t[0] = "a"
	t[1] = "b"
	t[2]= "c"
	t[3] = "d"
	t[4] = "e"

	fmt.Println(t)

}