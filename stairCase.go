package main

import (
	"fmt"
)

func staircase(n int32) {
	for i := int(n) - 1; i >= 0; i-- {
        for j := 0; j < int(n); j++ {
            if j < i {
                fmt.Print(" ")
            } else {
                fmt.Print("#")
            }
        }
        fmt.Println()
    }
}

func main() {
	staircase(6)
}
