package main

import "fmt"

func luasPersegi(sisi1 float64, sisi2 float64) {
	fmt.Println("sisi 1 : ", sisi1)
	fmt.Println("sisi 2 : ", sisi2)
	fmt.Println("area   : ", sisi1*sisi2, "m^2")
}

func main() {
	fmt.Println("fungsi untuk menghitung area persegi ")
	//mengisi Sisi ke 1 dan Sisi ke 2
	luasPersegi(32.1231, 10.512)

}
