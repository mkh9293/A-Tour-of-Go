package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	z := complex128(1)
	for i:= 0; i<10; i++{
		z = z - (z*z*z -x) / (3 * z*z)
	}
	return z
}

func exerciseLoopsAndFunctions() {
	fmt.Println(Cbrt(2 + 3i))
	fmt.Println(cmplx.Pow(2.0 + 3i, 1.0/3))
}