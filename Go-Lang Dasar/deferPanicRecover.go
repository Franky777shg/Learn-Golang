// package main

// import "fmt"

// func logging() {
// 	fmt.Println("selesai")
// 	// Add on for recover example
// 	message := recover()
// 	fmt.Println("data error", message)
// }

// // // Defer use for run something when a function finish, even error happen
// // func runningApp() {
// // 	defer logging()
// // 	fmt.Println("running application")
// // }

// // func main() {
// // 	runningApp()
// // }

// // // Panic use for stop a function when error happen
// // func runningApp(error bool) {
// // 	defer logging()
// // 	if error {
// // 		panic("ERROR")
// // 	}
// // 	fmt.Println("running app")
// // }

// // func main() {
// // 	runningApp(false)
// // }

// // Recover use for catch error data from panic
// func runningApp(error bool) {
// 	defer logging()
// 	if error {
// 		panic("ERROR")
// 	}
// }

// func main() {
// 	runningApp(true)
// }
