package main

import "fmt"

func fibonacci3(c, quit chan int) {
	x, y := 0, 1

	for {
		select {
		case c <- x :
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit!")
			return
		}
	}
}

func selects() {
	c := make(chan int)
	quit := make(chan int)
	
	go func() {
		for i:=0; i<10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci3(c, quit)
}