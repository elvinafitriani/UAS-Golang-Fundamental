package database

import (
	"ProjectFutsal/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() *gorm.DB {
	var err error
	db := "root:@tcp(127.0.0.1:3306)/akbar_futsal?parseTime=true"

	Db, err = gorm.Open(mysql.Open(db), &gorm.Config{})
	if err != nil {
		log.Fatal("Can't Connect To Database")
	}

	Db.AutoMigrate(&models.Lapangan{}, &models.Transaksi{}, &models.Booking{}, &models.Riwayat{})

	return Db
}
