package models

type Dosen struct {
	Nip          int      `gorm:"primaryKey;autoIncrement;" json:"nip"`
	Nama_Dosen   string   `json:"nama_dosen"`
	Alamat_Dosen string   `json:"alamat_dosen"`
	Telp         string   `json:"telp"`
	Jadwal       []Jadwal `gorm:"foreignKey:Nip_Dosen;" json:"jadwal"`
}

type Jadwal struct {
	Id_Jadwal int    `gorm:"primaryKey;autoIncrement;" json:"id_jadwal"`
	Mk_Kode   int    `json:"mk_kode"`
	Nip_Dosen int    `json:"nip_dosen"`
	Ruangan   string `json:"ruangan"`
	Hari      string `json:"hari"`
	Matkul    Matkul `gorm:"foreignKey:Mk_Kode;references:Kode_Mk;" json:"matkul"`
}

type Matkul struct {
	Kode_Mk     int    `gorm:"primaryKey;autoIncrement;" json:"kode_mk"`
	Nama_Matkul string `json:"nama_matkul"`
	Sks         int    `json:"sks"`
}

type Khs struct {
	Id_Khs       int    `gorm:"primaryKey;autoIncrement;" json:"id_khs"`
	Nim_Mhs      int    `json:"nim_mhs"`
	Kode_Matkul  int    `json:"kode_matkul"`
	Semester     int    `json:"semester"`
	Thn_Akademik int    `json:"thn_akademik"`
	Grade        string `json:"grade"`
	Matkul       Matkul `gorm:"foreignKey:Kode_Matkul;references:Kode_Mk;" json:"matkul"`
}

type Mahasiswa struct {
	Nim              int    `gorm:"primaryKey;autoIncrement;" json:"nim"`
	Foto             string `json:"foto"`
	Nama_Mahasiswa   string `json:"nama_mahasiswa"`
	Alamat_Mahasiswa string `json:"alamat_mahasiswa"`
	Jk_Mhs           string `json:"jk_mhs"`
	Khs              []Khs  `gorm:"foreignKey:Nim_Mhs" json:"khs"`
}
