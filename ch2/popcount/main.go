// Package popcount will calculate the popcount of integers
package main

import (
	"fmt"

	"popcount"
)

func main() {
	fmt.Println(popcount.PopCount(45))
	fmt.Println(popcount.PopCountLoop(45))
}
