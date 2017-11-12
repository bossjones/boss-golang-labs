package main

import "fmt"

const usixteenbixmax float64 = 65535
const kmh_multiple float64 = 1.60934 // 1mile per hour = 1.60934 Km/h

// METHODS
// #######################################################################
// Two types: Value receivers + Pointer receivers
// Value Receiver: Calculations on values
// Pointer receivers: If you want to actually modify a value in a struct
// #######################################################################

type car struct {
	gas_pedal uint16 // anything from 0 - 65535
	break_pedal uint16
	steering_wheel int16 // -32K - +32k
	top_speed_kmh float64
}

// Value Receiver Method
func (c *car) kmh() float64 {
	// NOTE: If you don't make this into a pointer receiver, then setting value below will not persist other calls!
	c.top_speed_kmh = 500
	// Returns top speed of car in km/h
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbixmax)
}

// Value Receiver Method
func (c car) mph() float64 {
	// Returns top speed of car in mi/h
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbixmax/kmh_multiple)
}

// Reading through pointer to get value, then modify it
func (c *car) new_top_speed(newspeed float64){
	c.top_speed_kmh = newspeed
}

// Function that does same thing as method, requires writing over entire car object ( NOT EFFICENT, but YOU CAN DO THIS!)
func newer_top_speed(c car, speed float64) car {
	c.top_speed_kmh = speed
	return c
}

// ####################################################################################################
// # When to use value receivers?? vs pointer receivers??
// ####################################################################################################
// What's the point of a value receiver ? It makes a copy, instead of modifying the exact value.
// If the struct is small, value receiver more effecient ... less to copy, less garbage
// If struct is big, too much to copy, lots of overhead, inefficent
// ####################################################################################################

func main() {
	a_car := car{gas_pedal: 65000,
				break_pedal: 0,
				steering_wheel: 12561,
				top_speed_kmh: 225.0}

	// NOTE: This form below is also valid, but less legible
	// a_car := car{22341, 0, 12561, 225.0}

	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
	// a_car.new_top_speed(500)
	a_car = newer_top_speed(a_car, 500)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
}
