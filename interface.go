package main

import (
	"fmt"
)

// Rehber Adında Bir Struct oluşturalım..
type Rehber struct {
	id          int
	name        string
	phoneNumber int
}

// İnterface ile rehber foksionlarını belirtelim..
type RehberFunc interface {
	save()
	delete()
	update()
}

// Save Methodu
func (R *Rehber) save() string {
	return R.name + " Başarıyla Eklendi.."
}

// Delete Methodu
func (R *Rehber) delete() string {
	return R.name + " Adlı kişiyi sildiniz.."
}

//Update Methodu
func (R *Rehber) update() string {
	return R.name + "Adlı Kişi bilgileri güncenlendi.."

}

func main() {

	r1 := Rehber{1, "Murat", 535987654}
	// Ekle
	fmt.Println(r1.save())
	// Güncelle
	fmt.Println(r1.update())
	// Sil
	fmt.Println(r1.delete())

}
