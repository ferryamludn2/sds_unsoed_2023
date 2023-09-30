package main

import (
	"fmt"
)

// Struktur untuk menggambarkan seorang pengguna
type User struct {
	ID       int
	Pengguna string
	Email    string
}

// Metode untuk mengeluarkan informasi pengguna
func (u User) DapatkanInfo() string {
	return fmt.Sprintf("ID: %d\nPengguna: %s\nEmail: %s", u.ID, u.Pengguna, u.Email)
}

// Struktur untuk merepresentasikan produk
type Product struct {
	ID    int
	Nama  string
	Harga float64
}

// Metode untuk menghitung harga produk setelah diskon
func (p Product) HitungHargaSetelahDiskon(persenDiskon float64) float64 {
	hargaSetelahDiskon := p.Harga - (p.Harga * persenDiskon / 100)
	return hargaSetelahDiskon
}

func main() {
	// Membuat instance pengguna
	pengguna := User{ID: 1, Pengguna: "Ferry Amaludin", Email: "ferryganteng@pedenomersatu.com"}

	// Memanggil metode DapatkanInfo untuk mendapatkan informasi pengguna
	infoPengguna := pengguna.DapatkanInfo()
	fmt.Println("Informasi Pengguna:")
	fmt.Println(infoPengguna)

	// Membuat instance produk
	produk := Product{ID: 102, Nama: "Laptop ROG", Harga: 15000000}

	// Memanggil metode HitungHargaSetelahDiskon untuk menghitung harga produk setelah diskon
	hargaSetelahDiskon := produk.HitungHargaSetelahDiskon(10)
	fmt.Printf("\nHarga Produk Setelah Diskon: Rp.%.2f\n", hargaSetelahDiskon)
}
