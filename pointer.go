package main

import (
	"fmt"
)

func main() {
	degisken := 30
	degisken1 := &degisken

	fmt.Println("Normal Değişken", degisken)
	fmt.Println("Pointeri Alınmış Değişken", degisken1)
	fmt.Println("Pointer ", *degisken1)
}
