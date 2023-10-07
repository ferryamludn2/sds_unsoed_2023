package route

import (
	"belajar-backend/database"
	"belajar-backend/models"

	"github.com/gofiber/fiber/v2"
)

///data akan dimasukan ke data base

func CreateMahasiswa(c *fiber.Ctx) error {

	var data map[string]string

	mahasiswa := models.Mahasiswa{
		Id_mahasiswa: 0,
		Phone:        data["phone"],
		Email:        data["email"],
		Password:     data["password"],
	}

	database.DB.Create(mahasiswa)

	return c.JSON(fiber.Map{})
}

func CreateDosen(c *fiber.Ctx) error {

	var data map[string]string

	dosen := models.Dosen{
		Phone:      data["phone"],
		Email:      data["email"],
		Nama_dosen: data["nama_dosen"],
		Password:   data["password"],
		Jurusan:    data["jurusan"],
	}

	database.DB.Create(dosen)

	return c.JSON(fiber.Map{})
}

func CreateTugas(c *fiber.Ctx) error {

	var data map[string]string

	tugas := models.Tugas{
		Matakuliah:  data["mata_kuliah"],
		Pengumpulan: data["batas_akhir"],
	}

	database.DB.Create(tugas)

	return c.JSON(fiber.Map{})
}

func UpdateMahasiswa(c *fiber.Ctx) error {

	id_mahasiswa := c.Params("id_mahasiswa")

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var mahasiswa models.Mahasiswa
	database.DB.Find(&mahasiswa)

	update := models.Mahasiswa{
		Phone:          data["phone"],
		Email:          data["email"],
		Nama_mahasiswa: data["nama_mahasiswa"],
		Password:       data["password"],
		Jurusan:        data["jurusan"],
	}
	//mengambil database untuk di update

	database.DB.Model(&mahasiswa).Where("id_mahasiswa = ?", id_mahasiswa).Updates(update)

	return c.JSON(fiber.Map{
		"Pesan": "data mahasiswa terupdate",
	})
}

func UpdateDosen(c *fiber.Ctx) error {

	id_dosen := c.Params("id_dosen")

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var dosen models.Dosen
	database.DB.Find(&dosen)

	update := models.Dosen{
		Phone:      data["phone"],
		Email:      data["email"],
		Nama_Dosen: data["nama_dosen"],
		Password:   data["password"],
		Jurusan:    data["jurusan"],
	}
	//mengambil database untuk di update

	database.DB.Model(&dosen).Where("id_dosen = ?", id_dosen).Updates(update)

	return c.JSON(fiber.Map{
		"Pesan": "data dosen terupdate",
	})
}

func UpdateTugas(c *fiber.Ctx) error {

	id_tugas := c.Params("id_Tugas")

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var tugas models.Tugas
	database.DB.Find(&tugas)

	update := models.Tugas{
		Matakuliah:  data["mata_kuliah"],
		Pengumpulan: data["batas_akhir"],
	}
	//mengambil database untuk di update

	database.DB.Model(&tugas).Where("id_tugas = ?", id_tugas).Updates(update)

	return c.JSON(fiber.Map{
		"Pesan": "data tugas terupdate",
	})
}

func GetMahasiswa(c *fiber.Ctx) error {
	var mahasiswa []models.Mahasiswa
	database.DB.Find(&mahasiswa)
	return c.JSON(fiber.Map{
		"data": mahasiswa,
	})
}

func GetDosen(c *fiber.Ctx) error {
	var dosen []models.Dosen
	database.DB.Find(&dosen)
	return c.JSON(fiber.Map{
		"data": dosen,
	})
}

func GetTugas(c *fiber.Ctx) error {
	var tugas []models.Tugas
	database.DB.Find(&tugas)
	return c.JSON(fiber.Map{
		"data": tugas,
	})
}

func DeleteDosen(c *fiber.Ctx) error {

	var dosen models.Dosen

	id_dosen := c.Params("id_dosen")

	database.DB.Where("id_user = ?", id_dosen).Delete(dosen)

	return c.JSON(fiber.Map{
		"Pesan": "data user terhapus",
	})

}

func DeleteMahasiswa(c *fiber.Ctx) error {

	var mahasiswa models.Mahasiswa

	id_mahasiswa := c.Params("id_mahasiswa")

	database.DB.Where("id_user = ?", id_mahasiswa).Delete(mahasiswa)

	return c.JSON(fiber.Map{
		"Pesan": "data mahasiswa terhapus",
	})

}

func DeleteTugas(c *fiber.Ctx) error {

	var tugas models.Tugas

	id_tugas := c.Params("id_tugas")

	database.DB.Where("id_user = ?", id_tugas).Delete(tugas)

	return c.JSON(fiber.Map{
		"Pesan": "data tugas terhapus",
	})

}
