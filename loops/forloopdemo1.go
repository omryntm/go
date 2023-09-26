package loops

import "fmt"

func Demo1() {
	var metin string = "Merhaba Dünya"
	//fmt.Println(metin)
	//fmt.Println(metin)
	//fmt.Println(metin)
	//fmt.Println(metin)
	//fmt.Println(metin)

	i := 1
	for i <= 5 {
		fmt.Println(metin)
		i = i + 1
		//i++
	}
	fmt.Println("Döngü bitti")
}
