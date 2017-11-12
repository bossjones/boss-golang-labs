package main

// For imports, you need to import the specific packages
// Use the / notation to import them
// Eg. math/rand
import ("fmt"
		"math/rand")

func main() {
	// Capital letter means go will export the function
	fmt.Println("Welcome to go")
	// fmt.Println("The square root of 4 is",math.Sqrt(4))
	fmt.Println("A number from 1-100", rand.Intn(100))
}
