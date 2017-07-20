package main

import (
	"fmt"
)

func main() {

	yas := 30

	switch yas {
	case 20:
		fmt.Println("Askerlik Çağındasın..")
	case 23:
		fmt.Println("İş Bulman Gerek ..")
	case 27:
		fmt.Println("Eş Bulman Gerek...")
	default:
		fmt.Println("Aradığımız Yaş Aralığında Değilsin...")
	}
}
