package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func SwapGlobal(x, y string) (string, string) {
	return y, x
}

func main() {
	// Private call
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// Public call
	c, d := SwapGlobal("world", "hello")
	fmt.Println(c, d)
}
