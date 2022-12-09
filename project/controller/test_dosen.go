package controller

import (
	"encoding/json"
	"net/http"
	"project/connection"
	"project/models"
	"strings"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = connection.ConnectToDb()
}

func GetDosen(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Dosen
		db.Find(&data)

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

func PostDosen(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.Dosen
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON ", 500)
			return
		}
		db.Create(&data)
		w.Write([]byte("Sukses post data"))
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error Not Found ", 404)
}

func DeleteDosen(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		u := r.URL.String()
		var nip []string = strings.Split(u, "/")
		if nip[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := db.First(&models.Dosen{}, "nip = ?", nip[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		db.Delete(&models.Dosen{}, db.Where("nip = ?", nip[2]))

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func UpdateDosen(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		u := r.URL.String()
		var nip []string = strings.Split(u, "/")
		if nip[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var data models.Dosen
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON ", 500)
			return
		}

		err := db.First(&models.Dosen{}, "nip = ?", nip[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		db.Model(&models.Dosen{}).Where("nip = ? ", nip[2]).Updates(&data)

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
