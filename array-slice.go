package main

import (
	"fmt"
)

func main() {
	// Slice İle Dizi Oluşturma
	arraySlice := make([]int, 2)
	fmt.Println(arraySlice)

	// Sonradan Arraya Eleman Ekleme
	arraySlice = append(arraySlice, 10, 20, 30, 40, 50, 60)
	fmt.Println(arraySlice)

	// Dizi Eleman Sayısını Öğrenme
	fmt.Println(len(arraySlice))

	// Dizi Clonlama COPY
	arraySliceCopy := make([]int, len(arraySlice))
	copy(arraySliceCopy, arraySlice)
	fmt.Println("Orjinal:", arraySlice)
	fmt.Println("Copy:", arraySliceCopy)

	//Belirli sayıda elemanı alma

	arraySliceCut := arraySlice[:5] // İlk Beşini Al...
	fmt.Println("Cut", arraySliceCut)

}
