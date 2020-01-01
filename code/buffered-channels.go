package main

import "fmt"

func befferedChannels() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	
	fmt.Println(<-ch)
	fmt.Println(<-ch)	
}