package interfaces

import "fmt"

type Mortgage struct {
	crediPaymnetTotal float64
	address           string
	rate              float64
}

type Car struct {
	crediPaymnetTotal float64
	carInfo           string
	rate              float64
}

type CreditCalculator interface {
	Calculate() float64
}

func (m Mortgage) Calculate() float64 {
	return m.crediPaymnetTotal * m.rate / 100 / 12
}

func (c Car) Calculate() float64 {
	return c.crediPaymnetTotal * c.rate / 100 / 12
}

func CalculateMonthlyPayment(credits []CreditCalculator) float64 {
	paymentTotal := 0.0
	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + credits[i].Calculate()

	}
	return paymentTotal
}

func Demo2() {
	credit1 := Mortgage{rate: 10, crediPaymnetTotal: 100000, address: "Ankara"}
	credit2 := Mortgage{rate: 12, crediPaymnetTotal: 50000, address: "İstanbul"}
	credit3 := Car{rate: 15, crediPaymnetTotal: 60000, carInfo: "BMW"}
	credits := []CreditCalculator{credit1, credit2, credit3}
	total := CalculateMonthlyPayment(credits)
	fmt.Println("Toplam Ödeme: ", total)
}
