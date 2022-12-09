package connection

import(
	"project/models"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDb() *gorm.DB{
	dsn := "root:@tcp(127.0.0.1:3306)/akademik"

	db, err :=gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.Matkul{})
	db.AutoMigrate(&models.Mahasiswa{})
	db.AutoMigrate(&models.Khs{})
	db.AutoMigrate(&models.Dosen{})
	db.AutoMigrate(&models.Jadwal{})

	return db
}