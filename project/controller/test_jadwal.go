package controller

import (
	"encoding/json"
	"net/http"
	"project/connection"
	"project/models"
	"strings"

	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	Db = connection.ConnectToDb()
}

func GetJadwal(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Jadwal
		Db.Find(&data)

		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "error encode ke json ", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}
	http.Error(w, "error not found ", 404)
}

func PostJadwal(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.Jadwal
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON ", 500)
			return
		}
		Db.Create(&data)
		w.Write([]byte("Sukses post data"))
		w.WriteHeader(200)
		return

	}
	http.Error(w, "Error Not Found ", 404)
}

func DeleteJadwal(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		u := r.URL.String()
		var id_jadwal []string = strings.Split(u, "/")
		if id_jadwal[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := Db.First(&models.Jadwal{}, "id_jadwal = ?", id_jadwal[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		Db.Delete(&models.Jadwal{}, Db.Where("id_jadwal = ?", id_jadwal[2]))

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func UpdateJadwal(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		u := r.URL.String()
		var id_jadwal []string = strings.Split(u, "/")
		if id_jadwal[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		
		decoder := json.NewDecoder(r.Body)
		var data models.Jadwal
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON ", 500)
			return
		}

		err := Db.First(&models.Jadwal{}, "id_jadwal = ?", id_jadwal[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		Db.Model(&models.Jadwal{}).Where("id_jadwal = ? ", id_jadwal[2]).Updates(&data)

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func ViewJadwal(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var value []models.Jadwal
		Db.Find(&value)
		type All struct {
			Id_Jadwal int     `json:"id_jadwal"`
			Nama_Matkul string `json:"nama_matkul"`
			Ruangan     string `json:"ruangan"`
			Nama_Dosen  string `json:"nama_dosen"`
			Hari        string `json:"hari"`
			Semester    int    `json:"semester"`
		}
		var data []All
		Db.Table("dosens").Select("jadwals.id_jadwal, dosens.nama_dosen, matkuls.nama_matkul, jadwals.ruangan, jadwals.hari, khs.semester").
		Joins("inner join jadwals on dosens.nip = jadwals.nip_dosen inner join matkuls on jadwals.mk_kode = matkuls.kode_mk inner join khs on matkuls.kode_mk = khs.kode_matkul").
		Scan(&data)
		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "error encode ke json ", 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}
	http.Error(w, "error not found ", 404)
}
