package database

import (
	"go-fiber-spk/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root@tcp(127.0.0.1:3306)/gofiber_spk?charset=utf8mb4&parseTime=True&loc=Local"
var Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

func ConnectDatabase() {
	if err != nil {
		panic("failed to connect database")
	}
}

func RunMigration() {
	Db.AutoMigrate(&models.Mahasiswa{}, &models.Kategori{}, &models.BobotKategori{}, &models.NilaiMahasiswa{})
}
