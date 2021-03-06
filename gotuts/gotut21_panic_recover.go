package main

import (
	"fmt"
	"sync"
	"time"
)

// wg = wait group
var wg sync.WaitGroup

func cleanup() {
	// If theres no panic situation, this returns nil
	if r := recover(); r != nil {
		// If panic, and recover is not nil, do below
		fmt.Println("Recovered in cleanup:", r)
	}
	wg.Done()
}

func say(s string) {
	// Use Defer to make sure it ALWAYS runs
	// Notify wait group that we are done
	defer cleanup()
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
		if i == 2 {
			panic("Oh dear, a 2")
		}
	}
}

func foo() {
	// First in, last out order
	// Meaning ... we'll see Are we done? then Done! in that order.
	defer fmt.Println("Done!")
	defer fmt.Println("Are we done?")
	fmt.Println("Doing some stuff, who knows what?")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	// Everytime before you use a goroutine, add 1 to the waitgroup
	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("There")
	wg.Wait()
	// foo()
}
