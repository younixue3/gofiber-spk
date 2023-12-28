package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-spk/controller"
)

func SetupRoutes(app *fiber.App) {
	//API GROUP
	api := app.Group("/api")
	//Mahasiswa
	api.Get("/mahasiswa", controller.Mahasiswas)
	api.Get("/mahasiswa/:pk", controller.Mahasiswa)
	api.Post("/mahasiswa", controller.CreateMahasiswa)
	api.Put("/mahasiswa/:pk", controller.UpdateMahasiswa)
	api.Delete("/mahasiswa/:pk", controller.DeleteMahasiswa)
	//Kategori
	api.Get("/kategori", controller.Kategoris)
	api.Get("/kategori/:pk", controller.Kategori)
	api.Post("/kategori", controller.CreateKategori)
	api.Put("/kategori/:pk", controller.UpdateKategori)
	api.Delete("/kategori/:pk", controller.DeleteKategori)
	//Bobot Kategori
	api.Get("/bobot-kategori", controller.BobotKategoris)
	api.Get("/bobot-kategori/:pk", controller.BobotKategori)
	api.Post("/bobot-kategori", controller.CreateBobotKategori)
	api.Put("/bobot-kategori/:pk", controller.UpdateBobotKategori)
	api.Delete("/bobot-kategori/:pk", controller.DeleteBobotKategori)
	//Nilai
	api.Get("/nilai-mahasiswa", controller.NilaiMahasiswas)
	api.Get("/nilai-mahasiswa/:pk", controller.NilaiMahasiswa)
	api.Post("/nilai-mahasiswa", controller.CreateNilaiMahasiswa)
	api.Put("/nilai-mahasiswa/:pk", controller.UpdateNilaiMahasiswa)
	api.Delete("/nilai-mahasiswa/:pk", controller.DeleteNilaiMahasiswa)

	//api.Get("/book", controller.Books)
	//api.Get("/book/:pk", controller.Book)
	//api.Post("/book", controller.CreateBook)
	//api.Put("/book/:pk", controller.UpdateBook)
	//api.Delete("/book/:pk", controller.DeleteBook)
}
