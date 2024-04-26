package main

import "fmt"

// Struct sama dengan class di javascript

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello ", name, "My name is ", customer.Name)
}

func main() {
	// var person1 Customer
	// person1.Name = "Ali"
	// person1.Address = "Jakarta"
	// person1.Age = 13
	// fmt.Println(person1)

	// var person2 = Customer{
	// 	Name:    "Jaber",
	// 	Address: "Bandung",
	// 	Age:     10,
	// }
	// fmt.Println(person2)

	var person3 Customer
	person3.Name = "Dono"
	person3.sayHello("Budi")
}
