package models

type Lapangan struct {
	NoLapangan int         `gorm:"primaryKey" json:"no_lapangan"`
	Harga      int         `gorm:"Not Null;" json:"harga"`
	Images     string      `gorm:"Not Null" json:"image"`
	Booking    []Booking   `gorm:"foreignkey:no_lap;"`
	Riwayat    []Riwayat   `gorm:"foreignKey:no_lap"`
	Transaksi  []Transaksi `gorm:"foreignKey:no_lap"`
}
