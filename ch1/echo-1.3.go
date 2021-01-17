// Echos arguments sent to program
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	nano := time.Since(start).Nanoseconds()
	fmt.Printf("Loop Duration: %dns\n", nano)

	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	nano = time.Since(start).Nanoseconds()
	fmt.Printf("Join Duration: %dns\n", nano)
}
