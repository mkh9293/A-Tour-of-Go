package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex {
	"Bell Labs" : {40.6, -74.3},
	"Google"	: {37.42, -122.08},
}

func main() {
	fmt.Println(m)
}