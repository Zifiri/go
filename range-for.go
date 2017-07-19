package main

import (
	"fmt"
)

func main() {

	diller := []string{"PHP", "C+", "JAVA", "HTML", "CSS", "PYTHON"}
	//Range FOREACH mantığı
	for key, val := range diller {
		fmt.Println("key:", key, "value:", val)
	}

	// Başka bir örnek
	sepetim := map[string]int{"Patates": 20, "Domates": 30, "Kavun": 12, "Kestane": 5, "Üzüm": 27}
	toplam := 0
	for urun, fiyat := range sepetim {
		fmt.Println("Ürün:", urun, "Fiyatı:", fiyat)
		toplam += fiyat
	}
	fmt.Println("Toplam Tutar:", toplam)
}
