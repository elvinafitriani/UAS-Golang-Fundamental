package models

import "time"

type Transaksi struct {
	IdTransaksi      int       `gorm:"primaryKey;Not Null" json:"id_transaksi"`
	NoLap            int       `gorm:"Not Null" json:"no_lap"`
	DpStatus         string    `gorm:"Not Null" json:"dp_status"`
	TanggalTransaksi time.Time `gorm:"Not Null" json:"tanggal_transaksi"`
	Booking          []Booking `gorm:"foreignKey:kode_transaksi"`
}
