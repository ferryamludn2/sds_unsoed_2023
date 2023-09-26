package main

import "fmt"

func main() {

	//membuat variabel array tempat wisata yang berisi 8 indeks tentang nama dari tempat wisat terkenal
	var wisata = [8]string{"Mieobelix hills", "safari to the sky", "linggarjati", "nepal van java", "kaliangkrik", "suko makmur", "monas", "baturadem"}
	//menampilkan 8 indeks array
	fmt.Println(wisata)
	//membuat slice1 yang terdiri dari 5 buah data
	slice1 := wisata[:5]
	//menampilkan slice1
	fmt.Println(slice1)
	//menambahkan 3 data ke slice1

	//menambahkan data baru ke isi slice1 dimasukkan ke variabel slicebaru
	slicebaru := append(slice1, "omah kembang", "kedung kayang", "silancur")
	//menampilkan slicebaru
	fmt.Println(slicebaru)
}
