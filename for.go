package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// En Basit For Döngüsü
	for i := 0; i < 10; i++ {
		// fmt.Println(i)
	}
	// For Sonsuz Döngü ve Break...{Random Kütüphanesi}

	for {
		x := rand.Intn(100)
		fmt.Println(x)
		if x == 21 {
			break // Eğer 21 Olursa Döngüyü durdur ve Çık
		}
	}
}
