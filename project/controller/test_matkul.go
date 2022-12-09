package controller

import (
	"encoding/json"
	"net/http"
	"project/connection"
	"project/models"
	"strings"

	"gorm.io/gorm"
)

var DBm *gorm.DB

func init() {
	DBm = connection.ConnectToDb()
}

func GetMatkul(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Matkul
		DBm.Find(&data)

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

func PostMatkul(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.Matkul
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON ", 500)
			return
		}
		DBm.Create(&data)
		w.Write([]byte("Sukses post data"))
		w.WriteHeader(200)
		return

	}
	http.Error(w, "Error Not Found ", 404)
}

func DeleteMatkul(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		u := r.URL.String()
		var kode_mk []string = strings.Split(u, "/")
		if kode_mk[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := DBm.First(&models.Matkul{}, "kode_mk = ?", kode_mk[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		DBm.Delete(&models.Matkul{}, DBm.Where("kode_mk = ?", kode_mk[2]))

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func UpdateMatkul(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		u := r.URL.String()
		var kode_mk []string = strings.Split(u, "/")
		if kode_mk[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var data models.Matkul
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON ", 500)
			return
		}

		err := DBm.First(&models.Matkul{}, "kode_mk = ?", kode_mk[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		DBm.Model(&models.Matkul{}).Where("kode_mk = ? ", kode_mk[2]).Updates(&data)

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
