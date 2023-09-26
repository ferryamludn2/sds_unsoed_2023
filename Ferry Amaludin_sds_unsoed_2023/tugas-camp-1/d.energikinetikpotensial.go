package main

import "fmt"

func Kinetik(massa float64, kecepatan float64) {
	fmt.Println("dengan kecepatan", kecepatan, " m/s dan massa dari sebuah benda sebesar", massa, "kg maka energi kinetik yang dihasilkan adalah :", 0.5*massa*kecepatan*kecepatan, "Joule")
}

func Potensial(massa float64, tinggi float64) {
	fmt.Println("Dengan massa", massa, "kg dan tinggi antara tanah dengan benda", tinggi, "m dan gravitasi di bumi sebesar 9,8 m/s^2 maka energi potensial yang dihasilkan adalah :", massa*9.8*tinggi, "Joule")
}

func main() {
	fmt.Println("Contoh Energi Kinetik : ")
	//mengisi dengan massa dan kecepatan
	Kinetik(2, 4)
	fmt.Println("Contoh Energi Potensial : ")
	//mengisi dengan massa dan tinggi
	Potensial(5, 10)
}
