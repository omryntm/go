package functions

func DortIslem(sayi1 int, sayi2 int) (int, int, int, float32) {
	var toplam = sayi1 + sayi2
	var fark = sayi1 - sayi2
	var carpim = sayi1 * sayi2
	var bolum = float32(sayi1 / sayi2)
	return toplam, fark, carpim, bolum

}
