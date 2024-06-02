package main

import (
	"fmt"

	"example.com/go_learn/mascot"

	"rsc.io/quote"
)

func main() {
	fmt.Println(mascot.BestMascot())
	fmt.Println(quote.Go())
	fmt.Println(pascal_triangle(1))
	fmt.Println(pascal_triangle(10))
}