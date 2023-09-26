package main

import "fmt"

func luassegitiga(alas float64, tinggi float64) {
	fmt.Println("alas   : ", alas)
	fmt.Println("tinggi : ", tinggi)
	fmt.Println("area   : ", (alas*tinggi)/2, "m^2")
}

func main() {
	fmt.Println("Fungsi untuk menghitung area segitiga ")
	//mengisi alas dan tinggi
	luassegitiga(6, 2)

}
