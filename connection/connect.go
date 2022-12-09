package connection

import (
	"log"
	"projek/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDb() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/obdam_rsj"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.JobSpef{}, &models.Doctor{}, &models.Medicine{}, &models.Room{}, &models.Pasien{}, &models.Detaildoc{})

	return db
}
