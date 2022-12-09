package models

type Doctor struct {
	Doc_id    int         `gorm:"primaryKey;autoIncrement;not null;" json:"doctor_id"`
	Doc_nama  string      `gorm:"not null;" json:"Doctor_name"`
	Job_spes  int         `gorm:"not null;" json:"job"`
	Jadwal    string      `gorm:"not null;" json:"jadwal"`
	Doc_age   string      `gorm:"not null" json:"doc_age"`
	Doc_pict  string      `gorm:"not null" json:"doc_pict"`
	Pasien    []Pasien    `gorm:"foreignKey:Pas_doc; not null;" json:"dokter"`
	Detaildoc []Detaildoc `gorm:"foreignKey:Detail_id" json:"detail_id"`
}

type JobSpef struct {
	Job_id   int      `gorm:"primaryKey;" json:"Job_Spef"`
	Job_nama string   `json:"Job_Name"`
	Doctor   []Doctor `gorm:"foreignKey:Job_spes" json:"Job_Spes"`
}

type Medicine struct {
	Med_id    int      `gorm:"primaryKey;" json:"medicine_id"`
	Med_nama  string   `json:"Medicine_name"`
	Med_stok  int      `json:"stok"`
	Med_rules string   `json:"aturan"`
	Med_exp   int      `gorm:"not null" json:"med_exp"`
	Pasien    []Pasien `gorm:"foreignKey:Pas_obat" json:"obat"`
}

type Room struct {
	Room_id   int      `gorm:"primaryKey;" json:"room_id"`
	Room_nama string   `json:"Room_name"`
	Room_spes string   `json:"space"`
	Pasien    []Pasien `gorm:"foreignKey:Pas_room" json:"kamar"`
}

type Pasien struct {
	Pas_id        int    `gorm:"primaryKey;" json:"pasien_id"`
	Pas_nama      string `json:"nama"`
	Pas_room      int    `json:"kamar"`
	Pas_obat      int    `json:"obat"`
	Pas_doc       int    `json:"dokter"`
	Detail_ttl    string `json:"detail_ttl"`
	Detail_age    string `json:"detail_age"`
	Detail_jk     string `json:"detail_jk"`
	Detail_addres string `json:"detail_address"`
	Detail_stat   string `json:"detail_stat"`
	Detail_desc   string `json:"detail_desc"`
}

type Detaildoc struct {
	Detail_id  int    `json:"detail_id"`
	Detail_ttl string `json:"detail_ttl"`
	Detail_edu string `json:"detail_edu"`
	Detail_sex string `json:"detail_sex"`
}
