package models

type PaketWisata struct {
	Id_paket        int             `gorm:"primaryKey;autoIncrement;" json:"id_paket"`
	Nama_paket      string          `json:"nama_paket"`
	Harga_paket     int             `json:"harga_paket"`
	Gambar_paket    string          `json:"gambar_paket"`
	Destinasi       int             `json:"destinasi"`
	AkomodasiWisata int             `json:"akomodasi"`
	DestinasiWisata DestinasiWisata `gorm:"foreignKey:Destinasi;references:Id_destinasi;"`
	Akomodasi       Akomodasi       `gorm:"foreignKey:AkomodasiWisata;references:Id_akomodasi;"`
}

type Join struct {
	Nama_paket       string `json:"nama_paket"`
	Harga_paket      int    `json:"harga_paket"`
	Gambar_paket     string `json:"gambar_paket"`
	Nama_destinasi   string `json:"nama_destinasi"`
	Lokasi_wisata    string `json:"lokasi_wisata"`
	Nama_akomodasi   string `json:"nama_akomodasi"`
	Lokasi_akomodasi string `json:"lokasi_akomodasi"`
	Fasilitas        string `json:"fasilitas"`
}
