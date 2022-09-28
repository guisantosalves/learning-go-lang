package main

import (
	"fmt"
	"time"
)

// defer -> runs the command in the end of the function
func giveNumber(intChannel chan int) {
	defer close(intChannel)
	for i := 0; i <= 10; i++ {
		intChannel <- i
	}
}

func doubleNumber(intChannel chan int) {
	// verify if the channel is open or not
	// for {
	// 	number, isOpen := <-intChannel
	// 	if !isOpen {
	// 		break
	// 	}
	// 	fmt.Println(number, " doubled is:", number*2)
	// 	time.Sleep(time.Millisecond * 500)
	// }

	// in this case the range get the quantity of the position that the channel has,
	// not the quantity of the number
	for number := range intChannel {
		fmt.Println(number, " doubled is:", number*2)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	intChannel := make(chan int)
	go giveNumber(intChannel)
	doubleNumber(intChannel)
}
