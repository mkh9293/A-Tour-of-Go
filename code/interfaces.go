package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func interfaces() {
	var a Abser
	f := MyFloat2(-math.Sqrt2)
	v := Vertex4{3, 4}

	a = f
	a = &v

	fmt.Println(a.Abs())
}

type MyFloat2 float64

func (v MyFloat2) Abs() float64 {
	if(v < 0) {
		return float64(-v)
	}

	return float64(v)
}

type Vertex4 struct {
	X, Y float64
}

func (v *Vertex4) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}