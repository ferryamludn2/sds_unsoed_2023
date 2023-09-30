package main

import (
	"fmt"
)

// Struktur penggambaran orang
type Pengguna struct {
	ID           int
	NamaPengguna string
	Email        string
}

// informasi pengguna
func (u Pengguna) DapatkanInfo() string {
	return fmt.Sprintf("ID: %d\nNama Pengguna: %s\nEmail: %s", u.ID, u.NamaPengguna, u.Email)
}

// mengubah email pengguna dengan pointer
func (u *Pengguna) UbahEmail(emailBaru string) {
	u.Email = emailBaru
}

func main() {
	// Membuat instance pengguna dengan menggunakan pointer
	pengguna := &Pengguna{ID: 1, NamaPengguna: "ferryamaludin", Email: "ferryamludn@apaajadehya.com"}

	// Memanggil metode DapatkanInfo untuk mendapatkan informasi dari pengguna
	infoPengguna := pengguna.DapatkanInfo()
	fmt.Println("Informasi Pengguna Sebelum Perubahan Email:")
	fmt.Println(infoPengguna)

	// Mengubah email pengguna dengan menggunakan metode UbahEmail
	pengguna.UbahEmail("ferrysiapaaja@apaajadehya.com")

	// Memanggil metode DapatkanInfo lagi setelah mengubah email
	infoPengguna = pengguna.DapatkanInfo()
	fmt.Println("\nInformasi Pengguna Setelah Perubahan Email:")
	fmt.Println(infoPengguna)
}
