package main

import "fmt"

func fibonacci() func() int {
	cur, prev := 0, -1

	return func() int {
		val := cur + prev
		cur, prev = val, cur

		if val < 1 {
			cur++
			return cur
		}

		return val
	}
}

func exerciseFibonacciClosure() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
