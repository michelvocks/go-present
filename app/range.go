package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v1 := Vertex{1, 2}
	v2 := Vertex{X: 3, Y: 4}

	var slice []Vertex
	slice = append(slice, v1)
	slice = append(slice, v2)

	for i, v := range slice {
		fmt.Printf("id %d = %+v\n", i, v)
	}

}
