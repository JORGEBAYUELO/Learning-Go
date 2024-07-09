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

// func main() {
// 	var x, y int = 100,90
// 	fmt.Println(!(((x+y) >> 2 ) == 47))
// }

//If, Else, Else-if, Switch//

// func main() {
// 	var a, b = 100, 5
// 	switch {
// 	case a/b == 10:
// 		fmt.Println("10")
// 	case a/b == 20:
// 		fmt.Println("20")
// 	case a/b == 10:
// 		fmt.Println("30")
// 	default:
// 		fmt.Println("default") 
		
// 	}
// }

// func main() {
// 	var a, b = 100, 5
// 	switch a {
// 	case a/b == 10:
// 		fmt.Println("10")
// 	case a/b == 20:
// 		fmt.Println("20")
// 	case a/b == 10:
// 		fmt.Println("30")
// 	default:
// 		fmt.Println("default") 
		
// 	}
// }

//Arrays & Slice//

// func main() {
// 	arr := []int{-1, -2}
// 	for _, value := range arr {
// 		fmt.Println(value)
// 	}
// }

// func main() {
// 	arr := [5]string{"a", "b", "c", "d", "e"}
// 	slice := arr[:4]
// 	fmt.Println(arr)
// 	fmt.Println(slice)
// 	slice[1] = "x"
// 	fmt.Println(arr)
// 	fmt.Println(slice)
// }

// func main() {
// 	arr := [5]int{10, 20, 90, 70, 60}
// 	slice := arr[:3]
// 	fmt.Println(cap(slice))
// 	new_slice := append(slice, 100, 200)
// 	fmt.Println(cap(new_slice))
// }

// func main() {
// 	arr := [5]int{10, 20, 90, 70, 60}
// 	slice := arr[:3]
// 	fmt.Println(cap(slice))

// 	slice_2 := make([]int, 5, 20)
// 	new_slice := append(slice, slice_2...)
// 	fmt.Println(cap(new_slice))

// }

// func main() {
// 	arr := [5]int{10, 20, 90, 70, 60}
// 	slice := append(arr[:2], arr[3:])
// 	fmt.Println(slice)
// }

// func main() {
// 	arr := []int{10, 20, 90, 70, 60}
// 	slice := make([]int, 10)
// 	num := copy(slice, arr)
// 	fmt.Println(slice)
// 	fmt.Println(num)
// }

// func main() {
// 	arr := []int{10, 20, 90, 70, 60}
// 	slice := make([]int, 10)
// 	copy(slice, arr)
// 	slice[1] = 1000
// 	fmt.Println(arr)
// 	fmt.Println(slice)
// }

func main() {
	arr := [10]int{10, 20}
	slice := arr[2:8]
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}