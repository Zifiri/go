package main

import (
	"fmt"
)

func main() {

	// Array Key value
	arrayMap := make(map[string]int)
	arrayMap["Key1"] = 1
	arrayMap["Key2"] = 2
	arrayMap["Key3"] = 3
	arrayMap["Key4"] = 4
	arrayMap["Key5"] = 5
	fmt.Println(arrayMap)

	// Silme Delete
	delete(arrayMap, "Key2") // Key2 olanı diziden sil...append
	fmt.Println("Silindi mi:", arrayMap)

	// Key Value 2.Yol

	arrayMap2 := map[string]bool{"Map": true, "Dizi": false}
	fmt.Println(arrayMap2)

}
