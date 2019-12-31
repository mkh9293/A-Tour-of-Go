package main

import (
	"fmt"
	"math"
)

type Vertex3 struct {
	X, Y float64
}

func (v Vertex3) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex3) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func methodsPointers() {
	v := Vertex3{3, 4}
	v.Scale(10)
	fmt.Println(v, v.Abs())
}
