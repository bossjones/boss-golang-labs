package main

import "fmt"

// ########################################################################
// NOTE: Structs
// In go, there are no classes
// We use structs instead of classes.
// Structs can use methods as well
// ########################################################################

type car struct {
	gas_pedal uint16 // anything from 0 - 65535
	break_pedal uint16
	steering_wheel int16 // -32K - +32k
	top_speed_kmh float64
}


func main(){
	a_car := car{gas_pedal: 22341,
				break_pedal: 0,
				steering_wheel: 12561,
				top_speed_kmh: 225.0}

	// NOTE: This form below is also valid, but less legible
	// a_car := car{22341, 0, 12561, 225.0}

	fmt.Println(a_car.gas_pedal)
}
