package main

import "fmt"

func main() {
	// key is string, value is float32. Think dictonary in python
	// Refactor 1
	// var grades map[string]float32

	// Refactor 2
	// NOTE: A map usually doesn't have any values. To initialize one, we use golang built-in make command
	Grades := make(map[string]float32)

	Grades["Timmy"] = 42
	Grades["Jess"] = 92
	Grades["Sam"] = 67
	fmt.Println(Grades)

	TimsGrade := Grades["Timmy"]
	fmt.Println(TimsGrade)

	delete(Grades, "Timmy")
	fmt.Println(Grades)

	// Iterate through map
	for k, v := range Grades {
		fmt.Println(k,":",v)
	}

}
