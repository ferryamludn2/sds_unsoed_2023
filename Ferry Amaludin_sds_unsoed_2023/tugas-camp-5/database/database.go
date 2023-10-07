package database

import (
	"belajar-backend/models"
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var db *sql.DB

func Connect() {

	dsn := "root@tcp(127.0.0.1:3307)/eldirufadhil?charset=utf8mb4&parseTime=True&loc=Local" //ini aku pake 3307 di mysql xampp sama phpmyadmin soalnya yang port 3306 dipake sama mysql workbench
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("database gak kehubung")
	}

	fmt.Println("database terhubung")

	fmt.Println(db)
	DB.AutoMigrate(models.Dosen{}, models.Mahasiswa{}, models.Tugas{})

}
