package loops

import "fmt"

//kullanıcıdan bir sayı girmesini iste
//girilen sayının asal olup olmadığını bulan program yaz
//2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97

// asalSayilar := [25]int {2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
func Demo5() {
	girilenSayi := 0
	fmt.Println("Bir sayı giriniz")
	fmt.Scanln(&girilenSayi)
	asalMi := true
	for i := 2; i < girilenSayi; i++ {
		if girilenSayi%i == 0 {
			asalMi = false
		}
	}

	if asalMi == true {
		fmt.Println("Girilen sayı asaldır")
	} else {
		fmt.Println("Girilen sayı asal değildir")

	}

}
