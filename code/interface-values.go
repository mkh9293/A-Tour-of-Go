package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

func interfaceValues() {
	var i I

	// 1 
	i = &T{"Hello"}
	describe(i)
	i.M()

	// 2
	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// 1
type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

// 2
type F float64

func (f F) M() {
	fmt.Println(f)
}

