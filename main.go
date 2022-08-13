package main

import (
	"fmt"
	"math/rand"
)

var foo = "bar"

func main() {
	fmt.Println("Here we go...", "\n", foo, rand.Intn(10))
}
