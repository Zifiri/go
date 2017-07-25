package main

import "errors"
import "fmt"

func checkAirPort(airport string) error {
	if airport == "sabiha-gokcen" {
		return errors.New("Bu Hava Alanı Çok Uzak Trafikte Bunalabilirsin...")
	}
	return nil
}

func main() {
	havalanlari := map[string]int{"sabiha-gokcen": 40, "ataturk": 20, "yeni-havalani": 200}
	for key, val := range havalanlari {
		if a := checkAirPort(key); a != nil {
			fmt.Println(key+":", a)
		} else if val > 100 {
			fmt.Println("Sakın", key, "seçme !! Çünkü", val, "dakika!! ")
		} else {
			fmt.Println(key, "Hava alanı mesafe olarak sadece", val, "dakika fav'a at...")
		}

	}
}
