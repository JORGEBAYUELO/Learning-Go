package main

import "fmt"

const PI float64 = 3.14

// func main() {
// 	fmt.Println("Hello World!")
// }

//Declaring Variables//

// func main() {
// 	var greeting string = "Hello World"
// 	fmt.Println(greeting)
// }

//Print variable and string//

// func main() {
// 	var name string = "KodeKloud"
// 	var user string = "Harry"

// 	fmt.Print("Welcome to ", name, ", ",user)
// }

//Printing on newline//

// func main() {
// 	var name string = "KodeKloud"
// 	var user string = "Harry"

// 	fmt.Println(name)
// 	fmt.Println(user)
// }

// func main() {
// 	var radius float64 = 5.154
// 	var area float64

// 	area = PI * radius * radius
// 	fmt.Printf("Radius: %.2f \nPI:%.1f \n", radius, PI)
// 	fmt.Printf("Area of Circle is : %f", area)
// 	fmt.Println("Thank You")
// }

//Bitwise Operators//

// func main() {
// 	var x, y int = 100,90
// 	fmt.Println(x & y)
// 	fmt.Println(x | y)
// }

// func main() {
// 	var x, y int = 100,90
// 	fmt.Println((x+y) >> 2)
// }

func main() {
	var x, y int = 100,90
	fmt.Println(!(((x+y) >> 2 ) == 47))
}