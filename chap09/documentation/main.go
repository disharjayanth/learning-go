package main

import (
	"fmt"

	"github.com/disharjayanth/learningGo/chap09/documentation/money"
)

func main() {
	var options int
	var amount float64
	var err error
	convertedMoney := money.Money{}
	fmt.Println("Enter the amount of $ to convert:")
	fmt.Scanf("%f", &amount)
	fmt.Println("1 - CAD, 2 - EU, 3 - INR")
	fmt.Scanf("%d", &options)
	if options == 1 {
		convertedMoney, err = money.Convert(money.Money{
			Currency: "USD",
			Value:    amount,
		}, "CAD")
	} else if options == 2 {
		convertedMoney, err = money.Convert(money.Money{
			Currency: "USD",
			Value:    amount,
		}, "EU")
	} else if options == 3 {
		convertedMoney, err = money.Convert(money.Money{
			Currency: "USD",
			Value:    amount,
		}, "INR")
	}
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(convertedMoney)
}
