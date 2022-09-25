package main

import (
	"fmt"
)

func main() {
	fmt.Println(calculateAvarage(5.6, 8.9, 7.7))

	someStuff := []string{"aeaeae", "gaydasdasd", "idjfdsofk"}

	fmt.Println(someStuff, 123)

	boringFunction("eaeae", 0, 4)

	// function inside a function
	func() {
		fmt.Println("it's my internal function")
	}()

	// when pass a anonymous func to a variable
	// we need to remove the () in the final
	times := func(numberOne, numberTwo int) int {
		return numberOne * numberTwo
	}

	fmt.Println(times(1, 2))

	hundred := 100
	increment := func(number int) int {
		number = number + hundred
		return number
	}

	incrementTwo := func() int {
		hundred = hundred + 100
		return hundred
	}

	fmt.Println("incrementing 100: ", increment(23))
	fmt.Println("incrementing 100: ", increment(25))
	fmt.Println("incrementing 100: ", incrementTwo())
}

func boringFunction(someStuff ...interface{}) {
	// when use ... is like to pack the data
	// when use ... after is to unpack data

	fmt.Println(someStuff...)
}

// ... -> means we can pass many values in parameters
func calculateAvarage(numbers ...float64) float64 {
	total := 0.0

	for _, number := range numbers {
		total = total + number
	}

	return total / float64(len(numbers))
}
