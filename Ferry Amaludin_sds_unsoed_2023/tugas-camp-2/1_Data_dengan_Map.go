package main

import (
	"fmt"
)

func main() {
	// Membuat map kosong sebagai tempat penyimpanan data pengguna
	Nama := make(map[int]string)

	// Menambahkan data pengguna ke dalam map
	Nama[1] = "Ferry"
	Nama[2] = "Ferroy"
	Nama[3] = "Fernaldo"

	// Menampilkan data pengguna dari map
	fmt.Println("Data Pengguna:")
	fmt.Println(Nama)

	// Mengakses data pengguna berdasarkan kunci
	fmt.Println("Pengguna dengan ID 2:", Nama[2])

	// Memeriksa apakah sebuah kunci ada dalam map
	_, isExists := Nama[4]
	if isExists {
		fmt.Println("Pengguna dengan ID 4 ditemukan")
	} else {
		fmt.Println("Pengguna dengan ID 4 tidak ditemukan")
	}
	//Ferry Amaludin
	// Menghapus data pengguna dari map
	delete(Nama, 3)
	fmt.Println("Setelah menghapus data pengguna dengan ID 3:")
	fmt.Println(Nama)
	//Ferry Amaludin
	// Inisialisasi map dengan harga tiket
	Wisata := map[string]int{
		"Sukomakmur":   5000,
		"Kaliangkrik":  5000,
		"NepalVanJava": 5000,
	}
	//Ferry Amaludin
	// Menampilkan data wisata yang ada
	fmt.Println("Harga tiket wisata:")
	fmt.Println(Wisata)
}
