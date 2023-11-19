package main

import "fmt"

func main() {
	// fmt.Println("Hello World!")

	// data number type
	// fmt.Println("satu", 1)
	// fmt.Println("dua", 2)
	// fmt.Println("tiga koma lima", 3.5)

	// func string
	// fmt.Println(len("franky"))
	// fmt.Println("franky"[2])

	// variable
	// var name string

	// name = "franky"
	// fmt.Println(name)

	// name = "frankyshg"
	// fmt.Println(name)

	// var name2 = "yoona"
	// fmt.Println(name2)

	// name2 = "yoona yoona"
	// fmt.Println(name2)

	// name3 := "pasti bisa"
	// fmt.Println(name3)

	// var (
	// 	firstName = "Frengky"
	// 	lastName  = "Sihombing"
	// )

	// fmt.Println(firstName)
	// fmt.Println(lastName)

	// constant
	// const firstName = "franky"
	// const lastName = "sihombing"
	// firstName := "franky" => same as var
	// firstName = "asd" => can't re assign to constant
	// const (
	// 	car1 = "ferrary"
	// 	car2 = "tesla"
	// 	car3 = "porsche"
	// )
	// fmt.Println(firstName, lastName)
	// fmt.Println(car1, car2, car3)

	// konversi tipe data
	// var nilai32 int32 = 100000
	// var nilai64 int64 = int64(nilai32)
	// var nilai8 int8 = int8(nilai64) // -96, karena melebihi limit int8 yaitu 127
	// fmt.Println(nilai32)
	// fmt.Println(nilai64)
	// fmt.Println(nilai8)

	// konversi tipe data 2
	// var name = "frengky"
	// var e = name[2]
	// fmt.Println((e))
	// fmt.Println((string(e))) // string() merubah byte menjadi character

	// type noKTP string

	// var serialNumber noKTP = "6amd730cnd7"
	// fmt.Println(serialNumber)

	// Boolean sama seperti javascript
	// &&, ||, !

	// var nilaiAkhir = 90
	// var nilaiAbsen = 80

	// // var lulusNilaiAkhir = nilaiAkhir >= 80
	// // var lulusNilaiAbsen = nilaiAbsen >= 80

	// // var lulus = lulusNilaiAbsen && lulusNilaiAkhir
	// // fmt.Println(lulus)
	// fmt.Println(nilaiAkhir >= 80 && nilaiAbsen >= 80)

	// Array
	// var names [3]string
	// names[0] = "lala"
	// names[1] = "lili"
	// names[2] = "lulu"
	// // names[3] = "lulu" => error
	// fmt.Println(names)

	// var numbers = [3]int{
	// 	1, 2, 3,
	// }
	// fmt.Println(numbers)
	// fmt.Println(len(numbers))
	// numbers[1] = 4
	// fmt.Println(numbers)

	// months := [...]string{
	// 	"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
	// }

	// slice1 := months[4:7]
	// fmt.Println(slice1)
	// fmt.Println(len(slice1))
	// fmt.Println(cap(slice1))

	// // saat array asal diubah, slice juga ikut berubah
	// months[5] = "diubah"
	// fmt.Println(slice1)

	// // saat slice berubah pun, array asal juga akan ikut berubah
	// slice1[0] = "May lagi"
	// fmt.Println(months)

	// slice2 := months[2:4]
	// fmt.Println(slice2)
	// // ketika kita melakukan append menggunakan slice yang kapasitas nya sudah penuh, maka dia akan membuat array baru
	// // sehingga array sebelum2nya tidak akan ikut terubah
	// // tetapi jika kita melakukan append pada slice yang kapasitas nya belum penuh, maka dia tidak akan membuat array baru
	// // sehingga array sebelum2nya akan ikut berubah
	// slice3 := append(slice2, "pengganti may", "pengganti jun")
	// fmt.Println(slice3)
	// slice3[1] = "Bukan apr"
	// fmt.Println(slice3)
	// fmt.Println(slice2)
	// fmt.Println(months)

	// newArray dengan "make"
	// newSlice := make([]string, 2, 5)
	// newSlice[0] = "Nice"
	// newSlice[1] = "Good"
	// fmt.Println(newSlice)
	// fmt.Println(len(newSlice))
	// fmt.Println(cap(newSlice))

	// copy slice dengan "copy"
	// copySlice := make([]string, 1, cap(newSlice))
	// copy(copySlice, newSlice)
	// fmt.Println(copySlice)

	// tipe data map
	// person := map[string]string{
	// 	"name":    "frengky",
	// 	"address": "jakarta",
	// }
	// fmt.Println(person["name"])

	// car := map[int]string{
	// 	1: "satu",
	// 	2: "dua",
	// }
	// fmt.Println(car[2])

	// if else same with javascript but with extra feature which is short if

	// for loop
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println(counter)
	// 	counter++
	// }

	// code di atas dapat diubah menjadi bentuk lebih singkat seperti di bawah
	// for counter := 1; counter <= 10; counter++ {
	// 	fmt.Println(counter)
	// }

	// looping collection of data
	slice := []string{"satu", "dua", "tiga"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for index, value := range slice {
		fmt.Println(index, value)
	}
	// without index
	for _, value := range slice {
		fmt.Println(value)
	}
}
