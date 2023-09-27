package slices

import "fmt"

func Demo1() {

	isimler := make([]string, 3)
	//fmt.Println(isimler)
	isimler[0] = "ömer"
	isimler[1] = "ömer"
	isimler[2] = "ömer"
	isimler = append(isimler, "ömer")
	fmt.Println(isimler)
	fmt.Println(len(isimler))

}
