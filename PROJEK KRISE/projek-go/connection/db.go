package connection

import (
	"log"
	"projek-go/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/panti"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Kepengurusan{},&models.Donatur{}, &models.Sumbangan{}, &models.Sekolah{}, &models.Anak{}, &models.RiwayatKesehatan{},&models.Adopsi{}) 

	return db
}
