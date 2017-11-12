package main

import "fmt"

func main() {
	// := only when initializing var
	// NOTE: Typical counter loop
	// for i:=0; i<10; i++{
	// 	fmt.Println(i)
	// }

	// refactor 2
	// i:=0
	// for i<10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// refactor 3 ( infinite loop, similar to while loop)
	// x := 5
	// for {
	// 	fmt.Println("Do stuff", x)
	// 	x+=3
	// 	if x > 25{
	// 		break
	// 	}
	// }

	// refactor 4 ( infinite loop, similar to while loop)
	// x := 5
	// for x:=5; x<25; x+=3{
	// 	fmt.Println("Do stuff", x)
	// }

	// refactor 5
	a:=3
	for x:=5; a<25; x+=3{
		fmt.Println("Do stuff", x)
		a+=4
	}
}
