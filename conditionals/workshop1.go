package conditionals

import (
	"fmt"
)

func Demo3() {
	//3 adet int değişken tanımla sayı olacak.
	//ekrana en büyük sayıyı yazdır.

	var sayi1, sayi2, sayi3 int = 10, 20, 30
	var enBuyuk int = sayi1

	if sayi2 > enBuyuk {
		enBuyuk = sayi2
	}
	if sayi3 > enBuyuk {
		enBuyuk = sayi3
	}

	fmt.Printf("En büyük sayı: %v", enBuyuk)

}
