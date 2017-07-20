package main

import (
	"fmt"
)

// Örnek1
func topla(a int, b int) int {
	return a + b
}

// Örnek 2
func topla2(a, b int) int {
	return a + b
}

// Örnek 3
func islem(a, b int) (int, int, int, int) {
	topla := a + b
	carp := a * b
	cikar := a - b
	bol := a / b
	return topla, carp, cikar, bol

}

//Örnek 4
func str(a, b string) (string, string) {
	return a, b

}
func main() {
	fmt.Println("Örnek1:", topla(2, 4))
	fmt.Println("Örnek2:", topla2(2, 4))
	// İşlem
	topla, carp, cikar, bol := islem(21, 30)
	fmt.Println("Toplama:", topla)
	fmt.Println("Çarpma:", carp)
	fmt.Println("Çıkarma:", cikar)
	fmt.Println("Bölme:", bol)
	// String
	ad, soyad := str("Müjdat", "Cengiz")
	fmt.Println("Örnek4:", ad, soyad)

}
