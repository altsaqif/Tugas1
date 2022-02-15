package main

import "fmt"

func main() {
	fmt.Println("halo")
	
	Philips := make(map[string]string)
	Philips["Nama"] = "Philips"
	Philips["Alamat"] = "Medan"
	fmt.Println(Philips)
	fmt.Println(Philips["Nama"])
	fmt.Println(Philips["Alamat"])
}
