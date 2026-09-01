package main
import "fmt"

func maps(){
	m := make(map[string]int)
	m["priyanshu"] = 1
	m["rohit"] = 2
	m["saurabh"]=  3

	fmt.Println(m)
	fmt.Println(len(m))

	u, v:=m["Priyanshu"] // we will get 0 when there is nothing found
	fmt.Println(u, v)

	delete(m, "rohit")
	fmt.Println(m)

	clear(m)
	fmt.Println(m)
}