package main

import (
	"fmt"
)

type gorevli struct {
	ad  string
	yas int
}

func main() {
	// Kullanımı
	fmt.Println(gorevli{"Mahmut Tuncer", 50})
	fmt.Println(gorevli{ad: "Tarık Akan", yas: 6})
	// Değişkene Atama
	gr := gorevli{"Hakan Taşıyan", 60}
	fmt.Println("Görevli Yaşı:", gr.yas)
	// Pointer ile Kullanım
	gr2 := &gr
	gr2.ad = "Müslüm Gürses"
	fmt.Println(gr2.ad)
}
