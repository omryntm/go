package loops

import "fmt"

func Demo3() {
	aklimdakiSayi := 50
	tahminEdilensayi := 0
	fmt.Println("Bir sayi tahmin ediniz")
	fmt.Scanln(&tahminEdilensayi)
	fmt.Println("Tahmin edilen sayi: ", tahminEdilensayi)
	for tahminEdilensayi != aklimdakiSayi {
		if tahminEdilensayi < aklimdakiSayi {
			fmt.Println("Daha buyuk bir sayi tahmin ediniz")
			fmt.Scanln(&tahminEdilensayi)
		} else {
			fmt.Println("Daha kucuk bir sayi tahmin ediniz")
			fmt.Scanln(&tahminEdilensayi)
		}

	}
	fmt.Println("Tebrikler")
}
