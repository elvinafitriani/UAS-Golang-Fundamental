package models

type Alumni struct {
	Id_alumni   int     `gorm:"primaryKey;" json:"id_alumni"`
	Foto        string  `json:"foto"`
	Nama_alumni string  `json:"nama_alumni"`
	Id_angkatan int     `json:"id_angkatan"`
	Id_prodi    int     `json:"id_prodi"`
	Total_infak int     `json:"total_infak"`
	Keterangan  string  `json:"keterangan"`
	Infak       []Infak `gorm:"foreignKey:Id_alumni;" json:"alumni"`
}

type Prodi struct {
	Id_prodi    int      `gorm:"primaryKey;" json:"id_prodi"`
	Nama_prodi  string   `json:"nama_prodi"`
	Id_angkatan int      `json:"id_angkatan"`
	Alumni      []Alumni `gorm:"foreignKey:Id_prodi;" json:"alumni"`
}

type Angkatan struct {
	Id_angkatan   int      `gorm:"primaryKey;" json:"id_angkatan"`
	Angkatan_ke   int      `json:"angkatan_ke"`
	Nama_angkatan string   `json:"nama_angkatan"`
	Prodi         Prodi    `gorm:"foreignKey:Id_angkatan;references:Id_angkatan;" json:"prodi"`
	Alumni        []Alumni `gorm:"foreignKey:Id_angkatan;" json:"alumni"`
}

type Infak struct {
	Tanggal      string `json:"tanggal"`
	Id_infak     int    `gorm:"primaryKey;" json:"id_infak"`
	Id_alumni    int    `json:"id_alumni"`
	Id_rekening  int    `json:"id_rekening"`
	Jumlah_infak int    `json:"jumlah_infak"`
}

type Rekening struct {
	Id_rekening   int     `gorm:"primaryKey;" json:"id_rekening"`
	Nama_rekening string  `json:"nama_rek"`
	Saldo         int     `json:"saldo"`
	Infak         []Infak `gorm:"foreignKey:Id_rekening;" json:"infak"`
}
