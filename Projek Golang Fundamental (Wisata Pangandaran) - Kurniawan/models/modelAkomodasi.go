package models

type Akomodasi struct {
	Id_akomodasi     int    `gorm:"primaryKey;autoIncrement;" json:"id_akomodasi"`
	Nama_akomodasi   string `json:"nama_akomodasi"`
	Lokasi_akomodasi string `json:"lokasi_akomodasi"`
	Harga_kamar      int    `json:"harga_kamar"`
	Fasilitas        string `json:"fasilitas"`
	Gambar_akomodasi string `json:"gambar_akomodasi"`
}
