package main

import "fmt"

type person struct {
	name    string
	age     int
	license bool
}

func main() {
	var john person = person{
		name:    "john",
		age:     20,
		license: true,
	}

	fmt.Println(john)

	sam := person{
		name:    "sam",
		age:     24,
		license: false,
	}

	fmt.Println(sam)

	var address struct {
		place       string
		landmark    string
		pincode     int
		phonenumber int
	}

	address.place = "St John's"
	address.landmark = "Hospital"
	address.pincode = 5423232
	address.phonenumber = 543232244
	fmt.Println(address)
}
