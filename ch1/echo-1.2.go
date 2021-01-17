// Echos arguments sent to program
package main

import (
	"fmt"
	"os"
)

func main() {
	s := ""
	for i, arg := range os.Args[1:] {
		s = fmt.Sprintf("%d: %s", i, arg)
		fmt.Println(s)
	}
}
