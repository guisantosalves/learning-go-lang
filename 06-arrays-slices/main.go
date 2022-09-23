package main

import (
	"fmt"
	"sort"
)

// array has a size limit

func main() {
	var sweets = [4]string{
		"milk",
		"dark chocolate",
		"marshmellow",
		"toffe apple",
	}

	fmt.Println("sweet at 0:", sweets[0])
	sweets[1] = "caramel"
	fmt.Println("sweet at 1:", sweets[1])
	fmt.Println("length of an array:", len(sweets))

	// slice
	var scores []int

	// it works like ...scores, new element or slice
	scores = append(scores, 20)
	scores = append(scores, 30)
	scores = append(scores, 2)
	scores = append(scores, 5)

	sort.Ints(scores)

	for i := 0; i < len(scores); i++ {
		fmt.Println(scores[i])
	}

	// for index, score := range scores {
	// 	fmt.Println(index, " and ", score)
	// }

	for _, score := range scores {
		fmt.Println("my elements from the slice: ", score)
	}

	fmt.Println("Print slice:", scores)

	newScores := []int{2, 6, 9, 8}
	newScores = append(newScores, 8)

	fmt.Println("new slice:", newScores)

	fmt.Println("my func: ", avarageScore(newScores))

}

func avarageScore(scores []int) float64 {
	total := 0

	for _, score := range scores {
		total = total + score
	}

	return float64(total / len(scores))
}
