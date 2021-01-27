// Package popcount will calculate the popcount of integers
package popcount

import (
	"fmt"
)

func main() {
	fmt.Println(PopCount(45))
	fmt.Println(PopCountLoop(45))
}
