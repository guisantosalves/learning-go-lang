package main

import "fmt"

var GlobalNweVariable = "this is my global"

func main() {

	// vars
	var awesomeString string = "Hello world with string"
	var reallyCoolInt int

	reallyCoolInt = 20
	coolFloat := 34.3

	fmt.Println("AwesomeString is:", awesomeString)
	fmt.Println("It's my int", reallyCoolInt)
	fmt.Println("a new variable:", coolFloat)
	fmt.Println("Global variable:", GlobalNweVariable)

	myName := getFirstName()

	fmt.Println("My function:", myName)

	guizaodozap, andre := getABunchOfNames()

	fmt.Println("return two values: " + guizaodozap + " e " + andre)

	// we can ignore any return from a function using _ like:
	//_, andree := getABunchOfNames()

	incrementing := incrementByOne(2)
	fmt.Println("incrementing number 2: ", incrementing)

	addingFloats := addTwoFloats(22.2, 12.1)

	fmt.Println("adding: ", addingFloats)
}

// when we start a function or global variables with capital letter,
//! it means that they are PUBLIC
func SubtractTwoIntegers(numberOne, numberTwo int) int {
	return numberOne - numberTwo
}

// in this way the func set float64 for boths
func addTwoFloats(numberOne, numberTwo float64) float64 {
	return numberOne + numberTwo
}

// set what returns after the ()
func getFirstName() string {
	return "guizaodozap"
}

// returning two things
func getABunchOfNames() (string, string) {
	return "guilherme", "andre"
}

func incrementByOne(number int) int {
	return number + 1
}
