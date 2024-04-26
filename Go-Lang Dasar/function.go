// package main

// import "fmt"

// // Standard Function
// func sayHello() {
// 	fmt.Println("Say Hello")
// }

// // With Param
// // The number of params and arguments must be the same
// func sayHello(firstName string, lastName string) {
// 	fmt.Println("Hello", firstName, lastName)
// }

// // Return Value
// func sayHello(name string) string {
// 	return "Hello " + name
// }

// // Return Multi Value
// func sayHello(name string) (string, int) {
// 	return "Hello " + name, 24
// }

// // Named Return Values
// func getCompleteName() (firstName, middleName, lastName string) {
// 	firstName = "Budi"
// 	middleName = "Agus"
// 	lastName = "Santoso"

// 	return firstName, middleName, lastName
// }

// // Variadic Function
// // ... is variabel argument
// func sumAll(numbers ...int) int {
// 	total := 0
// 	for _, item := range numbers {
// 		total += item
// 	}

// 	return total
// }

// Function as value is as same as javascript
// For example: car := thatFunction
// Dont use () because it will call the function

// func main() {
// 	// sayHello() // Standard Function

// 	// sayHello("Frengky", "Sihombing") // With Param

// 	// fmt.Println(sayHello("Frengky")) // Return Value

// 	// // Return Multi Value
// 	// greet, age := sayHello("Frengky")
// 	// fmt.Println(greet, age)
// 	// // Bisa diganti dengan _ jika ada variabel hasil return yang tidak dibutuhkan
// 	// dibutuhkan, _ := sayHello("frengky")
// 	// fmt.Println(dibutuhkan)

// 	// // Named Return Values
// 	// a, b, c := getCompleteName()
// 	// fmt.Println(a, b, c)

// 	// // Variadic Function
// 	// fmt.Println(sumAll(10, 1, 2, 3, 4, 5, 6))
// 	// // If already have a slice, we can change slice into var arg
// 	// numbers := []int{10, 10, 10}
// 	// fmt.Println(sumAll(numbers...))

// }

// // Function as Parameter
// // Type declaration can be useful if parameter declaration has long detail and for reuseable
// type Filter func(string) string

// func sayHelloWithFilter(name string, filter Filter) {
// 	filteredName := filter(name)
// 	fmt.Println("Hello", filteredName)
// }

// func spamFilter(name string) string {
// 	if name == "Anjing" {
// 		return "..."
// 	} else {
// 		return name
// 	}
// }

// func main() {
// 	sayHelloWithFilter("Frengky", spamFilter)
// 	sayHelloWithFilter("Anjing", spamFilter)
// }

// // Anonymous Function
// type Blacklist func(string) bool

// func registerName(name string, blacklist Blacklist) {
// 	if blacklist(name) {
// 		fmt.Println("You are blocked", name)
// 	} else {
// 		fmt.Println("You are welcome", name)
// 	}
// }

// func main() {
// 	blacklist := func(name string) bool {
// 		return name == "Anjing"
// 	}

// 	registerName("Frengky", blacklist)
// 	registerName("Anjing", blacklist)
// }

// Recursive Function
// // Using Loop
// func factorialLoop(number int) {
// 	result := 1

// 	for counter := 1; counter <= number; counter++ {
// 		result *= counter
// 	}

// 	fmt.Println(result)
// }

// func main() {
// 	factorialLoop(5)
// }

// Using Recursive
// func factorialLoop(number int) int {
// 	if number == 1 {
// 		return 1
// 	} else {
// 		return number * factorialLoop(number-1)
// 	}
// }

// func main() {
// 	fmt.Println(factorialLoop(5))
// }
