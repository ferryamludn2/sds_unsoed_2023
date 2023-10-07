package models

type Dosen struct {
	Id_dosen   uint `json:"id_dosen" gorm:"primaryKey;autoIncrement:true"`
	Phone      string
	Email      string
	nama_dosen string
	Password   string
	jurusan    string
}

type Mahasiswa struct {
	Id_mahasiswa   uint `json:"id_mahasiswa" gorm:"primaryKey;autoIncrement:true"`
	Phone          string
	Email          string
	nama_mahasiswa string
	Password       string
	jurusan        string
}

type Tugas struct {
	Id_tugas   uint `json:"id_user" gorm:"primaryKey;autoIncrement:true"`
	Matakuliah string
	deadline   string
}
