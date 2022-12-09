package connection

import (
	"log"
	"projek_infak/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDb() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/infak_pub"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Angkatan{})
	db.AutoMigrate(&models.Prodi{})
	db.AutoMigrate(&models.Alumni{})
	db.AutoMigrate(&models.Rekening{})
	db.AutoMigrate(&models.Infak{})

	return db
}
