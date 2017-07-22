package main

import (
	"fmt"
)

// Structların kendi fuctionlarına methods deniyor...
type hesapla struct {
	genislik, yukseklik int
}

func (h *hesapla) carp() int {
	return h.genislik * h.yukseklik
}
func (h *hesapla) topla() int {
	return h.genislik + h.yukseklik

}

func main() {
	// Sayıyı Gönderelim
	a := hesapla{12, 21}
	fmt.Println("Çarpım Sonucu:", a.carp())
	fmt.Println("Toplama Sonucu:", a.topla())
}
