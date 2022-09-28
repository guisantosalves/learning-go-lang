package shapes

import "math"

/*
	To use the method in another file, we need to set the first letter upper
*/

// method of Shape interface
func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) IsSquare() bool {
	return r.Width == r.Height
}
