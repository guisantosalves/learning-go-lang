package main

import "fmt"

func main() {
	var myFavoriteFruit string

	myFavoriteFruit = "apple"

	fmt.Println("my favorite fruit is:", myFavoriteFruit)

	myFavoriteFruit = "pear"

	fmt.Println("my favorite fruit is:", myFavoriteFruit)

	// with & we send the pointer instead the value
	changeFavouriteFruit(&myFavoriteFruit, "strawberry")

	fmt.Println("my favorite fruit is:", myFavoriteFruit)

	changeFavouriteFruit(&myFavoriteFruit, "bluebarry")

	fmt.Println("my favorite fruit is:", myFavoriteFruit)

}

// to access the value which is inside the address of pointer we need to use *myVariable
func changeFavouriteFruit(fruitPoiter *string, newFavourite string) {
	// fmt.Println("printing my pointer:", fruitPoiter, "the value is:", *fruitPoiter)
	*fruitPoiter = newFavourite
}
