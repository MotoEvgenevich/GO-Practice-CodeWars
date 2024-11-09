package kata

import (
	"math"
)

func StackHeight2d(layers int) float64 {
	if layers == 0 {
		return 0.0
	}
	height := 1.0 + float64(layers-1)*(math.Sqrt(3)/2)
	return math.Round(height*1000) / 1000

}

/*
Background

I have stacked some pool balls in a triangle.

Like this,

pool balls
Kata Task

Given the number of layers of my stack, what is the total height?

Return the height as multiple of the ball diameter.

Example

The image above shows a stack of 5 layers.

Notes

layers >= 0
approximate answers (within 0.001) are good enough
*/
