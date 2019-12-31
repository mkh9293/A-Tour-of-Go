package main

import (
	"fmt"
	"math"
)

type Vertex2 struct {
	X, Y float64
}

func (v Vertex2) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func methods() {
	v := Vertex2{3, 4}
	fmt.Println(v.Abs())
}