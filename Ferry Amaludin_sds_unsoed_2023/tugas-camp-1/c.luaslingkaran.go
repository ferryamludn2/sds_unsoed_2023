package main

import "fmt"

func luasLingkaran(radius float64) {
	fmt.Println("Radius : ", radius)
	fmt.Println("Area   : ", 3.14*radius*radius, "m^2")

}

func main() {
	fmt.Println("menghitung luas dari lingkaran")
	//mengisi dengan radius lingkaran
	luasLingkaran(4)
}
