package main

import ("time"
		"fmt"
)

func say(s string) {
	for i:=0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond*100)
	}
}

func main() {
	// // Goroutine is just a lightweight thread ( There's nothing that says this has to finish. Non blocking )
	// go say("Hey")
	// say("There")

	// // refactor 2 ( returns blank )
	// go say("Hey")
	// go say("There")

	// refactor 3 ( returns blank )
	go say("Hey")
	go say("There")
	time.Sleep(time.Second)
}
