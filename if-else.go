package main

import (
	"fmt"
)

func main() {

	// İf Else Mantığı
	sayi := 21
	if sayi == 20 {
		fmt.Println("Doğru Sayı! Sayım", sayi)
	} else if sayi == 21 {
		fmt.Println("Bu Sayıyıda Seviyorum, Sayım ", sayi)
	} else {
		fmt.Println("Oppss.. Sayı hatalı Sayım", sayi)
	}

}
