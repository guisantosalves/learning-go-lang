package shapes

import "math"

func (c *Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}
