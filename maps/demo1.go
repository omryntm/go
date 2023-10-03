package maps

import "fmt"

func Demo1() {
	// key-value mimarisi üzerine kurulmuş basit yapılar
	sozluk := make(map[string]string)
	sozluk["table"] = "masa"
	sozluk["book"] = "kitap"
	sozluk["pencil"] = "kalem"

	fmt.Println(sozluk["book"])
	fmt.Println("Eleman sayısı: ", len(sozluk))
	fmt.Println("Sözlük: ", sozluk)
	delete(sozluk, "book")
	fmt.Println("Eleman sayısı: ", len(sozluk))
	fmt.Println("Sözlük: ", sozluk)
	deger, varMi := sozluk["pencil"]
	fmt.Println(deger)
	fmt.Println("Listede olma durumu: ", varMi)
	sozluk2 := map[string]string{"glass": "bardak", "microphone": "mikrofon", "ball": "top"}
	fmt.Println(sozluk2)

}
