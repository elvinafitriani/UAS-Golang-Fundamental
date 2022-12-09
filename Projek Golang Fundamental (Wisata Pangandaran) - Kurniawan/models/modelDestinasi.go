package models

type DestinasiWisata struct {
	Id_destinasi   int    `gorm:"primaryKey;autoIncrement;" json:"id_destinasi"`
	Nama_destinasi string `json:"nama_destinasi"`
	Lokasi_wisata  string `json:"lokasi_wisata"`
	Harga_tiket    int    `json:"harga_tiket"`
	Deskripsi      string `json:"deskripsi"`
	Gambar         string `json:"gambar"`
}
