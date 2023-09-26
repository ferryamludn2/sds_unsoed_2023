package main

import "fmt"

func cariFrekuensi(periode float64) {
	fmt.Println("dengan Periode", periode, "sekon maka frekuensinya adalah :", 1/periode, "Hz")
}

func main() {
	fmt.Println("Mencari Frekuensi: ")
	//mengisi periode dalam satuan sekon
	cariFrekuensi(4)
}
