package controller

import (
	"encoding/json"
	"net/http"
	"project/connection"
	"project/models"
	"strings"

	"gorm.io/gorm"
)

var dB *gorm.DB

func init() {
	dB = connection.ConnectToDb()
}

func GetKhs(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Khs
		dB.Find(&data)

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

func PostKhs(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.Khs
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON ", 500)
			return
		}
		dB.Create(&data)
		w.Write([]byte("Sukses post data"))
		w.WriteHeader(200)
		return

	}
	http.Error(w, "Error Not Found ", 404)
}

func DeleteKhs(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		u := r.URL.String()
		var id_khs []string = strings.Split(u, "/")
		if id_khs[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := dB.First(&models.Khs{}, "id_khs = ?", id_khs[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		dB.Delete(&models.Khs{}, dB.Where("id_khs = ?", id_khs[2]))

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func UpdateKhs(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		u := r.URL.String()
		var id_khs []string = strings.Split(u, "/")
		if id_khs[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var data models.Khs
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON ", 500)
			return
		}

		err := dB.First(&models.Khs{}, "id_khs = ?", id_khs[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		dB.Model(&models.Khs{}).Where("id_khs = ? ", id_khs[2]).Updates(&data)

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func ViewKhs(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var value []models.Khs
		Db.Find(&value)
		type All struct {
			Nim_Mhs            int    `json:"nim_mhs"`
			Nama_Mahasiswa string `json:"nama_mahasiswa"`
			Nama_Matkul    string `json:"nama_matkul"`
			Semester       int    `json:"semester"`
			Thn_Akademik   int    `json:"thn_akademik"`
			Grade          string `json:"grade"`
		}
		var data []All
		Db.Table("matkuls").Select("khs.nim_mhs, mahasiswas.nama_mahasiswa, matkuls.nama_matkul,khs.semester, khs.thn_akademik, khs.grade").Joins("inner join khs on matkuls.kode_mk = khs.kode_matkul inner join mahasiswas on khs.nim_mhs = mahasiswas.nim").Scan(&data)
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
