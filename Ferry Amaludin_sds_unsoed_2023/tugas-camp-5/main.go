package main

import (
	"belajar-backend/database"
	"belajar-backend/route"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	database.Connect()

	log.Fatal(app.Listen(":3000"))

	app.Post("/insertdosen", route.CreateDosen)
	app.Post("/insertmahasiswa", route.CreateMahasiswa)
	app.Post("/inserttugas", route.CreateTugas)
	app.Get("/getDataDosen", route.GetDosen)
	app.Get("/getDataMahasiswa", route.GetMahasiswa)
	app.Get("/getDataTugas", route.GetTugas)
	app.Put("/update/:id_dosen", route.UpdateDosen)
	app.Put("/update/:id_mahasiswa", route.UpdateMahasiswa)
	app.Put("/update/:id_tugas", route.UpdateTugas)
	app.Get("/delete/:id_dosen", route.DeleteDosen)
	app.Get("/delete/:id_mahasiswa", route.DeleteMahasiswa)
	app.Get("/delete/:id_tugas", route.DeleteTugas)

}
