package models

import (
	"time"
)

type Booking struct {
	IdBooking     int       `gorm:"primaryKey;autoIncrement" json:"id_booking"`
	NamaTeam      string    `gorm:"Not Null" json:"nama_team"`
	NoHp          string    `gorm:"Not Null" json:"no_hp"`
	NoLap         int       `gorm:"Not Null;" json:"id_lapangan"`
	Tanggal       string    `gorm:"Not Null" json:"tanggal"`
	Harga         int       `gorm:"column:harga_lap" json:"harga"`
	TanggalMain   time.Time `gorm:"Not Null" sql:"type:timestamp without time zone;" json:"tanggal_main"`
	TanggalPesan  time.Time `gorm:"Not Null" json:"tanggal_booking"`
	KodeTransaksi int       `grom:"Not Null" json:"kode_trans"`
	DpStatus      string    `gorm:"Not Null" json:"dp_status"`
}
