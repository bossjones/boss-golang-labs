package main

import "fmt"

func main() {
	x := 15
	a := &x // memory address
	fmt.Println(a) // print: memory address
	fmt.Println(*a) // print: actual value at memory address of a
	*a = 5 // Overwrote refrence to x (from 5 -> 15)
	fmt.Println(x)
	*a = *a**a
	fmt.Println(x)
	fmt.Println(*a)
}
