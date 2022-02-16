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

	Philps := [3]int{
		1,
		2,
		3,
	}

	if Philps[0] == Philps[2] {
		fmt.Println("Hay Bro")
	} else if Philps[1] == Philps[2] {
		fmt.Println("Hay Bre")
	} else if Philps[2] == Philps[2] {
		fmt.Println("Hay Bray")
	} else {
		fmt.Println("Hay Ngab")
	}
	
}
