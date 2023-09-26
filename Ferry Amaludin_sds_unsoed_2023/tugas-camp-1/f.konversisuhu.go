package main

import "fmt"

// untuk mengkonversi suhu dari Celsius ke Fahrenheit
func celsiusKeFahrenheit(celsius float64) {
	fmt.Println(celsius, "Celsius =", (celsius*9/5)+32, "Fahrenheit")
}

// untuk mengkonversi suhu dari Fahrenheit ke Celsius
func fahrenheitKeCelsius(fahrenheit float64) {
	fmt.Println(fahrenheit, "Fahrenheit =", (fahrenheit-32)*5/9, "Celsius")

}

// untuk mengkonversi suhu dari Celsius ke Kelvin
func celsiusKeKelvin(celsius float64) {
	fmt.Println(celsius, "%.2f Celsius =", celsius+273.15, "Kelvin")

}

// untuk mengkonversi suhu dari Kelvin ke Celsius
func kelvinKeCelsius(kelvin float64) {
	fmt.Println(kelvin, "Kelvin =", kelvin-273.15, "Celsius")

}

// untuk mengkonversi suhu dari Celsius ke Réaumur
func celsiusKeReaumur(celsius float64) {
	fmt.Println(celsius, "Celsius =", celsius*4/5, "Réaumur")
}

// untuk mengkonversi suhu dari Réaumur ke Celsius
func reaumurKeCelsius(reaumur float64) {
	fmt.Println(reaumur, "Réaumur =", reaumur*5/4, "Celsius")
}

func main() {
	fmt.Println("Konversi Suhu")

	// Konversi Celsius ke Fahrenheit
	celsiusKeFahrenheit(30)

	// Konversi Fahrenheit ke Celsius

	fahrenheitKeCelsius(45)

	// // Konversi Celsius ke Kelvin

	celsiusKeKelvin(50)

	// // Konversi Kelvin ke Celsius

	kelvinKeCelsius(400)

	// // Konversi Celsius ke Réaumur

	celsiusKeReaumur(80)

	// // Konversi Réaumur ke Celsius

	reaumurKeCelsius(40)

}
