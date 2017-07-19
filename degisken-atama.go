package main

import (
	"fmt"
	"reflect"
)

// Değişken Tanımlama
func main() {
	var degisken string = "Degisken" // 1.yol
	fmt.Println(degisken)
	var degisken1 = "Degişken1" // 2.yol
	fmt.Println(degisken1)
	degisken2 := "Değişken2" //3.yol
	fmt.Println(degisken2)
	// #Çoklu Atama
	degisken3, degisken4 := "Değişken3", "Değişken4"
	fmt.Println(degisken3)
	fmt.Println(degisken4)
	// #Reflect ile Değişken Türünü Sorgulama
	deg := true           //bool
	deg1 := false         // bool
	deg2 := 200           // int
	deg3 := 1.333         // float64
	deg4 := "Bu Bir Yazı" // String
	fmt.Println(reflect.TypeOf(deg))
	fmt.Println(reflect.TypeOf(deg1))
	fmt.Println(reflect.TypeOf(deg2))
	fmt.Println(reflect.TypeOf(deg3))
	fmt.Println(reflect.TypeOf(deg4))

}
