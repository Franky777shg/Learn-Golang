package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	/* Number sama aja kaya javascript, tapi ada beberapa tipe data
	dan alias yang beda, gak perlu dihafal karena teori nya bisa liat di internet. */

	/* Boolean juga sama persis kaya javascript */

	/* String selalu dibungkus dengan kutip dua (") */
	fmt.Println("Frengky")
	fmt.Println(len("Frengky")) // hitung karakter di sebuah string
	fmt.Println("Frengky"[0])   // ambil ascii number dari sebuah string sesuai index

	/* Variable START
	Menggunakan keyword var dan wajib menyebutkan tipe data nya.
	Tapi jika langsung di assigned isi datanya, maka tidak perlu di kasih tipe data.*/
	var name string
	name = "Frengky"
	fmt.Println(name)

	var name2 = "Andi"
	fmt.Println(name2)

	/* Bisa tidak menggunakan var tapi memakai simbol := */
	name3 := "Lala"
	fmt.Println(name3)

	/* Bikin sekaligus variabel banyak*/
	var (
		name4 = "Budi"
		name5 = "Eko"
	)
	fmt.Println(name4)
	fmt.Println(name5)
}
