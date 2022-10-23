package main

import (
	"fmt"
)

var (
	hey bool = false
)

func main() {
	i := 42
	f := float64(i)
	u := uint(f)
	never := true
	const rick = "yes"

	fmt.Println(u, never, rick, hey)
	helloWorld()
}

func helloWorld() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)
}
