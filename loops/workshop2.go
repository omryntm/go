package loops

import "fmt"

func Demo4() {
	aklimdakiSayi := 50
	tahminEdilensayi := 0
	tahminSayisi := 0
	fmt.Println("Bir sayi tahmin ediniz")
	fmt.Scanln(&tahminEdilensayi)
	fmt.Println("Tahmin edilen sayi: ", tahminEdilensayi)
	for tahminEdilensayi != aklimdakiSayi {
		if tahminEdilensayi < aklimdakiSayi {
			fmt.Println("Daha buyuk bir sayi tahmin ediniz")
			fmt.Scanln(&tahminEdilensayi)
			tahminSayisi = tahminSayisi + 1
		} else {
			fmt.Println("Daha kucuk bir sayi tahmin ediniz")
			fmt.Scanln(&tahminEdilensayi)
			tahminSayisi = tahminSayisi + 1
		}

	}
	basariDurumu := ""
	if tahminSayisi > 0 || tahminSayisi <= 3 {
		basariDurumu = "Süper"
	} else if tahminSayisi <= 6 {
		basariDurumu = "Güzel"
	} else {
		basariDurumu = "Fena Değil"
	}
	fmt.Printf("Tebrikler %v tahmin  : %v ", tahminSayisi, basariDurumu)

}
