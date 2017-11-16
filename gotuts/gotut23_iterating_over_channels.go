package main

import ("fmt"
		"sync"
)

// *****************************************
// Used to syncronize threads
// *****************************************
var wg sync.WaitGroup

// If we need to pass a channel, it needs to be an argument
func foo(c chan int, someValue int) {
	defer wg.Done()
	// Send value over channel, don't return anything
	c <- someValue * 5
}


// // *************************************************************************
// // NOTE: Sending and receiving of data from channels is BLOCKING
// // *************************************************************************
func main() {
	// NOTE: We need a buffer to prevent values from exceeding cap of 10 items! ( This will prevent everything from blocking )
	fooVal := make(chan int, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(fooVal, i)
	}

	// This will ensure all the channels finish submitting data
	wg.Wait()

	// Need to close channel, but we should make sure this only runs when finished
	close(fooVal)

	for item := range fooVal {
		fmt.Println(item)
	}
	// go foo(fooVal, 5)
	// go foo(fooVal, 3)

	// // Since we know we are expecting 2 values, send one to v1, other to v2
	// // arrow never goes the other way, but you can use channel as a source
	// // *************************************************************************
	// // NOTE: Sending and receiving of data from channels is BLOCKING
	// // *************************************************************************
	// v1 := <-fooVal
	// v2 := <-fooVal

	// fmt.Println(v1,v2)

	// // NOTE: Returns
	// //  |2.2.3|   hyenatop in ~/dev/bossjones/boss-golang-labs/gotuts
	// // ± |master ?:1 ✗| → go run gotut22_channels.go
	// // 15 25
}
