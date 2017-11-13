package main

import "fmt"

// If we need to pass a channel, it needs to be an argument
func foo(c chan int, someValue int) {
	// Send value over channel, don't return anything
	c <- someValue * 5
}

func main() {
	fooVal := make(chan int)

	go foo(fooVal, 5)
	go foo(fooVal, 3)

	// Since we know we are expecting 2 values, send one to v1, other to v2
	// arrow never goes the other way, but you can use channel as a source
	v1 := <-fooVal
	v2 := <-fooVal

	fmt.Println(v1,v2)

	// NOTE: Returns
	//  |2.2.3|   hyenatop in ~/dev/bossjones/boss-golang-labs/gotuts
	// ± |master ?:1 ✗| → go run gotut22_channels.go
	// 15 25
}
