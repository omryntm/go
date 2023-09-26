package variables

import "fmt"

func Demo1() {
	fmt.Println("Hello, World!")
	var metin string = "Merhaba Dünya!"
	fmt.Println(metin)
	var kdv int = 10
	fmt.Println(kdv)
	var kdv2 float64 = 0.10
	fmt.Println(kdv2)

	fmt.Println(100 + 100*kdv2)
	var durum bool
	var name string = "Mehmet"
	var name2 string = "hasan"
	durum = name == name2
	fmt.Println(durum)
}
