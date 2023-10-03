package structs

import "fmt"

func Demo1() {
	fmt.Println(product{"bilgisayar", 500, "xyz marka"})
	fmt.Println(product{"bilgisayar", 500, ""})
	fmt.Println(product{name: "bilgisayar", unitPrice: 500, brand: ""})

}

type product struct {
	name      string
	unitPrice float64
	brand     string
	//discountRate int
}
