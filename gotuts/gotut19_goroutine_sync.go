package main

import ("time"
		"fmt"
		"sync"
)

// wg = wait group
var wg sync.WaitGroup

func say(s string) {
	for i:=0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond*100)
	}
	// Notify wait group that we are done
	wg.Done()
}

func main() {
	// // Goroutine is just a lightweight thread ( There's nothing that says this has to finish. Non blocking )
	// go say("Hey")
	// say("There")

	// // refactor 2 ( returns blank )
	// go say("Hey")
	// go say("There")

	// refactor 3 ( returns blank )

	// Everytime before you use a goroutine, add 1 to the waitgroup
	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("There")
	wg.Wait()

}
