package main

// NOTE: This tutorial is all about typing

// For imports, you need to import the specific packages
// Use the / notation to import them
// Eg. math/rand
import ("fmt")

// refactor 1
// func add(x float64,y float64) float64 {
// 	return x+y
// }

// refactor 2 ( get rid of type repetition )
func add(x,y float32) float32 {
	return x+y
}

// Need to return specific return type for each return value
func multiple(a,b string) (string,string) {
	return a,b
}

func main() {
	// var num1 float64 = 5.6
	// var num2 float64 = 9.5

	// refactor 2 ( get rid of type repetition )
	// var num1,num2 float64 = 5.6, 9.5

	// refactor 3 ( Inside of a main function, you don't need to specify type )
	// go can figure it out at compile time

	// NOTE: You will hit an error if you have variables that are defined, but not used
	// num1,num2 := 5.6, 9.5

	w1, w2 := "Hey", "there"

	fmt.Println(multiple(w1,w2))

}
