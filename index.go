package main

import "fmt"

func main() {
	fmt.Println("halo")
	
	
// Type Data Slice
var aray = [...]string{"Al Tsaqif", "Deri", "Ucup", "Nadie", "Azka"}

	var slice = aray[2:4]
	fmt.Println(slice)
	fmt.Println(slice[0])
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	aray[0] = "Asep"
	fmt.Println(aray)

	slice[0] = "Agus"
	fmt.Println(slice)

	var apent = append(slice, "Zanan")
	apent[0] = "Philips"
	fmt.Println(apent)
	fmt.Println(len(apent))
	fmt.Println(cap(apent))
	fmt.Println(slice)
	fmt.Println(aray)
}
