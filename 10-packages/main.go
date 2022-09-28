package main

import (
	"fmt"

	"github.com/guisantosalves/improvinggo/randomstuff"
	"github.com/guisantosalves/improvinggo/shapes"
)

func main() {
	Circle := shapes.Circle{Radius: 4.5}
	Rec := shapes.Rectangle{Width: 5.3, Height: 3.5}
	fmt.Println("Area of circle:", Circle.Area())
	fmt.Println("Area of rectangle:", Rec.Area())
	fmt.Println("is the rectangle square:", Rec.IsSquare())
	randomstuff.PrintSomeRubbish()
}
