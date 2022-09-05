package main

import "fmt"

func main() {

	/*
		if else
	*/
	var number int

	number = 15

	// even -> par
	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else if number%3 == 0 {
		fmt.Println(number, " is multiple of 3")
	} else {
		fmt.Println(number, "is odd")
	}

	/*
		switch case
	*/
	var favouriteFruits string
	favouriteFruits = "aa"

	switch favouriteFruits {
	case "apple":
		fmt.Println("yum apples rsrs")
	case "pear":
		fmt.Println("pear here is good")
	case "banana":
		fmt.Println("banana is cool")
	default:
		fmt.Println("I dont like this fruit")
	}

	/*
		for
	*/
	i := 0
	for i < 10 {
		fmt.Println(i, "doubled is: ", i*2)
		i = i + 1
	}

	for x := 0; x < 10; x++ {
		if x%2 == 0 {
			fmt.Println(x, "is even")
		} else {
			fmt.Println(x, "is odd")
		}
	}

}
