package main

import (
	"fmt"
)

// Clousure
func say() func() {
	count := 0
	return func() {
		for {
			count++
			if count == 6 {
				break
			}
			fmt.Println(count)
		}
	}

}

func main() {
	saydir := say()
	saydir()

}
