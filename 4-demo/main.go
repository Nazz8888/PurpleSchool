package main

import "fmt"

func main() {
	a := 42
	b := "hello"
	c := [3]int{1, 2, 3}
	c[0] = 100

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
