package models

type Kepengurusan struct {
	IdKepengurusan int    `gorm:"primaryKey;autoIncrement;" json:"idKepengurusan"`
	Foto           string `json:"foto"`
	NamaPengurus   string `gorm:"not null" json:"namaPengurus"`
	Jabatan        string `gorm:"not null" json:"jabatan"`
	TglMasuk       string `gorm:"not null" json:"tglMasuk"`
	JkPengurus     string `gorm:"not null" json:"jkPengurus"`
	TtlPengurus    string `json:"ttlPengurus"`
	AlamatPengurus string `json:"alamatPengurus"`
	TelpPengurus   string `gorm:"not null" json:"telpPengurus"`
}

type Donatur struct {
	IdDonatur        int    `gorm:"primaryKey;autoIncrement;" json:"idDonatur"`
	NamaDonatur      string `gorm:"not null" json:"namaDonatur"`
	PekerjaanDonatur string `gorm:"not null" json:"pekerjaanDonatur"`
	TelpDonatur      string `gorm:"not null" json:"telpDonatur"`
}

type Sumbangan struct {
	IdSumbangan int     `gorm:"primaryKey;autoIncrement;" json:"idSumbangan"`
	IDDonatur   int     `json:"idDonatur"`
	Tanggal     string  `gorm:"not null" json:"tanggal"`
	Donatur     Donatur `gorm:"foreignKey:IDDonatur;references:IdDonatur;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"Donatur"`
	Jumlah      int     `gorm:"not null" json:"jumlah"`
}

type Sekolah struct {
	IdSekolah     int    `gorm:"primaryKey;autoIncrement;" json:"idSekolah"`
	NamaSekolah   string `gorm:"not null" json:"namaSekolah"`
	AlamatSekolah string `gorm:"not null" json:"alamatSekolah"`
	TelpSekolah   string `gorm:"not null" json:"telpSekolah"`
	Anak          []Anak `gorm:"foreignKey:SekolahID;references:IdSekolah;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" `
}

type Anak struct {
	NoInduk          int                `gorm:"primaryKey;autoIncrement;" json:"noInduk"`
	SekolahID        int                `json:"sekolahId"`
	Kelas            int                `gorm:"not null" json:"kelas"`
	NamaAnak         string             `gorm:"not null" json:"namaAnak"`
	TtlAnak          string             `gorm:"not null" json:"ttlAnak"`
	JkAnak           string             `gorm:"not null" json:"jkAnak"`
	GolDarah         string             `json:"golDarah"`
	Agama            string             `json:"agama"`
	TinggiBadan      int                `json:"tinggiBadan"`
	BeratBadan       int                `json:"beratBadan"`
	StatusAnak       string             `gorm:"not null" json:"status"`
	OrangTua         string             `json:"orangTua"`
	JumlahSaudara    int                `json:"jumlahSaudara"`
	TglMasuk         string             `gorm:"not null" json:"tglMasuk"`
	StatusPanti      string             `json:"statusPanti"`
	RiwayatKesehatan []RiwayatKesehatan `gorm:"foreignKey:No_Induk;references:NoInduk;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" `
	Adopsi           []Adopsi           `gorm:"foreignKey:IndukNo;references:NoInduk;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" `
}

type RiwayatKesehatan struct {
	IdKesehatan      int    `gorm:"primaryKey;autoIncrement;" json:"idKesehatan"`
	No_Induk         int    `json:"no_Induk"`
	TglPeriksa       string `gorm:"not null" json:"tglPeriksa"`
	HasilPemeriksaan string `gorm:"not null" json:"hasilPemeriksaaan"`
}

type Adopsi struct {
	IdAdopsi         int    `gorm:"primaryKey;autoIncrement;" json:"idAdopsi"`
	IndukNo          int    `json:"nooinduk"`
	TglAdopsi        string `gorm:"not null" json:"tglAdopsi"`
	NamaPengadopsi   string `gorm:"not null" json:"namaPengadopsi"`
	Umur             int    `gorm:"not null" json:"umur"`
	Alamat           string `gorm:"not null" json:"alamat"`
	StatusPerkawinan string `gorm:"not null" json:"statusPerkawinan"`
	Pekerjaan        string `gorm:"not null" json:"pekerjaan"`
	Penghasilan      int    `gorm:"not null" json:"penghasilan"`
}
